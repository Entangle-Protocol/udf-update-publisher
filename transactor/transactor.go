package transactor

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/contrib/contracts/datafeeds/PullOracle"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"

	// "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type ITransactor interface {
	SendUpdate(update *types.MerkleRootUpdate) error
}

type Transactor struct {
	ctx    context.Context
	client bind.ContractBackend
	// Private key of the signer that sends price updates
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
	opts       *bind.TransactOpts
	pullOracle *PullOracle.PullOracle
}

func NewTransactor(
	ctx context.Context,
	client bind.ContractBackend,
	privateKey *ecdsa.PrivateKey,
	chainID *big.Int,
	pullOracleAddress common.Address,
) (*Transactor, error) {
	pullOracle, err := PullOracle.NewPullOracle(pullOracleAddress, client)
	if err != nil {
		return nil, err
	}

	transactor := &Transactor{
		client:     client,
		ctx:        ctx,
		privateKey: privateKey,
		chainID:    chainID,
		pullOracle: pullOracle,
	}

	opts, err := transactor.createTransactOpts(chainID)
	if err != nil {
		return nil, err
	}
	transactor.opts = opts

	return transactor, nil
}

func (t *Transactor) SendUpdate(update *types.MerkleRootUpdate) error {
	// Remap to correct type...
	signatures := make([]PullOracle.PullOracleSignature, len(update.Signatures))
	for i, s := range update.Signatures {
		signatures[i] = PullOracle.PullOracleSignature{
			R: s.R,
			S: s.S,
			V: s.V,
		}
	}

	log.Infof("Sending update to PullOracle contract")

	for {
		// Send update to PullOracle contract
		tx, err := t.pullOracle.GetLastPrice(
			t.opts,
			update.NewMerkleRoot,
			update.MerkleProof,
			signatures,
			update.DataKey,
			update.Price,
			update.Timestamp,
		)
		if err != nil {
			if strings.Contains(err.Error(), "invalid nonce") {
				// If the error is regarding nonce, fetch correct nonce and retry
				log.WithFields(log.Fields{
					"error": err,
					"nonce": t.opts.Nonce,
				}).Warn("Transactor: Invalid nonce, fetching correct nonce and retrying...")

				nonce, err := t.client.PendingNonceAt(
					t.ctx,
					t.opts.From,
				)
				if err != nil {
					log.WithError(err).Error("Transactor: Failed to fetch correct nonce")
					return err
				}

				log.WithField("nonce", nonce).Info("Transactor: Fetched correct nonce")
				t.opts.Nonce.Set(big.NewInt(int64(nonce)))

				continue
			}

			log.WithFields(log.Fields{
				"error":  err,
				"update": update,
			}).Error("Failed to execute PullOracle.GetLastPrice")

			return err
		}

		log.WithFields(log.Fields{
			"tx": tx.Hash().Hex(),
		}).Info("Sent update to PullOracle contract")

		break
	}

	t.opts.Nonce.Add(t.opts.Nonce, big.NewInt(1))

	return nil
}

func (t *Transactor) createTransactOpts(chainID *big.Int) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		t.privateKey,
		chainID,
	)
	if err != nil {
		return nil, err
	}

	// Get nonce
	nonce, err := t.client.PendingNonceAt(
		t.ctx,
		opts.From,
	)
	if err != nil {
		return nil, err
	}

	opts.Nonce = big.NewInt(int64(nonce))
	// opts.GasLimit = 500000

	return opts, nil
}
