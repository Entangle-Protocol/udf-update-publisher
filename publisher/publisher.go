package publisher

import (
	"fmt"
	"context"
	"regexp"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/transactor"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/utils"
)

type UpdatePublisher struct {
	dataKeys    []string
	transactors []transactor.ITransactor
	fetcher     fetcher.IFetcher
}

const MerkleRootRegex = "^0x[0-9A-Fa-f]{64}$"

func NewMerkleUpdateFromProof(proof *fetcher.EntangleFeedProof) (*types.MerkleRootUpdate, error) {
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

	dataKey, err := utils.AsciiToPaddedHex(proof.Key)
	if err != nil {
		log.Errorf("Failed to convert key to padded hex dataKey: %v", err)
		return nil, err
	}

	rootValid, _ := regexp.MatchString(MerkleRootRegex, proof.MerkleRoot)
	if (!rootValid) {
		log.Errorf("Failed to parse provided merkle root: %s", proof.MerkleRoot)
		return nil, fmt.Errorf("Invalid merkle root")
	}

	return &types.MerkleRootUpdate{
		DataKey:       dataKey,
		NewMerkleRoot: common.HexToHash(proof.MerkleRoot),
		MerkleProof:   merkleProof,
		Signatures:    signatures,
		Price:         big.NewInt(0).SetBytes(proof.Value.PriceData),
		Timestamp:     big.NewInt(proof.Value.Timestamp),
	}, nil
}

func NewUpdatePublisher(transactors []transactor.ITransactor, fetcher fetcher.IFetcher, dataKeys []string) *UpdatePublisher {
	return &UpdatePublisher{
		dataKeys:    dataKeys,
		transactors: transactors,
		fetcher:     fetcher,
	}
}

func (up *UpdatePublisher) PublishUpdate(ctx context.Context) error {
	for _, dataKey := range up.dataKeys {
		proofs, err := up.fetcher.GetFeedProofs(ctx, dataKey)
		if err != nil {
			log.Errorf("Failed to get feed proofs: %v", err)
			return err
		}
		log.Infof("Got feed proofs: %+v", proofs)

		update, err := NewMerkleUpdateFromProof(proofs)
		if err != nil {
			log.Errorf("Failed to create merkle update from proof: %v", err)
			return err
		}

		for _, transactor := range up.transactors {
			if err := transactor.SendUpdate(update); err != nil {
				log.Errorf("Failed to send update: %v", err)
				return err
			}
		}
	}

	return nil
}
