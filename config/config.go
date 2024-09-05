package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v3"
)

var (
	defaultPriceDiffThreshold = 100
	defaultUpdateThreshold    = 5 * time.Minute
	defaultUpdateInterval     = 30
)

var regex = regexp.MustCompile(`^0x[a-fA-F0-9]{64}$`)

type AssetSet struct {
	SourceID string   `yaml:"sourceID"`
	DataKeys []string `yaml:"dataKeys"`
}

type AppConfig struct {
	FinalizeSnapshotURL string                   `yaml:"finalizeSnapshotUrl"`
	DataKeys            []string                 `yaml:"dataKeys"`
	Assets              []AssetSet               `yaml:"assets"`
	Networks            map[string]NetworkConfig `yaml:"networks"`
	Publisher           PublisherConfig          `yaml:"publisher"`
}

type PublisherConfig struct {
	// Price diff threshold in percents where (1% = 100), below which an update will not be published
	PriceDiffThreshold uint `yaml:"priceDiffThreshold"`
	// Interval in seconds for publishing updates for DataKeys
	UpdateInterval uint `yaml:"updateInterval"`
	// Threshold in time duration, below which an update will not be published
	UpdateThreshold time.Duration `yaml:"updateThreshold"`
}

type NetworkConfig struct {
	Type              string         `yaml:"type"`
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

	if config.Publisher.UpdateInterval == 0 {
		return fmt.Errorf("update interval is required")
	}

	if config.Publisher.PriceDiffThreshold == 0 {
		config.Publisher.PriceDiffThreshold = uint(defaultPriceDiffThreshold)
	}

	if int64(config.Publisher.UpdateThreshold.Seconds()) == 0 {
		config.Publisher.UpdateThreshold = defaultUpdateThreshold
	}

	// Check if URLs are valid
	if _, err := url.ParseRequestURI(config.FinalizeSnapshotURL); err != nil {
		return fmt.Errorf("invalid FinalizeSnapshotURL: %w", err)
	}

	for name, net := range config.Networks {
		if net.Type == "nonevm" {
			continue
		}
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
