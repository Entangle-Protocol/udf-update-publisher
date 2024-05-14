package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const getAssetDataUriPath = "/getAssetData"

type RestFetcher struct {
	finalizedSnapshotUrl string
	client               *http.Client
}

func NewRestFetcher(
	client *http.Client,
	finalizedSnapshotUrl string,
) *RestFetcher {
	return &RestFetcher{
		client:               client,
		finalizedSnapshotUrl: finalizedSnapshotUrl,
	}
}

func (r *RestFetcher) GetFeedProofs(ctx context.Context, assetKey string) (*EntangleFeedProof, error) {
	url := r.finalizedSnapshotUrl + getAssetDataUriPath + "?assetKey=" + assetKey

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var proof EntangleFeedProof
		if err := json.NewDecoder(resp.Body).Decode(&proof); err != nil {
			return nil, err
		}

		return &proof, nil
	default:
		type getAssetDataError struct {
			Error string `json:"error"`
		}

		var dataError getAssetDataError
		if err := json.NewDecoder(resp.Body).Decode(&dataError); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get feed proofs: %v", dataError.Error)
	}
}
