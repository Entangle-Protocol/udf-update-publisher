package tests

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	ethsim "github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/require"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/config"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/contrib/contracts/datafeeds/PullOracle"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/keystore"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/publisher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/tests/deploy"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/tests/update"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/transactor"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/utils"
)

var (
	backend    *ethsim.Backend
	httpServer *http.Server

	adminKey *ecdsa.PrivateKey
	admin    common.Address

	appConfig      *config.AppConfig
	netConfig      config.NetworkConfig
	protocolConfig = deploy.ProtocolConfig{
		ID:            [32]byte{0x11},
		ConsensusRate: big.NewInt(6000),
	}

	pullOracleAddress common.Address
)

func TestMain(m *testing.M) {
	cfg, err := config.LoadConfig(filepath.Join("testdata", "config.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	appConfig = cfg
	netConfig = cfg.Networks["eth_sepolia"]

	adminKey, err = keystore.ParseKeyFromHex(netConfig.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	pub, _ := adminKey.Public().(*ecdsa.PublicKey)
	admin = crypto.PubkeyToAddress(*pub)

	balance, ok := big.NewInt(0).SetString("90000000000000000000", 10)
	if !ok {
		log.Fatal("Failed to parse bigInt")
	}

	backend = ethsim.NewBackend(ethtypes.GenesisAlloc{
		admin: {Balance: balance},
	})

	contracts, err := deploy.DeployContracts(context.Background(), backend, adminKey, protocolConfig)
	if err != nil {
		log.Fatal(err)
	}
	pullOracleAddress = contracts.PullOracle

	httpServer = &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: http.HandlerFunc(getAsset),
	}

	go httpServer.ListenAndServe()

	code := m.Run()

	_ = backend.Close()
	_ = httpServer.Shutdown(context.Background())

	os.Exit(code)
}

func TestPublisher_PublishUpdate_Simulate(t *testing.T) {
	r := require.New(t)

	ctx := context.Background()

	client := backend.Client()
	chainID, err := client.ChainID(ctx)
	r.NoError(err)

	tx, err := transactor.NewTransactor(ctx, client, adminKey, chainID, pullOracleAddress)
	r.NoError(err)

	url := "http://" + httpServer.Addr
	fetcher := fetcher.NewRestFetcher(http.DefaultClient, url)
	pub := publisher.NewUpdatePublisher(appConfig.Publisher, []transactor.ITransactor{tx}, fetcher, appConfig.DataKeys, []config.AssetSet{})

	pullOracle, err := PullOracle.NewPullOracle(pullOracleAddress, client)
	r.NoError(err)

	dataKey, err := utils.AsciiToPaddedHex(appConfig.DataKeys[0])
	r.NoError(err)

	info, err := pullOracle.LatestUpdate(&bind.CallOpts{}, dataKey)
	r.NoError(err)
	r.Zero(info.LatestPrice.Int64())
	r.Zero(info.LatestTimestamp.Int64())

	err = pub.PublishUpdate(ctx)
	r.NoError(err)
	backend.Commit()

	info, err = pullOracle.LatestUpdate(&bind.CallOpts{}, dataKey)
	r.NoError(err)
	r.NotZero(info.LatestPrice.Int64())
	r.NotZero(info.LatestTimestamp.Int64())
}

// func TestPublisher_PublishMultipleUpdate_Simulate(t *testing.T) {
// 	r := require.New(t)
// 	ctx := context.Background()

// 	spotterID := "prices-feed1"
// 	assetKeys := []string{"NGL/USD", "ETH/USD", "BTC/USD"}

// 	client := backend.Client()
// 	chainID, err := client.ChainID(ctx)
// 	r.NoError(err)

// 	tx, err := transactor.NewTransactor(ctx, client, adminKey, chainID, pullOracleAddress)
// 	r.NoError(err)

// 	assetSet := config.AssetSet{SourceID: spotterID, DataKeys: assetKeys}

// 	url := "https://pricefeed.entangle.fi"
// 	fetcher := fetcher.NewRestFetcher(http.DefaultClient, url)
// 	pub := publisher.NewUpdatePublisher(appConfig.Publisher, []transactor.ITransactor{tx}, fetcher, appConfig.DataKeys, []config.AssetSet{assetSet})

// 	pullOracle, err := PullOracle.NewPullOracle(pullOracleAddress, client)
// 	r.NoError(err)

// 	dataKey, err := utils.AsciiToPaddedHex(appConfig.DataKeys[0])
// 	r.NoError(err)

// 	info, err := pullOracle.LatestUpdate(&bind.CallOpts{}, dataKey)
// 	r.NoError(err)
// 	r.Zero(info.LatestPrice.Int64())
// 	r.Zero(info.LatestTimestamp.Int64())

// 	err = pub.PublishMultipleUpdate(ctx)
// 	r.NoError(err)
// 	backend.Commit()

// 	info, err = pullOracle.LatestUpdate(&bind.CallOpts{}, dataKey)
// 	r.NoError(err)
// 	r.NotZero(info.LatestPrice.Int64())
// 	r.NotZero(info.LatestTimestamp.Int64())
// }

// func TestPublisher_Testnet(t *testing.T) {
// 	r := require.New(t)
// 	ctx := context.Background()

// 	client, err := ethclient.Dial(netConfig.TargetChainURL)
// 	r.NoError(err)

// 	chain, err := client.ChainID(ctx)
// 	r.NoError(err)

// 	key, err := keystore.ParseKeyFromHex(netConfig.PrivateKey)
// 	r.NoError(err)

// 	tx, err := transactor.NewTransactor(ctx, client, key, chain, netConfig.PullOracleAddress)
// 	r.NoError(err)

// 	fetcher := fetcher.NewRestFetcher(http.DefaultClient, appConfig.FinalizeSnapshotURL)
// 	pub := publisher.NewUpdatePublisher(appConfig.Publisher, []transactor.ITransactor{tx}, fetcher, appConfig.DataKeys, appConfig.Assets)

// 	err = pub.PublishMultipleUpdate(ctx)
// 	r.NoError(err)
// }

func getAsset(w http.ResponseWriter, r *http.Request) {
	dataKey := strings.TrimPrefix(r.RequestURI, "/asset/")
	proof, err := update.GenerateProof(adminKey, dataKey)

	m := map[string]any{
		"calldata": proof,
	}

	if err != nil {
		m["error"] = err
	}

	json.NewEncoder(w).Encode(&m)
}
