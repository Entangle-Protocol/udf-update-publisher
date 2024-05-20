package transactor

import (
	"context"
	"math/big"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/contrib/contracts/datafeeds/PullOracle"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/keystore"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/tests/deploy"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/tests/update"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/utils"

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
	r := require.New(t)

	ctx := context.Background()
	// Create eth fake backend
	adminPk := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	admin := ethcommon.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	adminKey, err := keystore.ParseKeyFromHex(adminPk)
	r.NoError(err)

	protocolID := [32]byte{0x01}
	consensusRate := big.NewInt(6000)

	backend := simTestBackend(admin)
	defer func() {
		err := backend.Close()
		r.NoError(err)
	}()

	client := backend.Client()
	chainID, err := client.ChainID(ctx)
	r.NoError(err)

	deployed, err := deploy.DeployContracts(ctx, backend, adminKey, deploy.ProtocolConfig{
		ID:            protocolID,
		ConsensusRate: consensusRate,
	})
	r.NoError(err)

	transactor, err := NewTransactor(ctx, client, adminKey, chainID, deployed.PullOracle)
	r.NoError(err)

	pullOracle, err := PullOracle.NewPullOracle(deployed.PullOracle, client)
	r.NoError(err)

	testCases := []struct {
		desc      string
		getUpdate func() *types.MerkleRootUpdate
		wantErr   string
	}{
		{
			desc: "Valid Proof",
			getUpdate: func() *types.MerkleRootUpdate {
				update, err := update.GenerateUpdate(adminKey, "NGL/USDT")
				r.NoError(err)
				return update
			},
		},
		{
			desc: "Invalid",
			getUpdate: func() *types.MerkleRootUpdate {
				var u types.MerkleRootUpdate
				err := gofakeit.Struct(&u)
				r.NoError(err)
				return &u
			},
			wantErr: "execution reverted",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			update := tC.getUpdate()
			err = transactor.SendUpdate(update)
			backend.Commit()

			if tC.wantErr != "" {
				r.ErrorContains(err, tC.wantErr)
			} else {
				r.NoError(err)

				result, err := pullOracle.LatestUpdate(&bind.CallOpts{}, update.DataKey)
				r.NoError(err)
				r.Equal(update.Price, result.LatestPrice)
				r.Equal(update.Timestamp, result.LatestTimestamp)
			}
		})
	}
}

func NewMerkleUpdateFromProof(proof *fetcher.EntangleFeedProof) (*types.MerkleRootUpdate, error) {
	merkleProof := make([][32]byte, len(proof.MerkleProofs))
	for i, proof := range proof.MerkleProofs {
		copy(merkleProof[i][:], proof)
	}
	signatures := make([]types.ECDSASignature, len(proof.Signatures))
	for i, sig := range proof.Signatures {
		signatures[i] = types.ECDSASignature{
			V: sig.V,
			R: sig.R,
			S: sig.S,
		}
	}

	dataKey, err := utils.AsciiToPaddedHex(proof.Key)
	if err != nil {
		return nil, err
	}

	return &types.MerkleRootUpdate{
		DataKey:       dataKey,
		NewMerkleRoot: proof.MerkleRoot,
		MerkleProof:   merkleProof,
		Signatures:    signatures,
		Price:         big.NewInt(0).SetBytes(proof.Value.PriceData),
		Timestamp:     big.NewInt(proof.Value.Timestamp),
	}, nil
}
