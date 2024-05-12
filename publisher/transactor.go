package publisher

import (
	"math/big"
	"crypto/ecdsa"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/contrib/contracts/datafeeds/PullOracle"

	// "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type ITransactor interface {
	SendUpdate(update *MerkleRootUpdate) error
}

type Transactor struct {
	PullOracle *PullOracle.PullOracle
	// Private key of the signer that sends price updates
	PrivateKey *ecdsa.PrivateKey
	ChainID *big.Int
	opts *bind.TransactOpts
	client bind.ContractBackend
}

func NewTransactor(
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
		PullOracle: pullOracle,
		PrivateKey: privateKey,
		ChainID: chainID,
		client: client,
	}

	opts, err := transactor.createTransactOpts(chainID)
	if err != nil {
		return nil, err
	}
	transactor.opts = opts

	return transactor, nil
}

func (t *Transactor) createTransactOpts(chainID *big.Int) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		t.PrivateKey,
		chainID,
	)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

func (t *Transactor) SendUpdate(update *MerkleRootUpdate) error {

	// Remap to correct type...
	signatures := make([]PullOracle.PullOracleSignature, len(update.Signatures))
	for i, s := range update.Signatures {
		signatures[i] = PullOracle.PullOracleSignature{
			V: s.V,
			R: s.R,
			S: s.S,
		}
	}
	log.Infof("Sending update to PullOracle contract")

	// Send update to PullOracle contract
	tx, err := t.PullOracle.GetLastPrice(
		t.opts,
		update.NewMerkleRoot,
		update.MerkleProof,
		signatures,
		update.Price,
		update.Timestamp,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"update": update,
		}).Error("Failed to execute PullOracle.GetLastPrice")
		return err
	}

	log.WithFields(log.Fields{
		"tx": tx.Hash().Hex(),
	}).Info("Sent update to PullOracle contract")

	return nil
}
