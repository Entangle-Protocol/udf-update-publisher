package publisher

import (
	"fmt"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/utils"
)

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
	if !rootValid {
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
