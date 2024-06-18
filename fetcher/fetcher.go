package fetcher

import "context"

type IFetcher interface {
	GetFeedProofs(ctx context.Context, assetKey string) (*EntangleFeedProof, error)
	GetSpotterFeedsProofs(ctx context.Context, spotterID string, assetKeys []string) (*EntangleFeedsProofs, error)
}
