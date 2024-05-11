package publisher

import (
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"

	log "github.com/sirupsen/logrus"
)

type UpdatePublisher struct {
	transactor ITransactor
	fetcher	fetcher.IFetcher
}

func NewUpdatePublisher(transactor ITransactor, fetcher fetcher.IFetcher) *UpdatePublisher {
	return &UpdatePublisher{
		transactor: transactor,
		fetcher: fetcher,
	}
}

func (up *UpdatePublisher) PublishUpdate() error {
	proofs, err := up.fetcher.GetFeedProofs()
	if err != nil {
		log.Errorf("Failed to get feed proofs: %v", err)
		return err
	}

	log.Infof("Got feed proofs: %+v", proofs)

	return nil
}

