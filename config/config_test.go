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

	// Load config
	config, err := LoadConfigFromEnv()
	assert.Nil(t, err)

	// Assert config values
	assert.Equal(t, config.FinalizeSnapshotUrl, finalizeSnapshotUrl)
}
