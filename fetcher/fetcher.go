package fetcher

import "context"

type IFetcher interface {
	GetFeedProofs(ctx context.Context, assetKey string) (*EntangleFeedProof, error)
}
