package publisher

import (
	"gitlab.ent-dx.com/entangle/pull-update-publisher/contrib/contracts/datafeeds/PullOracle"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type ITransactor interface {
	SendUpdate() error
}

type Transactor struct {
	PullOracle *PullOracle.PullOracle
}

func NewTransactor(client *ethclient.Client, pullOracleAddress common.Address) (*Transactor, error) {
	pullOracle, err := PullOracle.NewPullOracle(pullOracleAddress, client)
	if err != nil {
		return nil, err
	}

	return &Transactor{
		PullOracle: pullOracle,
	}, nil
}

func (t *Transactor) SendUpdate() error {
	log.Infof("Sending update to PullOracle contract")
	// Send update to PullOracle contract
	return nil
}
