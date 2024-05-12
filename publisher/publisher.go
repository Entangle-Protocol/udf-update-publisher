package publisher

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
)

type UpdatePublisher struct {
	transactor ITransactor
	fetcher	fetcher.IFetcher
}

type ECDSASignature struct {
	V byte // Recovery ID to reduce the correct pk
	R common.Hash // X coordinate of signature point
	S common.Hash // Signature component
}

// Type that aggregates argument fields that gets passed to PullOracle
type MerkleRootUpdate struct {
	NewMerkleRoot [32]byte
	MerkleProof [][32]byte
	Signatures []ECDSASignature
	Price *big.Int
	Timestamp *big.Int
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

