package transactor

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/contrib/contracts/datafeeds/PullOracle"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/keystore"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethsim "github.com/ethereum/go-ethereum/ethclient/simulated"
)

func simTestBackend(testAddr ethcommon.Address) *ethsim.Backend {
	balance, ok := big.NewInt(0).SetString("90000000000000000000", 10)
	if !ok {
		panic("Failed to parse bigInt")
	}
	return ethsim.NewBackend(
		ethtypes.GenesisAlloc{
			testAddr: {Balance: balance},
		},
	)
}

func TestSendUpdate(t *testing.T) {
	ctx := context.Background()
	// Create eth fake backend
	testAddrPk := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	key, err := keystore.ParseKeyFromHex(testAddrPk)
	assert.Nil(t, err)

	testAddr := ethcommon.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	backend := simTestBackend(testAddr)
	defer func() {
		err := backend.Close()
		assert.Nil(t, err)
	}()

	client := backend.Client()
	chainID, err := client.ChainID(ctx)
	assert.Nil(t, err)

	deployOpts, err := bind.NewKeyedTransactorWithChainID(
		key,
		chainID,
	)
	assert.Nil(t, err)
	pullOracleAddress, _, _, err := PullOracle.DeployPullOracle(deployOpts, client)
	assert.Nil(t, err)

	blockHash := backend.Commit()
	fmt.Println("commited state with block hash:", blockHash)

	transactor, err := NewTransactor(ctx, client, key, chainID, pullOracleAddress)
	assert.Nil(t, err)

	var merkleUpdate types.MerkleRootUpdate
	err = gofakeit.Struct(&merkleUpdate)
	assert.Nil(t, err)

	err = transactor.SendUpdate(&merkleUpdate)
	// Expect error due to invalid merkle proof (since we generate it using random data)
	assert.Contains(t, err.Error(), "execution reverted")
}
