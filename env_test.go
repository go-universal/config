package config_test

import (
	"os"
	"testing"

	"github.com/go-universal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnv(t *testing.T) {
	const testFile = "test.env"
	const testContent = `APP_TITLE="My App"`

	// Create test file
	err := os.WriteFile(testFile, []byte(testContent), os.ModePerm)
	require.NoError(t, err)
	defer func() {
		_ = os.Remove(testFile)
	}()

	// Load configuration
	cfg, err := config.NewEnv(testFile)
	if err != nil {
		t.Fatalf("failed to load env file: %v", err)
	}

	// Validate configuration
	v, err := cfg.Cast("APP_TITLE").String()
	if err != nil {
		t.Fatalf("failed to cast APP_TITLE: %v", err)
	}

	assert.Equal(t, "My App", v, `APP_TITLE should be "My App"`)
}
