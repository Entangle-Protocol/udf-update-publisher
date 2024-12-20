package config_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/config"
)

func TestLoadConfigWithCorrectEnv(t *testing.T) {
	r := require.New(t)

	// Setup environment variables
	// finalizeSnapshotUrl := "http://localhost:3000"
	network := "ethereum"
	finalizeSnapshotUrl := gofakeit.URL()
	assetKey := "NGL/USDT"
	sourceID := "prices-feed"
	updateInterval := 30
	updateThreshold := "5m0s"
	priceDiffThreshold := 1
	targetChainUrl := gofakeit.URL()
	pullOracleAddress := "0x9EeF2FA023ADbfe260EC8164BAfB454ffEF3E2bd"
	privateKey := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" // 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

	path := filepath.Join(os.TempDir(), "config.yaml")
	payload := fmt.Sprintf(template,
		finalizeSnapshotUrl,
		assetKey,
		sourceID,
		assetKey,
		updateInterval,
		updateThreshold,
		priceDiffThreshold,
		network,
		targetChainUrl,
		pullOracleAddress,
		privateKey,
	)
	fmt.Printf("payload: %s\n", payload)
	err := os.WriteFile(path, []byte(payload), os.ModePerm)
	r.NoError(err)

	// Load config
	config, err := config.LoadConfig(path)
	r.NoError(err)

	err = config.Verify()
	r.NoError(err)

	r.Len(config.Networks, 1)
	r.NotEmpty(config.Networks[network])
	r.NotEmpty(config.DataKeys)

	fmt.Printf("config: %+v\n", *config)

	// Assert config values
	r.Equal(finalizeSnapshotUrl, config.FinalizeSnapshotURL)
	r.Equal(assetKey, config.DataKeys[0])
	r.Equal(targetChainUrl, config.Networks[network].TargetChainURL)
	r.Equal(pullOracleAddress, config.Networks[network].PullOracleAddress.String())
	r.Equal(privateKey, config.Networks[network].PrivateKey)
	r.Equal(updateInterval, int(config.Publisher.UpdateInterval))
	r.Equal(updateThreshold, config.Publisher.UpdateThreshold.String())
	r.Equal(priceDiffThreshold, int(config.Publisher.PriceDiffThreshold))
}

var template = `
finalizeSnapshotUrl: %s
dataKeys:
  - %s
assets:
  - sourceID: %s
    dataKeys:
    - %s
publisher:
  updateInterval: %d 
  updateThreshold: %s 
  priceDiffThreshold: %d
networks:
  %s:
    targetChainUrl: %s
    pullOracleAddress: %s
    privateKey: %s
`
