package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethmath "github.com/ethereum/go-ethereum/common/math"
)

// Types that is used throughout the project

type ECDSASignature struct {
	R common.Hash // X coordinate of signature point
	S common.Hash // Signature component
	V byte        // Recovery ID to reduce the correct pk
}

func (s *ECDSASignature) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, s)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Type that aggregates argument fields that gets passed to PullOracle
type MerkleRootUpdate struct {
	DataKey       [32]byte
	NewMerkleRoot [32]byte
	MerkleProof   [][32]byte
	Signatures    []ECDSASignature
	Price         *big.Int
	Timestamp     *big.Int
}

type MultipleUpdateData struct {
	DataKey     [32]byte
	MerkleProof [][32]byte
	Price       big.Int
	Timestamp   big.Int
}

type MerkleRootUpdateMultiple struct {
	MerkleRoot [32]byte
	Signatures []ECDSASignature
	UpdateData []MultipleUpdateData
}

func NewMekrleRootUpdateMultipleFromUpdates(
	updates []*MerkleRootUpdate,
) (*MerkleRootUpdateMultiple, error) {
	if len(updates) == 0 {
		return nil, fmt.Errorf("no updates provided")
	}

	updateData := make([]MultipleUpdateData, len(updates))
	// QUESTIONABLE: Check that all updates have the same merkle root
	for i, update := range updates {
		if update.NewMerkleRoot != updates[0].NewMerkleRoot {
			return nil, fmt.Errorf("all updates must have the same merkle root")
		}

		updateData[i] = MultipleUpdateData{
			DataKey:     update.DataKey,
			MerkleProof: update.MerkleProof,
			Price:       *update.Price,
			Timestamp:   *update.Timestamp,
		}
	}

	merkleRoot := updates[0].NewMerkleRoot
	signatures := updates[0].Signatures

	return &MerkleRootUpdateMultiple{
		MerkleRoot: merkleRoot,
		Signatures: signatures,
		UpdateData: updateData,
	}, nil
}

/**
* Encode MerkleRootUpdateMultiple into a calldata to PullOracle.updateMultipleAssets.
*
* Update ABI specification for EVM PullOracle
* Format: <merkle_root><sigs_length><sigs_array><updates_length><updates_array>
*
* Structure:
* - `merkle_root` (32 bytes): The Merkle root hash.
* - `sigs_length` (1 byte): Number of signatures.
* - `sigs_array` (variable length): Array of signatures, each 65 bytes (RSV format).
*   - Each signature is composed of:
*     - `r` (32 bytes): R component of the signature.
*     - `s` (32 bytes): S component of the signature.
*     - `v` (1 byte): V component of the signature.
* - `updates_length` (1 byte): Number of updates.
* - `updates_array` (variable length): Array of updates, each following the Update ABI format.
*
* Update ABI Format: <merkle_proof_length><merkle_proof_array><timestamp><price><dataKey>
*
* Update Structure:
* - `merkle_proof_length` (1 byte): Length of the Merkle proof array.
* - `merkle_proof_array` (variable length): Array of Merkle proofs, each 32 bytes.
*   - `merkle_proof` (array of bytes32): Merkle proof elements.
*     - Each element is 32 bytes.
* - `timestamp` (32 bytes): Update timestamp (uint256).
* - `price` (32 bytes): Asset price (uint256).
* - `dataKey` (32 bytes): Asset data key (bytes32).
*
* Size Calculations:
* - `sigs_array` size: `sigs_length * 65` bytes.
* - `merkle_proof_array` size: `merkle_proof_length * 32` bytes.
* - `update` size: `(1 + merkle_proof_length * 32 + 32 + 32 + 32)` bytes.
* - Total updates size: `(1 + merkle_proof_length * 32 + 96) * updates_length` bytes.
 */
func (up *MerkleRootUpdateMultiple) ToCalldata() ([]byte, error) {
	// Build hex encoder
	updateBuffer := new(bytes.Buffer)

	// Encode merkle root
	merkleRoot := up.MerkleRoot
	if err := binary.Write(updateBuffer, binary.BigEndian, merkleRoot[:]); err != nil {
		return nil, err
	}

	// Encode signatures
	if err := binary.Write(updateBuffer, binary.BigEndian, []byte{byte(len(up.Signatures))}); err != nil {
		return nil, err
	}
	for _, sig := range up.Signatures {
		sigBytes, err := sig.ToBytes()
		if err != nil {
			return nil, err
		}
		if err := binary.Write(updateBuffer, binary.BigEndian, sigBytes); err != nil {
			return nil, err
		}

		// enc.Write(sig.R[:])
		// enc.Write(sig.S[:])
		// enc.Write([]byte{sig.V})
	}

	// Encode update data
	if err := binary.Write(updateBuffer, binary.BigEndian, []byte{byte(len(up.UpdateData))}); err != nil {
		return nil, err
	}
	for _, update := range up.UpdateData {
		// Encode merkle proof
		if err := binary.Write(updateBuffer, binary.BigEndian, []byte{byte(len(update.MerkleProof))}); err != nil {
			return nil, err
		}
		for _, proof := range update.MerkleProof {
			if err := binary.Write(updateBuffer, binary.BigEndian, proof[:]); err != nil {
				return nil, err
			}
		}

		// Encode update values
		if err := binary.Write(updateBuffer, binary.BigEndian, ethmath.U256Bytes(&update.Timestamp)); err != nil {
			return nil, err
		}
		if err := binary.Write(updateBuffer, binary.BigEndian, ethmath.U256Bytes(&update.Price)); err != nil {
			return nil, err
		}
		if err := binary.Write(updateBuffer, binary.BigEndian, update.DataKey[:]); err != nil {
			return nil, err
		}
	}

	updateCalldata := updateBuffer.Bytes()
	return updateCalldata, nil
}
