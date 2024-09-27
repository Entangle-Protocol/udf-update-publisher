package fetcher

import (
	"github.com/ethereum/go-ethereum/common"
)

// FeedProof holds the latest feed proof. Matches the schema from finalized-data-snapshot backend API
type EntangleFeedProof struct {
	MerkleRoot   string                    `json:"merkleRoot"`
	Signatures   []HashEncodedSignatureDoc `json:"signatures"`
	MerkleProofs [][]byte                  `json:"merkleProofs,omitempty"`
	Key          string                    `json:"key"`
	Value        FinalizedDataDoc          `json:"value"`
}

type EntangleFeedsProofs struct {
	MerkleRoot string                    `json:"merkleRoot"`
	Signatures []HashEncodedSignatureDoc `json:"signatures"`
	Feeds      []FeedProof               `json:"feeds"`
}

func (f *EntangleFeedsProofs) Proofs() []*EntangleFeedProof {
	proofs := make([]*EntangleFeedProof, len(f.Feeds))
	for i, feed := range f.Feeds {
		proofs[i] = &EntangleFeedProof{
			MerkleRoot:   f.MerkleRoot,
			Signatures:   f.Signatures,
			MerkleProofs: feed.MerkleProofs,
			Key:          feed.Key,
			Value:        feed.Value,
		}
	}
	return proofs
}

type FeedProof struct {
	Key          string           `json:"key"`
	Value        FinalizedDataDoc `json:"value"`
	MerkleProofs [][]byte         `json:"merkleProofs,omitempty"`
}

type HashEncodedSignatureDoc struct {
	R common.Hash
	S common.Hash
	V byte
}

type FinalizedDataDoc struct {
	Timestamp int64  `json:"timestamp"`
	PriceData []byte `json:"data"`
}
