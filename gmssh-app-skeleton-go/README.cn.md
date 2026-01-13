
# simplejrpc-go

`simplejrpc-go` 是一个轻量级的 Go 语言 JSON-RPC 框架，旨在帮助开发者构建高性能、可维护的 JSON-RPC 服务，并内置多种实用功能。

## 功能特性
- **配置管理**：从 JSON 文件加载配置，支持使用点号表示法访问嵌套配置
- **结构化日志**：基于 `go.uber.org/zap` 实现，支持日志轮转、标准输出和自定义日志级别
- **数据验证**：全面的结构体字段验证，支持必填、长度、数值范围等规则
- **国际化支持**：国际化资源加载（当前支持 INI 文件格式）
- **网络通信**：JSON-RPC 服务器实现，支持中间件

## 安装
```sh
go get github.com/DemonZack/simplejrpc-go@latest
```

## 配置管理

示例 `config.json`：
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

使用示例：
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

## 日志

日志配置示例：
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
		panic(fmt.Sprintf("初始化日志失败: %v", err))
	}
	defer logger.Sync()

	logger.Info("日志已初始化",
		zap.String("path", config["path"].(string)),
		zap.Int("backupLimit", config["rotateBackupLimit"].(int)),
	)
}
```

## 数据验证

验证示例：
```go
package main

import (
	"fmt"
	"github.com/DemonZack/simplejrpc-go/core"
)

type User struct {
	Username string `validate:"min_length:6#用户名至少需要6个字符"`
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
		fmt.Println("验证错误:", err)
	}
}
```

## 国际化 (i18n)

国际化示例：
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

## JSON-RPC 服务器

服务器实现：
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

## 贡献指南
欢迎贡献代码！请按以下步骤操作：
1. Fork 本仓库
2. 创建您的功能分支
3. 提交您的更改
4. 推送到分支
5. 创建 Pull Request
