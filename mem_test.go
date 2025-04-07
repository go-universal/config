package config_test

import (
	"testing"

	"github.com/go-universal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemory(t *testing.T) {
	data := map[string]any{"title": "My App"}
	cfg, err := config.NewMemory(data)
	require.NoError(t, err, "Failed to create memory config")

	v, err := cfg.Cast("title").String()
	require.NoError(t, err, "Failed to cast title to string")

	assert.Equal(t, "My App", v, `title should be "My App"`)
}
