package fetcher

import (
	"net/http"
	"path"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type HashEncodedSignatureDoc struct {
	R common.Hash
	S common.Hash
	V byte
}

const GetAssetDataUriPath = "/getAssetData"

type FinalizedDataDoc struct {
	Timestamp int64  `bson:"timestamp"`
	PriceData []byte `bson:"priceData"`
}

// FeedProof holds the latest feed proof. Matches the schema from finalized-data-snapshot backend API
type EntangleFeedProof struct {
	MerkleRoot   common.Hash               `json:"merkleRoot"`
	Signatures   []HashEncodedSignatureDoc `json:"signatures"`
	MerkleProofs [][]byte                  `json:"merkleProofs"`
	Key          string                    `json:"key"`
	Value        FinalizedDataDoc    `json:"value"`
}

type RestFetcher struct {
	client *http.Client
	FinalizedSnapshotUrl string
}

func NewRestFetcher(
	finalizedSnapshotUrl string,
	client *http.Client,
) (*RestFetcher, error) {
	return &RestFetcher{
		client: client,
		FinalizedSnapshotUrl: finalizedSnapshotUrl,
	}, nil
}

func (r *RestFetcher) GetFeedProofs() (*EntangleFeedProof, error) {
	resp, err := r.client.Get(path.Join(r.FinalizedSnapshotUrl, GetAssetDataUriPath))
	if err != nil {
		return nil, err
	}

	log.Infof("Got response: %v", resp)
	return nil, nil
}
