package transactor

import (
	"math/big"
	"runtime"
	"sync"
	"unsafe"

	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/utils"
)

/*
#cgo LDFLAGS: -L${SRCDIR}/.. -lprice_publisher
#include <stdint.h>
#include <string.h>

typedef struct {
    uint8_t data_key[32];
    const uint8_t * merkle_proof;
    size_t merkle_proof_len;
    uint8_t data[32];
    uint64_t timestamp;
} MultipleUpdateData;

typedef struct {
    uint8_t r[32];
    uint8_t s[32];
    uint8_t v;
} ECDSASignature;

typedef struct {
    uint8_t merkle_root[32];
    const ECDSASignature* signatures;
    size_t signatures_len;
    const MultipleUpdateData* update_data;
    size_t update_data_len;
} MerkleRootUpdateMultiple;

typedef struct {
    uint8_t data[32];
    uint64_t timestamp;
} LatestUpdate;

void update_multiple_assets(const MerkleRootUpdateMultiple * merkle_root_update_multiple);
LatestUpdate get_latest_update(const uint8_t * data_key);
void get_chain_id(uint8_t * chain_id_ptr);
*/
import "C"

type SolanaTransactor struct {
	mutex sync.Mutex
}

func NewSolanaTransactor() (*SolanaTransactor, error) {
	transactor := &SolanaTransactor{
		mutex: sync.Mutex{},
	}
	return transactor, nil
}

func (t *SolanaTransactor) SendMultipleUpdate(update *types.MerkleRootUpdateMultiple) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	var pinner runtime.Pinner
	defer pinner.Unpin()

	log.WithFields(log.Fields{
		"merkleRoot": ethcommon.Bytes2Hex(update.MerkleRoot[:]),
		"chainID":    t.ChainID().String(),
		"dataKeys": utils.Map(update.UpdateData, func(u types.MultipleUpdateData) string {
			return ethcommon.Bytes2Hex(u.DataKey[:])
		}),
	}).Info("Sending PullOracle.UpdateMultipleAssets tx")

	cMerkleRootUpdate, err := cUpdateMultipleFrom(*update, &pinner)
	if err != nil {
		log.Errorf("Failed to make cUpdateMultipleFrom update: %v\n", err)
		return err
	}
	C.update_multiple_assets(&cMerkleRootUpdate)
	return nil
}

func (t *SolanaTransactor) SendUpdate(update *types.MerkleRootUpdate) error {
	log.Errorf("Sending update to solana is deprecated")
	return nil
}

func (t *SolanaTransactor) LatestUpdate(dataKey [32]byte) (*big.Int, *big.Int) {
	latestUpdate := C.get_latest_update((*C.uint8_t)(&dataKey[0]))
	data := make([]byte, 32)
	C.memcpy(unsafe.Pointer(&data[0]), unsafe.Pointer(&latestUpdate.data[0]), 32)
	dataInt := big.NewInt(0)
	dataInt.SetBytes(data)

	timestamp := latestUpdate.timestamp
	timeInt := big.NewInt((int64)(timestamp))
	log.Infof("Latest update: %d, timestamp: %d", dataInt.Uint64(), timeInt.Uint64())
	return dataInt, timeInt
}

func (t *SolanaTransactor) ChainID() *big.Int {
	chainIdData := make([]C.uint8_t, 16)
	C.get_chain_id(&chainIdData[0])
	result := make([]byte, 16)
	C.memcpy(unsafe.Pointer(&result[0]), unsafe.Pointer(&chainIdData[0]), 16)
	chainId := big.NewInt(0).SetBytes(result)
	return chainId
}

func cUpdateMultipleFrom(updateMultiple types.MerkleRootUpdateMultiple, pinner *runtime.Pinner) (C.MerkleRootUpdateMultiple, error) {
	var cUpdateData []C.MultipleUpdateData

	for _, dataFeed := range updateMultiple.UpdateData {

		cMerkleProof := make([][32]C.uint8_t, len(dataFeed.MerkleProof))
		for i, proof := range dataFeed.MerkleProof {
			C.memcpy(unsafe.Pointer(&cMerkleProof[i]), unsafe.Pointer(&proof[0]), 32)
		}
		pinner.Pin(&cMerkleProof[0])

		cDataFeed := C.MultipleUpdateData{
			timestamp:        C.uint64_t(dataFeed.Timestamp.Uint64()),
			merkle_proof:     (*C.uint8_t)(unsafe.Pointer(&cMerkleProof[0])),
			merkle_proof_len: C.size_t(len(cMerkleProof)),
		}
		C.memcpy(unsafe.Pointer(&cDataFeed.data_key[0]), unsafe.Pointer(&dataFeed.DataKey[0]), 32)

		data := make([]byte, 32)
		dataFeed.Price.FillBytes(data)
		C.memcpy(unsafe.Pointer(&cDataFeed.data[0]), unsafe.Pointer(&data[0]), 32)
		cUpdateData = append(cUpdateData, cDataFeed)
	}

	pinner.Pin(&updateMultiple.Signatures[0])
	pinner.Pin(&cUpdateData[0])

	var cMerkleRoot [32]C.uint8_t
	C.memcpy(unsafe.Pointer(&cMerkleRoot[0]), unsafe.Pointer(&updateMultiple.MerkleRoot[0]), 32)

	cMerkleRootUpdate := C.MerkleRootUpdateMultiple{
		merkle_root:     cMerkleRoot,
		signatures:      (*C.ECDSASignature)(unsafe.Pointer(&updateMultiple.Signatures[0])),
		signatures_len:  C.size_t(len(updateMultiple.Signatures)),
		update_data:     &cUpdateData[0],
		update_data_len: C.size_t(len(cUpdateData)),
	}

	return cMerkleRootUpdate, nil
}
