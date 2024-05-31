package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v3"
)

var regex = regexp.MustCompile(`^0x[a-fA-F0-9]{64}$`)

type AppConfig struct {
	FinalizeSnapshotURL string                   `yaml:"finalizeSnapshotUrl"`
	DataKeys            []string                 `yaml:"dataKeys"`
	Networks            map[string]NetworkConfig `yaml:"networks"`
	// Interval in seconds for publishing updates for DataKeys
	UpdateInterval      uint                     `yaml:"updateInterval"`
}

type NetworkConfig struct {
	TargetChainURL    string         `yaml:"targetChainUrl"`
	PullOracleAddress common.Address `yaml:"pullOracleAddress"`
	PrivateKey        string         `yaml:"privateKey"`
}

// This function verifies the symbolic correctness of the configuration.
// This however does not guarantee the semantic correctness of the configuration.
func (config AppConfig) Verify() error {
	if len(config.Networks) == 0 {
		return fmt.Errorf("networks are required")
	}

	if len(config.DataKeys) == 0 {
		return fmt.Errorf("data keys are required")
	}

	if (config.UpdateInterval == 0) {
		return fmt.Errorf("update interval is required")
	}

	// Check if URLs are valid
	if _, err := url.ParseRequestURI(config.FinalizeSnapshotURL); err != nil {
		return fmt.Errorf("invalid FinalizeSnapshotURL: %w", err)
	}

	for name, net := range config.Networks {
		if _, err := url.ParseRequestURI(net.TargetChainURL); err != nil {
			return fmt.Errorf("invalid %s TargetChainURL: %w", name, err)
		}

		// Don't need to check common.Address
		// if common.IsHexAddress(config.PullOracleAddress.Hex()) == false { }

		// Check if PrivateKey is in the correct format (0x followed by 64 hex characters)
		if !regex.MatchString(net.PrivateKey) {
			return errors.New("invalid PrivateKey: must be 0x[0-9a-zA-Z]{64}")
		}
	}

	return nil
}

func LoadConfig(configPath string) (*AppConfig, error) {
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	var conf AppConfig
	if err := yaml.Unmarshal(bytes, &conf); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if err := conf.Verify(); err != nil {
		return nil, fmt.Errorf("failed config verify: %w", err)
	}

	return &conf, nil
}
