package fetcher

type IFetcher interface {
	GetFeedProofs() (*EntangleFeedProof, error)
}
