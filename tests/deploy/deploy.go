package deploy

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethsim "github.com/ethereum/go-ethereum/ethclient/simulated"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/contrib/contracts/datafeeds/PullOracle"
	abiContracts "gitlab.ent-dx.com/entangle/pull-update-publisher/tests/testdata/abi"
)

type ProtocolConfig struct {
	ID            [32]byte
	ConsensusRate *big.Int
}

type DeployedContracts struct {
	EndPoint   common.Address
	PullOracle common.Address
}

func DeployContracts(ctx context.Context, backend *ethsim.Backend, privKey *ecdsa.PrivateKey, protoCfg ProtocolConfig) (DeployedContracts, error) {
	client := backend.Client()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return DeployedContracts{}, err
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(
		privKey,
		chainID,
	)
	if err != nil {
		return DeployedContracts{}, err
	}

	admin := txOpts.From

	endPointAddress, _, endPoint, err := abiContracts.DeployMockEndPoint(txOpts, client)
	if err != nil {
		return DeployedContracts{}, err
	}
	backend.Commit()

	_, err = endPoint.AddAllowedProtocol(txOpts, protoCfg.ID, protoCfg.ConsensusRate, []common.Address{admin})
	if err != nil {
		return DeployedContracts{}, err
	}
	backend.Commit()

	pullOracleAddress, _, pullOracle, err := PullOracle.DeployPullOracle(txOpts, client)
	if err != nil {
		return DeployedContracts{}, err
	}
	backend.Commit()

	_, err = pullOracle.Initialize(txOpts, protoCfg.ID, endPointAddress)
	if err != nil {
		return DeployedContracts{}, err
	}
	backend.Commit()

	gotEndPointAddress, err := pullOracle.EndPoint(&bind.CallOpts{})
	if err != nil {
		return DeployedContracts{}, err
	}

	if endPointAddress.String() != gotEndPointAddress.String() {
		return DeployedContracts{}, fmt.Errorf("diff end point addr")
	}

	return DeployedContracts{
		EndPoint:   endPointAddress,
		PullOracle: pullOracleAddress,
	}, nil
}
