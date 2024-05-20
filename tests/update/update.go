package update

import (
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/txaty/go-merkletree"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/utils"
)

type leaf struct {
	DataKey   [32]byte
	Price     *big.Int
	Timestamp *big.Int
}

func (data *leaf) Serialize() ([]byte, error) {
	uint256Type, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return nil, err
	}

	bytes32Type, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		return nil, err
	}

	bytesType, err := abi.NewType("bytes", "bytes", nil)
	if err != nil {
		return nil, err
	}

	arguments := abi.Arguments{
		{
			Type: uint256Type,
		},
		{
			Type: bytesType,
		},
		{
			Type: bytes32Type,
		},
	}

	bytes, err := arguments.Pack(
		data.Timestamp,
		math.U256Bytes(data.Price),
		data.DataKey,
	)
	if err != nil {
		return nil, err
	}

	return crypto.Keccak256(bytes), nil
}

func GenerateUpdate(privKey *ecdsa.PrivateKey, asset string) (*types.MerkleRootUpdate, error) {
	dataKey, err := utils.AsciiToPaddedHex(asset)
	if err != nil {
		return nil, err
	}

	price := big.NewInt(100)
	timestamp := big.NewInt(time.Now().Unix())

	blocks := make([]merkletree.DataBlock, 0)
	for i := 0; i < 2; i++ {
		blocks = append(blocks, &leaf{
			DataKey:   dataKey,
			Price:     price,
			Timestamp: timestamp,
		})
	}

	tree, err := merkletree.New(&merkletree.Config{
		Mode:             merkletree.ModeProofGenAndTreeBuild,
		SortSiblingPairs: true,
		HashFunc: func(bytes []byte) ([]byte, error) {
			return crypto.Keccak256(bytes), nil
		},
	}, blocks)
	if err != nil {
		return nil, err
	}

	b := []byte("\x19Ethereum Signed Message:\n32")
	res := append(common.TrimLeftZeroes(b), crypto.Keccak256(tree.Root)...)

	sig, err := crypto.Sign(crypto.Keccak256(res), privKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27

	merkleProof := make([][32]byte, len(tree.Proofs[0].Siblings))
	for i, proof := range tree.Proofs[0].Siblings {
		copy(merkleProof[i][:], proof)
	}

	return &types.MerkleRootUpdate{
		DataKey:       dataKey,
		NewMerkleRoot: [32]byte(tree.Root),
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

	blocks := make([]merkletree.DataBlock, 0)
	for i := 0; i < 2; i++ {
		blocks = append(blocks, &leaf{
			DataKey:   dataKey,
			Price:     price,
			Timestamp: timestamp,
		})
	}

	tree, err := merkletree.New(&merkletree.Config{
		Mode:             merkletree.ModeProofGenAndTreeBuild,
		SortSiblingPairs: true,
		HashFunc: func(bytes []byte) ([]byte, error) {
			return crypto.Keccak256(bytes), nil
		},
	}, blocks)
	if err != nil {
		return nil, err
	}

	b := []byte("\x19Ethereum Signed Message:\n32")
	res := append(common.TrimLeftZeroes(b), crypto.Keccak256(tree.Root)...)

	sig, err := crypto.Sign(crypto.Keccak256(res), privKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27

	return &fetcher.EntangleFeedProof{
		Key:          asset,
		MerkleRoot:   common.BytesToHash(tree.Root),
		MerkleProofs: tree.Proofs[0].Siblings,
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
