package transactor

import (
	"context"
	"math/big"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/contrib/contracts/datafeeds/PullOracle"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/keystore"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/tests/deploy"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/tests/update"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethsim "github.com/ethereum/go-ethereum/ethclient/simulated"


	// ethmath "github.com/ethereum/go-ethereum/common/math"
	// "github.com/ethereum/go-ethereum/accounts"
	// smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	// "github.com/ethereum/go-ethereum/crypto"
	// "gitlab.ent-dx.com/entangle/pull-update-publisher/utils"
	// "fmt"
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
			_, err = transactor.SendUpdate(update)
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

// func TestSendMultipleUpdate(t *testing.T) {
// 	r := require.New(t)
//
// 	ctx := context.Background()
// 	// Create eth fake backend
// 	adminPk := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
// 	admin := ethcommon.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
// 	adminKey, err := keystore.ParseKeyFromHex(adminPk)
// 	r.NoError(err)
//
// 	protocolID := [32]byte{0x01}
// 	consensusRate := big.NewInt(6000)
//
// 	backend := simTestBackend(admin)
// 	defer func() {
// 		err := backend.Close()
// 		r.NoError(err)
// 	}()
//
// 	client := backend.Client()
// 	chainID, err := client.ChainID(ctx)
// 	r.NoError(err)
//
// 	deployed, err := deploy.DeployContracts(ctx, backend, adminKey, deploy.ProtocolConfig{
// 		ID:            protocolID,
// 		ConsensusRate: consensusRate,
// 	})
// 	r.NoError(err)
//
// 	transactor, err := NewTransactor(ctx, client, adminKey, chainID, deployed.PullOracle)
// 	r.NoError(err)
//
// 	pullOracle, err := PullOracle.NewPullOracle(deployed.PullOracle, client)
// 	r.NoError(err)
//
// 	type updateData struct {
// 		DataKey string
// 		Price string
// 		Timestamp *big.Int
// 	}
//
// 	testCases := []struct {
// 		desc string
// 		updateData []updateData
// 		expectedUpdateCalldata string
// 		containsErr string
// 		wantErr bool
// 	}{
// 		{
// 			desc: "Test with 2 valid updates",
// 			updateData: []updateData{
// 				{
// 					DataKey: "USDC/USD",
// 					Price: "1010000000000000000",
// 					Timestamp: big.NewInt(1712156481),
// 				},
// 				{
// 					DataKey: "ETH/USD",
// 					Price: "3230000000000000000000",
// 					Timestamp: big.NewInt(1712156481),
// 				},
// 				{
// 					DataKey: "BTC/USD",
// 					Price: "69510000000000000000000",
// 					Timestamp: big.NewInt(1712156481),
// 				},
// 				{
// 					DataKey: "NGL/USD",
// 					Price: "830000000000000000",
// 					Timestamp: big.NewInt(1712156481),
// 				},
// 			},
// 		},
// 	}
//
// 	var LeafEncodings = []string{
// 		smt.SOL_UINT256,
// 		smt.SOL_BYTES,
// 		smt.SOL_BYTES32,
// 	}
//
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			// Build merkle tree with the updates
// 			var merkleValues [][]any
// 			for _, update := range tC.updateData {
// 				dataKey, err := utils.AsciiToPaddedHex(update.DataKey)
// 				r.NoError(err)
//
// 				price, success := new(big.Int).SetString(update.Price, 10)
// 				r.True(success)
// 				
// 				value := []any{update.Timestamp, ethmath.U256Bytes(price), dataKey}
// 				merkleValues = append(merkleValues, value)
// 			}
//
// 			tree, err := smt.Of(merkleValues, LeafEncodings)
// 			r.NoError(err)
//
// 			root := accounts.TextHash(tree.GetRoot())
//
// 			// Sign the merkle root
// 			sig, err := crypto.Sign(root, adminKey)
// 			r.NoError(err)
// 			sig[64] += 27
// 			signatures := make([]types.ECDSASignature, 0)
// 			signatures = append(signatures, types.ECDSASignature{
// 				R: ethcommon.BytesToHash(sig[:32]),
// 				S: ethcommon.BytesToHash(sig[32:64]),
// 				V: sig[64],
// 			})
//
// 			// Build merkle root updates array
// 			merkleUpdates := make([]*types.MerkleRootUpdate, len(tC.updateData))
// 			for i, update := range tC.updateData {
// 				dataKey, err := utils.AsciiToPaddedHex(update.DataKey)
// 				r.NoError(err)
// 				proof, err := tree.GetProofWithIndex(i)
// 				r.NoError(err)
// 				merkleProof := make([][32]byte, len(proof))
// 				for i, proof := range proof {
// 					copy(merkleProof[i][:], proof)
// 				}
//
// 				price, success := new(big.Int).SetString(update.Price, 10)
// 				r.True(success)
//
// 				update := &types.MerkleRootUpdate{
// 					DataKey: dataKey,
// 					NewMerkleRoot: ethcommon.BytesToHash(root[:]),
// 					MerkleProof: merkleProof,
// 					Signatures: signatures,
// 					Price: price,
// 					Timestamp: update.Timestamp,
// 				}
//
// 				merkleUpdates[i] = update
// 			}
//
// 			// Get multiple update object
// 			multipleUpdate, err := types.NewMekrleRootUpdateMultipleFromUpdates(merkleUpdates)
// 			r.NoError(err)
//
// 			// Send update to PullOracle
// 			tx, err := transactor.SendMultipleUpdate(multipleUpdate)
// 			r.NoError(err)
// 			r.NotNil(tx)
// 			backend.Commit()
//
// 			rec, err := backend.Client().TransactionReceipt(ctx, tx.Hash())
// 			r.NoError(err)
//
// 			_ = rec
// 			r.Nil(rec)
//
// 			// Verify data in PullOracle
// 			for _, update := range tC.updateData {
// 				dataKey, err := utils.AsciiToPaddedHex(update.DataKey)
// 				r.NoError(err)
//
// 				result, err := pullOracle.LatestUpdate(&bind.CallOpts{}, dataKey)
// 				r.NoError(err)
//
// 				fmt.Println("result", result)
// 				// FIXME: Check results
//
// 			}
// 		})
// 	}
// }
