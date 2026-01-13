
# simplejrpc-go

`simplejrpc-go` is a lightweight JSON-RPC framework for Go, designed to help developers build high-performance and maintainable JSON-RPC services with essential built-in features.

## Features
- **Configuration Management**: Load and access nested configuration from JSON files using dot notation
- **Structured Logging**: Powered by `go.uber.org/zap` with log rotation, stdout output, and customizable log levels
- **Data Validation**: Comprehensive struct field validation with rules for required fields, length, numeric ranges, etc.
- **i18n Support**: Internationalization resource loading (currently INI file format)
- **Network Communication**: JSON-RPC server implementation with middleware support

## Installation
```sh
go get github.com/DemonZack/simplejrpc-go@latest
```

## Configuration Management

Example `config.json`:
```json
{
    "test": {
        "version": "1.0.0",
        "jsonrpc": {
            "sockets": "rpc.sock"
        },
        "logger": {
            "path": "logs/",
            "file": "{Y-m-d}.log",
            "level": "error",
            "stdout": false,
            "StStatus": 0,
            "rotateBackupLimit": 7,
            "writerColorEnable": true,
            "RotateBackupCompress": 9,
            "rotateExpire": "1d",
            "Flag": 44
        }
    }
}
```

Usage example:
```go
package main

import (
	"fmt"
	"path/filepath"
	"github.com/DemonZack/simplejrpc-go/core"
	"github.com/DemonZack/simplejrpc-go/core/config"
	"github.com/DemonZack/simplejrpc-go/os/gpath"
)

func main() {
	env := "test"
	fullPath, _ := filepath.Abs(filepath.Dir("."))
	gpath.GmCfgPath = filepath.Join(filepath.Dir(fullPath), "..")
	core.InitContainer(config.WithConfigEnvFormatterOptionFunc(env))

	val, err := core.Container.CfgFmt().GetValue("logger.level").String()
	if err != nil {
		panic(err)
	}
	fmt.Println("[*] logger.level:", val)
}
```

## Logging

Example logging setup:
```go
package main

import (
	"fmt"
	"time"
	"go.uber.org/zap"
	"github.com/DemonZack/simplejrpc-go/core/glog"
)

func main() {
	config := map[string]any{
		"path":                 "logs/",
		"file":                 "{Y-m-d}.log",
		"level":                "error",
		"stdout":               false,
		"rotateBackupLimit":    7,
		"writerColorEnable":    true,
		"RotateBackupCompress": 9,
		"rotateExpire":         "1d"
	}

	logger, err := glog.NewLogger(config)
	if err != nil {
		panic(fmt.Sprintf("init logger failed: %v", err))
	}
	defer logger.Sync()

	logger.Info("Logger initialized",
		zap.String("path", config["path"].(string)),
		zap.Int("backupLimit", config["rotateBackupLimit"].(int)),
	)
}
```

## Data Validation

Validation example:
```go
package main

import (
	"fmt"
	"github.com/DemonZack/simplejrpc-go/core"
)

type User struct {
	Username string `validate:"min_length:6#Username must be at least 6 characters"`
	Age      any    `validate:"required|range:18,100|int"`
	Email    string `validate:"required"`
}

func main() {
	user := User{
		Username: "test",
		Age:      15,
	}

	err := core.Container.Valid().Walk(&user)
	if err != nil {
		fmt.Println("Validation error:", err)
	}
}
```

## Internationalization (i18n)

i18n example:
```go
package main

import (
	"path/filepath"
	"github.com/DemonZack/simplejrpc-go/core/gi18n"
)

func main() {
	path := filepath.Join("testdata", "i18n")
	gi18n.Instance().SetPath(path)
	gi18n.Instance().SetLanguage(gi18n.English)
	
	println(gi18n.Instance().T("Welcome"))
}
```

## JSON-RPC Server

Server implementation:
```go
package main

import (
	"github.com/DemonZack/simplejrpc-go/net/gsock"
	rpc "github.com/DemonZack/simplejrpc-go"
)

type Handler struct{}

func (h *Handler) Hello(req *gsock.Request) (any, error) {
	return "Hello World", nil
}

func main() {
	srv := rpc.NewDefaultServer(
		gsock.WithJsonRpcSimpleServiceHandler(gsock.NewJsonRpcSimpleServiceHandler()),
	)
	
	srv.RegisterHandle("hello", (&Handler{}).Hello)
	err := srv.StartServer("rpc.sock")
	if err != nil {
		panic(err)
	}
}
```

## Contributing
We welcome contributions! Please follow these steps:
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a pull request

