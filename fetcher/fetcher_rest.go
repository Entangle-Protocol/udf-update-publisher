package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	getAssetDataUriPath = "/asset"
	getSpottersUriPath  = "/spotters"
)

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
		type errorResponse struct {
			Error string `json:"error"`
		}

		var errResp errorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get feed proofs: %v", errResp.Error)
	}
}

func (r *RestFetcher) GetSpotterFeedsProofs(ctx context.Context, spotterID string, assetKeys []string) (*EntangleFeedsProofs, error) {
	url, err := url.JoinPath(r.finalizedSnapshotUrl, getSpottersUriPath, spotterID)
	if err != nil {
		return nil, err
	}

	url += "?assets=" + strings.Join(assetKeys, ",")

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
			Calldata *EntangleFeedsProofs `json:"calldata"`
		}

		var res result
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return nil, err
		}

		return res.Calldata, nil
	default:
		type errorResponse struct {
			Error string `json:"error"`
		}

		var errResp errorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get spotter feeds proofs: %v", errResp.Error)
	}
}
