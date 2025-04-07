package config

import "github.com/go-universal/cast"

// Config provide methods for loading, setting, getting,
// checking existence, and casting configuration values.
type Config interface {
	// Load loads the configuration from a source.
	Load() error

	// Set sets/override the value for a given key in the configuration.
	Set(key string, value any)

	// Get retrieves the value associated with the given key.
	// on json driver if multiple file passed you must append
	// file name to access setting .e.g. file1.some.field
	Get(key string) any

	// Exists checks if a given key exists in the configuration.
	Exists(key string) bool

	// Cast retrieves the value associated with the given key and returns it
	// as a cast.Caster
	Cast(key string) cast.Caster
}
