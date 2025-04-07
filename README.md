# Config Library

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/config?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/config.svg)](https://pkg.go.dev/github.com/go-universal/config)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/config/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/config)](https://goreportcard.com/report/github.com/go-universal/config)
![Contributors](https://img.shields.io/github/contributors/go-universal/config)
![Issues](https://img.shields.io/github/issues/go-universal/config)

The `config` library offers a unified, thread-safe interface for managing configuration data from multiple sources, including in-memory maps, JSON files, and environment variables. It also supports type casting for configuration values.

## Installation

```bash
go get github.com/go-universal/config
```

## Features

- Unified configuration management from in-memory maps, JSON files, and environment variables.
- Thread-safe operations for accessing and modifying configuration data.
- Built-in type casting for configuration values using `cast`.

## API Overview

### `Config` Interface

The `Config` interface provides methods for managing configuration data:

- **`Load() error`**: Loads configuration from the source.
- **`Set(key string, value any)`**: Sets or updates a value for the specified key.
- **`Get(key string) any`**: Retrieves the value for the specified key.
- **`Exists(key string) bool`**: Checks if a key exists in the configuration.
- **`Cast(key string) cast.Caster`**: Retrieves and casts the value for the specified key.

### In-Memory Configuration

Create an in-memory configuration instance:

```go
NewMemory(config map[string]any) (Config, error)
```

```go
import (
    "fmt"

    "github.com/go-universal/config"
)

func main() {
    data := map[string]any{"app_name": "My App"}
    cfg, err := config.NewMemory(data)
    if err != nil {
        panic(err)
    }

    fmt.Println(cfg.Get("app_name")) // Output: My App
}
```

### JSON Configuration

Creates a new configuration instance that loads data from JSON files. When multiple files are loaded, the configuration keys are prefixed with the respective file name, forming a path like `{file}.path.to.key`.

```go
NewJSON(files ...string) (Config, error)
```

```go
import (
    "fmt"

    "github.com/go-universal/config"
)

func main() {
    cfg, err := config.NewJSON("app.json", "db.json")
    if err != nil {
        panic(err)
    }

    fmt.Println(cfg.Get("app.title")) // Output: My App
}
```

### Environment Variable Configuration

Creates a new configuration instance that loads data from environment variables. Optionally, `.env` files can be loaded.

```go
NewEnv(files ...string) (Config, error)
```

```go
import (
    "fmt"

    "github.com/go-universal/config"
)

func main() {
    cfg, err := config.NewEnv(".env")
    if err != nil {
        panic(err)
    }

    fmt.Println(cfg.Get("APP_TITLE")) // Output: My App
}
```

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.
