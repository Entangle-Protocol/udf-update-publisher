package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const getAssetDataUriPath = "/asset"

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
	url, err := url.JoinPath(r.finalizedSnapshotUrl, getAssetDataUriPath, assetKey)
	if err != nil {
		return nil, err
	}

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
		type result struct {
			Calldata *EntangleFeedProof `json:"calldata"`
		}

		var res result
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return nil, err
		}

		return res.Calldata, nil
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
