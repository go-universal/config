package config_test

import (
	"os"
	"testing"

	"github.com/go-universal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJsonSingle(t *testing.T) {
	const testFile = "test.json"
	const testContent = `{ "app": { "title": "My App" } }`

	// Create test file
	err := os.WriteFile(testFile, []byte(testContent), 0644)
	require.NoError(t, err)
	defer func() {
		_ = os.Remove(testFile)
	}()

	// Load configuration
	cfg, err := config.NewJSON(testFile)
	require.NoError(t, err)

	// Validate configuration value
	v, err := cfg.Cast("app.title").String()
	require.NoError(t, err)
	assert.Equal(t, "My App", v, `"app.title" should be "My App"`)
}

func TestJsonMultiple(t *testing.T) {
	const appFile = "app.json"
	const appContent = `{ "title": "My App" }`
	const dbFile = "db.json"
	const dbContent = `{ "name": "my_app" }`

	// Create app file
	err := os.WriteFile(appFile, []byte(appContent), 0644)
	require.NoError(t, err)
	defer func() {
		_ = os.Remove(appFile)
	}()

	// Create db file
	err = os.WriteFile(dbFile, []byte(dbContent), 0644)
	require.NoError(t, err)
	defer func() {
		_ = os.Remove(dbFile)
	}()

	// Load configuration
	cfg, err := config.NewJSON(appFile, dbFile)
	require.NoError(t, err)

	// Validate configuration value
	v, err := cfg.Cast("app.title").String()
	require.NoError(t, err)
	assert.Equal(t, "My App", v, `"app.title" should be "My App"`)

	v, err = cfg.Cast("db.name").String()
	require.NoError(t, err)
	assert.Equal(t, "my_app", v, `"db.name" should be "my_app"`)
}
