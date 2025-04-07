package config

import (
	"os"
	"sync"

	"github.com/go-universal/cast"
	"github.com/joho/godotenv"
)

// env is a struct that manages environment variables and provides thread-safe access.
type env struct {
	files []string
	data  map[string]any
	mutex sync.RWMutex
}

// NewEnv initializes a new environment configuration loader.
// It optionally accepts file paths to load additional environment variables.
// Returns a Config instance or an error if loading fails.
func NewEnv(files ...string) (Config, error) {
	driver := &env{
		files: files,
		data:  make(map[string]any),
	}

	if err := driver.Load(); err != nil {
		return nil, err
	}

	return driver, nil
}

func (e *env) Load() error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return godotenv.Overload(e.files...)
}

func (e *env) Set(key string, value any) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.data[key] = value
}

func (e *env) Get(key string) any {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	if value, exists := e.data[key]; exists {
		return value
	}

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return nil
}

func (e *env) Exists(key string) bool {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	if _, exists := e.data[key]; exists {
		return true
	}

	if _, exists := os.LookupEnv(key); exists {
		return true
	}

	return false
}

func (e *env) Cast(key string) cast.Caster {
	return cast.NewCaster(e.Get(key))
}
