package config

import (
	"fmt"
	"net/url"
	"errors"
	"regexp"

	"github.com/kelseyhightower/envconfig"
	"github.com/ethereum/go-ethereum/common"
)

type AppConfig struct {
	FinalizeSnapshotUrl string `envconfig:"FINALIZE_SNAPSHOT_URL" required:"true"`
	TargetChainUrl string `envconfig:"TARGET_CHAIN_URL" required:"true"`
	PullOracleAddress common.Address `envconfig:"PULL_ORACLE_ADDRESS" required:"true"`
	PrivateKey string `envconfig:"PRIVATE_KEY" required:"true"`
}

// This function verifies the symbolic correctness of the configuration. 
// This however does not guarantee the semantic correctness of the configuration.
func (config AppConfig) Verify() error {
	// Check if URLs are valid
	if _, err := url.ParseRequestURI(config.FinalizeSnapshotUrl); err != nil {
		return fmt.Errorf("invalid FinalizeSnapshotUrl: %w", err)
	}

	if _, err := url.ParseRequestURI(config.TargetChainUrl); err != nil {
		return fmt.Errorf("invalid TargetChainUrl: %w", err)
	}

	// Don't need to check common.Address
	// if common.IsHexAddress(config.PullOracleAddress.Hex()) == false { }

	// Check if PrivateKey is in the correct format (0x followed by 64 hex characters)
	re := regexp.MustCompile(`^0x[a-fA-F0-9]{64}$`)
	if !re.MatchString(config.PrivateKey) {
		return errors.New("invalid PrivateKey: must be 0x[0-9a-zA-Z]{64}")
	}

	return nil
}

// LoadConfig reads environment variables and initializes an AppConfig struct.
func LoadConfigFromEnv() (*AppConfig, error) {
	var config AppConfig
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}
	if err := config.Verify(); err != nil {
		return nil, err
	}
	return &config, nil
}
