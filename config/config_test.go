package config

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigWithCorrectEnv(t *testing.T) {
	// Setup environment variables
	finalizeSnapshotUrl := "http://localhost:3000"
	assert.Nil(t, os.Setenv("FINALIZE_SNAPSHOT_URL", finalizeSnapshotUrl))
	assert.Nil(t, os.Setenv("TARGET_CHAIN_URL", finalizeSnapshotUrl))
	assert.Nil(t, os.Setenv("PULL_ORACLE_ADDRESS", "0x5ca636af0aB140A75515Bd708E3e382aa7A70aEb"))
	assert.Nil(t, os.Setenv("PRIVATE_KEY", "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")) // 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

	// Load config
	config, err := LoadConfigFromEnv()
	assert.Nil(t, err)

	// Assert config values
	assert.Equal(t, config.FinalizeSnapshotUrl, finalizeSnapshotUrl)
}
