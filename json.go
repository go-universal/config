package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-universal/cast"
	"github.com/tidwall/gjson"
)

// json is a driver for managing JSON-based configuration.
type json struct {
	files []string
	raw   string
	data  map[string]any
	mutex sync.RWMutex
}

// NewJSON initializes a new Config instance that loads configuration from JSON files.
// Accepts a list of file paths and returns a Config instance or an error if loading fails.
func NewJSON(files ...string) (Config, error) {
	driver := &json{
		files: files,
		data:  make(map[string]any),
	}

	if err := driver.Load(); err != nil {
		return nil, err
	}

	return driver, nil
}

func (j *json) Load() error {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	var contents []string

	for _, file := range j.files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", file, err)
		}

		content := string(bytes)
		if !gjson.Valid(content) {
			return fmt.Errorf("invalid JSON in file %s", file)
		}

		if len(j.files) > 1 {
			fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
			contents = append(contents, fmt.Sprintf(`"%s":%s`, fileName, content))
		} else {
			contents = append(contents, content)
		}
	}

	if len(j.files) > 1 {
		j.raw = "{" + strings.Join(contents, ",") + "}"
	} else if len(contents) > 0 {
		j.raw = contents[0]
	} else {
		j.raw = "{}"
	}

	return nil
}

func (j *json) Set(key string, value any) {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	j.data[key] = value
}

func (j *json) Get(key string) any {
	j.mutex.RLock()
	defer j.mutex.RUnlock()

	if value, exists := j.data[key]; exists {
		return value
	}

	if value := gjson.Get(j.raw, key); value.Exists() {
		return value.Value()
	}

	return nil
}

func (j *json) Exists(key string) bool {
	j.mutex.RLock()
	defer j.mutex.RUnlock()

	if _, exists := j.data[key]; exists {
		return true
	}

	return gjson.Get(j.raw, key).Exists()
}

func (j *json) Cast(key string) cast.Caster {
	return cast.NewCaster(j.Get(key))
}
