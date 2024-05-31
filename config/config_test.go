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
	updateInterval := 30
	targetChainUrl := gofakeit.URL()
	pullOracleAddress := "0x9EeF2FA023ADbfe260EC8164BAfB454ffEF3E2bd"
	privateKey := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" // 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

	path := filepath.Join(os.TempDir(), "config.yaml")
	payload := fmt.Sprintf(template,
		finalizeSnapshotUrl,
		assetKey,
		updateInterval,
		network,
		targetChainUrl,
		pullOracleAddress,
		privateKey,
	)
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

	// Assert config values
	r.Equal(finalizeSnapshotUrl, config.FinalizeSnapshotURL)
	r.Equal(assetKey, config.DataKeys[0])
	r.Equal(targetChainUrl, config.Networks[network].TargetChainURL)
	r.Equal(pullOracleAddress, config.Networks[network].PullOracleAddress.String())
	r.Equal(privateKey, config.Networks[network].PrivateKey)
}

var template = `
finalizeSnapshotUrl: %s
dataKeys:
  - %s
updateInterval: %d
networks:
  %s:
    targetChainUrl: %s
    pullOracleAddress: %s
    privateKey: %s
`
