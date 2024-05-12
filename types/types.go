package types

import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

// Types that is used throughout the project

type ECDSASignature struct {
	V byte // Recovery ID to reduce the correct pk
	R common.Hash // X coordinate of signature point
	S common.Hash // Signature component
}

// Type that aggregates argument fields that gets passed to PullOracle
type MerkleRootUpdate struct {
	NewMerkleRoot [32]byte
	MerkleProof [][32]byte
	Signatures []ECDSASignature
	Price *big.Int
	Timestamp *big.Int
}
