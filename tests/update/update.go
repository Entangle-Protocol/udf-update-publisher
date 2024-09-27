package update

import (
	"crypto/ecdsa"
	"math/big"
	"time"

	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/utils"
)

var LeafEncodings = []string{
	smt.SOL_UINT256,
	smt.SOL_BYTES,
	smt.SOL_BYTES32,
}

func GenerateUpdate(privKey *ecdsa.PrivateKey, asset string) (*types.MerkleRootUpdate, error) {
	dataKey, err := utils.AsciiToPaddedHex(asset)
	if err != nil {
		return nil, err
	}

	price := big.NewInt(100)
	timestamp := big.NewInt(time.Now().Unix())

	value := []any{timestamp, math.U256Bytes(price), dataKey}

	tree, err := smt.Of([][]any{value}, LeafEncodings)
	if err != nil {
		return nil, err
	}

	res := accounts.TextHash(tree.GetRoot())

	sig, err := crypto.Sign(res, privKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27

	proof, err := tree.GetProofWithIndex(0)
	if err != nil {
		return nil, err
	}

	merkleProof := make([][32]byte, len(proof))
	for i, proof := range proof {
		copy(merkleProof[i][:], proof)
	}

	return &types.MerkleRootUpdate{
		DataKey:       dataKey,
		NewMerkleRoot: [32]byte(tree.GetRoot()),
		MerkleProof:   merkleProof,
		Signatures: []types.ECDSASignature{
			{
				R: common.BytesToHash(sig[:32]),
				S: common.BytesToHash(sig[32:64]),
				V: sig[64],
			},
		},
		Price:     price,
		Timestamp: timestamp,
	}, nil
}

func GenerateProof(privKey *ecdsa.PrivateKey, asset string) (*fetcher.EntangleFeedProof, error) {
	dataKey, err := utils.AsciiToPaddedHex(asset)
	if err != nil {
		return nil, err
	}

	price := big.NewInt(100)
	timestamp := big.NewInt(time.Now().Unix())

	value := []any{timestamp, math.U256Bytes(price), dataKey}

	tree, err := smt.Of([][]any{value}, LeafEncodings)
	if err != nil {
		return nil, err
	}

	res := accounts.TextHash(tree.GetRoot())

	sig, err := crypto.Sign(res, privKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27

	proof, err := tree.GetProofWithIndex(0)
	if err != nil {
		return nil, err
	}

	return &fetcher.EntangleFeedProof{
		Key:          asset,
		MerkleRoot:   common.BytesToHash(tree.GetRoot()).Hex(),
		MerkleProofs: proof,
		Signatures: []fetcher.HashEncodedSignatureDoc{
			{
				R: common.BytesToHash(sig[:32]),
				S: common.BytesToHash(sig[32:64]),
				V: sig[64],
			},
		},
		Value: fetcher.FinalizedDataDoc{
			PriceData: price.Bytes(),
			Timestamp: timestamp.Int64(),
		},
	}, nil
}
