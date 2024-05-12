package config

import (
	"os"
	"testing"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/ethereum/go-ethereum/common"
)

func TestLoadConfigWithCorrectEnv(t *testing.T) {
	// Setup environment variables
	// finalizeSnapshotUrl := "http://localhost:3000"
	finalizeSnapshotUrl := gofakeit.URL()
	targetChainUrl := gofakeit.URL()
	pullOracleAddress :=  "0x5ca636af0aB140A75515Bd708E3e382aa7A70aEb"
	privateKey := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" // 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
	assert.Nil(t, os.Setenv("FINALIZE_SNAPSHOT_URL", finalizeSnapshotUrl))
	assert.Nil(t, os.Setenv("TARGET_CHAIN_URL", targetChainUrl))
	assert.Nil(t, os.Setenv("PULL_ORACLE_ADDRESS", pullOracleAddress))
	assert.Nil(t, os.Setenv("PRIVATE_KEY", privateKey))

	// Load config
	config, err := LoadConfigFromEnv()
	assert.Nil(t, err)

	err = config.Verify()
	assert.Nil(t, err)

	// Assert config values
	assert.Equal(t, config.FinalizeSnapshotUrl, finalizeSnapshotUrl)
	assert.Equal(t, config.TargetChainUrl, targetChainUrl)
	assert.Equal(t, config.PullOracleAddress, common.HexToAddress(pullOracleAddress))
	assert.Equal(t, config.PrivateKey, privateKey)
}
