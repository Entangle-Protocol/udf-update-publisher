package fetcher

import (
	"github.com/ethereum/go-ethereum/common"
)

// FeedProof holds the latest feed proof. Matches the schema from finalized-data-snapshot backend API
type EntangleFeedProof struct {
	MerkleRoot   common.Hash               `json:"merkleRoot"`
	Signatures   []HashEncodedSignatureDoc `json:"signatures"`
	MerkleProofs [][]byte                  `json:"merkleProofs"`
	Key          string                    `json:"key"`
	Value        FinalizedDataDoc          `json:"value"`
}

type HashEncodedSignatureDoc struct {
	R common.Hash
	S common.Hash
	V byte
}

type FinalizedDataDoc struct {
	Timestamp int64
	PriceData []byte
}
