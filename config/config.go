package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/ethereum/go-ethereum/common"
)

type AppConfig struct {
	FinalizeSnapshotUrl string `envconfig:"FINALIZE_SNAPSHOT_URL" required:"true"`
	TargetChainUrl string `envconfig:"TARGET_CHAIN_URL" required:"true"`
	PullOracleAddress common.Address `envconfig:"PULL_ORACLE_ADDRESS" required:"true"`
}

// LoadConfig reads environment variables and initializes an AppConfig struct.
func LoadConfigFromEnv() (*AppConfig, error) {
	var config AppConfig
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
