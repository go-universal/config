package config

import (
	"sync"

	"github.com/go-universal/cast"
)

// memory is an in-memory implementation of the Config interface.
// It stores configuration data in a thread-safe map.
type memory struct {
	data  map[string]any
	mutex sync.RWMutex
}

// NewMemory initializes a new in-memory Config instance.
// It accepts a map of configuration key-value pairs and populates the in-memory store.
// Returns the Config instance or an error if initialization fails.
func NewMemory(config map[string]any) (Config, error) {
	driver := &memory{
		data: make(map[string]any),
	}

	for k, v := range config {
		driver.Set(k, v)
	}

	return driver, nil
}

func (m *memory) Load() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.data == nil {
		m.data = make(map[string]any)
	}

	return nil
}

func (m *memory) Set(key string, value any) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data[key] = value
}

func (m *memory) Get(key string) any {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.data[key]
}

func (m *memory) Exists(key string) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	_, ok := m.data[key]
	return ok
}

func (m *memory) Cast(key string) cast.Caster {
	return cast.NewCaster(m.Get(key))
}
