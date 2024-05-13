package publisher

import (
	"math/big"

	log "github.com/sirupsen/logrus"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/transactor"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"
)

type UpdatePublisher struct {
	transactors []transactor.ITransactor
	fetcher     fetcher.IFetcher
}

func NewMerkleUpdateFromProof(proof *fetcher.EntangleFeedProof) *types.MerkleRootUpdate {
	merkleProof := make([][32]byte, len(proof.MerkleProofs))
	for i, proof := range proof.MerkleProofs {
		copy(merkleProof[i][:], proof)
	}
	signatures := make([]types.ECDSASignature, len(proof.Signatures))
	for i, sig := range proof.Signatures {
		signatures[i] = types.ECDSASignature{
			V: sig.V,
			R: sig.R,
			S: sig.S,
		}
	}

	return &types.MerkleRootUpdate{
		NewMerkleRoot: proof.MerkleRoot,
		MerkleProof:   merkleProof,
		Signatures:    signatures,
		Price:         big.NewInt(0).SetBytes(proof.Value.PriceData),
		Timestamp:     big.NewInt(proof.Value.Timestamp),
	}
}

func NewUpdatePublisher(transactors []transactor.ITransactor, fetcher fetcher.IFetcher) *UpdatePublisher {
	return &UpdatePublisher{
		transactors: transactors,
		fetcher:     fetcher,
	}
}

func (up *UpdatePublisher) PublishUpdate() error {
	proofs, err := up.fetcher.GetFeedProofs()
	if err != nil {
		log.Errorf("Failed to get feed proofs: %v", err)
		return err
	}

	update := NewMerkleUpdateFromProof(proofs)

	for _, transactor := range up.transactors {
		if err := transactor.SendUpdate(update); err != nil {
			log.Errorf("Failed to send update: %v", err)
			return err
		}
	}

	log.Infof("Got feed proofs: %+v", proofs)

	return nil
}
