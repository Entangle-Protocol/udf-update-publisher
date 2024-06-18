package publisher

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/config"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/transactor"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"
)

type UpdatePublisher struct {
	dataKeys    []string
	assets      []config.AssetSet
	transactors []*updateTransactor
	fetcher     fetcher.IFetcher
}

type updateTransactor struct {
	transactor.ITransactor

	priceDiffThreshold *big.Int
	updateThreshold    *big.Int // as unix timestamp

	mutex    *sync.Mutex
	dataKeys map[common.Hash]*latestUpdate
}

type latestUpdate struct {
	price     *big.Int
	updatedAt *big.Int // as unix timestamp
}

func NewUpdatePublisher(
	conf config.PublisherConfig,
	transactors []transactor.ITransactor,
	fetcher fetcher.IFetcher,
	dataKeys []string,
	assets []config.AssetSet,
) *UpdatePublisher {
	t := make([]*updateTransactor, 0, len(transactors))
	for _, v := range transactors {
		t = append(t, newUpdateTransactor(conf, v))
	}

	return &UpdatePublisher{
		dataKeys:    dataKeys,
		assets:      assets,
		transactors: t,
		fetcher:     fetcher,
	}
}

func (up *UpdatePublisher) PublishUpdate(ctx context.Context) error {
	for _, dataKey := range up.dataKeys {
		proofs, err := up.fetcher.GetFeedProofs(ctx, dataKey)
		if err != nil {
			log.Errorf("Failed to get feed proofs: %v", err)
			continue
		}
		log.Debugf("Got feed proofs: %+v", proofs)

		update, err := NewMerkleUpdateFromProof(proofs)
		if err != nil {
			log.Errorf("Failed to create merkle update from proof: %v", err)
			continue
		}

		for _, transactor := range up.transactors {
			if err := transactor.sendUpdate(update); err != nil {
				log.Errorf("Failed to send update: %v", err)
				continue
			}
		}

		log.Infof("Successfully updated datakey: %s", proofs.Key)
	}

	return nil
}

func (up *UpdatePublisher) PublishMultipleUpdate(ctx context.Context) error {
	// Publish valid updates for each asset set
	for _, asset := range up.assets {
		feedsProofs, err := up.fetcher.GetSpotterFeedsProofs(ctx, asset.SourceID, asset.DataKeys)
		if err != nil {
			log.Errorf("Failed to get feeds proofs: %v", err)
			continue
		}

		merkleUpdates := make([]*types.MerkleRootUpdate, 0)

		for _, proof := range feedsProofs.Proofs() {
			log.Debugf("Got feed proofs: %+v", proof)

			update, err := NewMerkleUpdateFromProof(proof)
			if err != nil {
				log.Errorf("Failed to create merkle update from proof: %v", err)
				continue
			}

			merkleUpdates = append(merkleUpdates, update)
		}

		for _, transactor := range up.transactors {
			// Select only valid updates
			validUpdates := make([]*types.MerkleRootUpdate, 0)
			for _, update := range merkleUpdates {
				if err := transactor.validateUpdate(update); err == nil {
					validUpdates = append(validUpdates, update)
				}
			}

			// If no valid update, skip sending to this network
			if len(validUpdates) == 0 {
				log.Debugf("No valid updates for chainID %d", transactor.ChainID())
				continue
			}

			multipleUpdate, err := types.NewMekrleRootUpdateMultipleFromUpdates(validUpdates)
			if err != nil {
				log.WithFields(log.Fields{
					"error":   err,
					"chainID": transactor.ChainID(),
				}).Errorf("Failed to create multiple update object")
				continue
			}

			if err := transactor.sendMultipleUpdate(multipleUpdate); err != nil {
				log.Errorf("Failed to send multiple updates: %v", err)
				continue
			}
		}
	}

	return nil
}

func newUpdateTransactor(conf config.PublisherConfig, t transactor.ITransactor) *updateTransactor {
	return &updateTransactor{
		ITransactor:        t,
		updateThreshold:    big.NewInt(int64(conf.UpdateThreshold.Seconds())),
		priceDiffThreshold: big.NewInt(int64(conf.PriceDiffThreshold)),
		mutex:              &sync.Mutex{},
		dataKeys:           make(map[common.Hash]*latestUpdate),
	}
}

func (t *updateTransactor) sendMultipleUpdate(update *types.MerkleRootUpdateMultiple) error {
	if _, err := t.SendMultipleUpdate(update); err != nil {
		return fmt.Errorf("failed to send multiple update: %w", err)
	}

	t.mutex.Lock()
	for _, u := range update.UpdateData {
		dataKey := common.BytesToHash(u.DataKey[:])
		t.dataKeys[dataKey] = &latestUpdate{
			price:     new(big.Int).Set(u.Price),
			updatedAt: new(big.Int).Set(u.Timestamp),
		}
	}
	t.mutex.Unlock()
	return nil
}

func (t *updateTransactor) sendUpdate(update *types.MerkleRootUpdate) error {
	if err := t.validateUpdate(update); err != nil {
		return fmt.Errorf("skip update: %w", err)
	}

	if _, err := t.SendUpdate(update); err != nil {
		return fmt.Errorf("failed to send update: %w", err)
	}

	dataKey := common.BytesToHash(update.DataKey[:])

	t.mutex.Lock()
	t.dataKeys[dataKey] = &latestUpdate{
		price:     new(big.Int).Set(update.Price),
		updatedAt: new(big.Int).Set(update.Timestamp),
	}
	t.mutex.Unlock()

	return nil
}

func (t *updateTransactor) validateUpdate(update *types.MerkleRootUpdate) error {
	if update.Price.Cmp(common.Big0) <= 0 {
		return fmt.Errorf("invalid update price")
	}

	if update.Timestamp.Int64() <= 0 {
		return fmt.Errorf("invalid update timestamp")
	}

	dataKey := common.BytesToHash(update.DataKey[:])

	t.mutex.Lock()
	defer t.mutex.Unlock()

	latest, ok := t.dataKeys[dataKey]
	if !ok {
		// if data key not exists in map, fetch info from transactor
		price, timestamp := t.LatestUpdate(update.DataKey)
		latest = &latestUpdate{
			price:     new(big.Int).Set(price),
			updatedAt: new(big.Int).Set(timestamp),
		}
	}

	since := new(big.Int).Sub(update.Timestamp, latest.updatedAt)

	diff := new(big.Int).Sub(update.Price, latest.price)                                  // diff = new - old
	percent := new(big.Int).Quo(new(big.Int).Mul(diff, big.NewInt(10_000)), update.Price) // percent = (diff * 10_000) / new

	if percent.CmpAbs(t.priceDiffThreshold) < 0 && since.Cmp(t.updateThreshold) < 0 {
		return fmt.Errorf("too often update with same price, diff=%s previous=%s got=%s", percent, latest.updatedAt, latest.updatedAt)
	}

	return nil
}
