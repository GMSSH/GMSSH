# Go 后端脚手架架构规范

> 本文档是 GMSSH 外置应用 Go 后端脚手架的完整架构参考。AI 在生成后端代码时 **必须** 严格遵循此规范。

---

## 1. 技术栈

| 项目 | 技术 | 版本 |
|------|------|------|
| 语言 | Go | 1.24+ |
| RPC 框架 | `simplejrpc-go` | v1.0.4 |
| 传输层 | Unix Socket + JSON-RPC | — |
| 日志 | zap (uber-go) | — |
| 国际化 | INI 配置 + gi18n | — |
| 配置 | JSON (test/prod 双环境) | — |

---

## 2. 入口文件 — `main.go`

```go
package main

import (
    "fmt"
    "github.com/DemonZack/simplejrpc-go/core"
    "template/internal/handler/rpc"
)

func main() {
    sock := "../../tmp/app.sock"

    server := rpc.NewServer()

    if err := server.Start(sock); err != nil {
        core.Container.Log().Error(fmt.Sprintf("启动RPC服务器失败: %v", err))
        return
    }

    core.Container.Log().Info(fmt.Sprintf("服务启动成功, socket: %s", sock))
}
```

> [!IMPORTANT]
> `sock` 路径由 GMSSH 平台动态分配，开发时使用 `../../tmp/app.sock` 进行本地调试。

---

## 3. 分层架构详解

```
handler (接入层)        ← 接收 RPC 请求，解析参数，返回响应
   ↓ 调用
service (业务逻辑层)    ← 核心业务实现，编排 model 和 adapter
   ↓ 调用
model (数据模型层)      ← 业务实体定义
   ↓ 使用
core/adapter (基础设施) ← 外部服务封装（DB、缓存、三方 API）
```

**公共层** (可被任意层引用)：
- `common/` — 常量、枚举、错误码
- `dto/` — 请求/响应传输对象

### 3.1 handler/rpc/ — 接入层

**职责**：接收请求 → 解析参数 → 调用 Service → 返回结果

**`server.go` 核心结构**：

```go
// Server 结构体承载所有 RPC Handler
type Server struct{}

func NewServer() *Server {
    return &Server{}
}

// RegisterHandles 集中注册路由 (camelCase 方法名)
func (s *Server) RegisterHandles(ds interface {
    RegisterHandle(name string, handler func(*gsock.Request) (any, error), middlewares ...gsock.RPCMiddleware)
}) {
    ds.RegisterHandle("ping", s.Ping)           // ⚠️ 必须保留
    ds.RegisterHandle("getItemList", s.GetItemList) // 业务接口
    ds.RegisterHandle("createItem", s.CreateItem)
}

// Start 启动服务器
func (s *Server) Start(sockPath string) error {
    _ = os.MkdirAll(sockPath, 0o755)
    ds := simplejrpc.NewDefaultServer(
        gsock.WithJsonRpcSimpleServiceHandler(gsock.NewJsonRpcSimpleServiceHandler()),
    )
    gpath.GmCfgPath = "./"
    gi18n.Instance().SetPath("./i18n")
    core.InitContainer(config.WithConfigEnvFormatterOptionFunc("test"))
    s.RegisterHandles(ds)
    return ds.StartServer(sockPath)
}
```

**Handler 标准模板**：

```go
func (s *Server) GetItemList(req *gsock.Request) (any, error) {
    // 1. 解析参数 (自动处理 i18n)
    var args dto.GetItemListReq
    if err := rpcutil.ParseParams(req, &args); err != nil {
        return nil, err
    }

    // 2. 参数校验
    if args.Page < 1 {
        return nil, fmt.Errorf(gi18n.Instance().T("InvalidParams"))
    }

    // 3. 调用 Service
    result, err := service.NewItemService().GetList(args)
    if err != nil {
        return nil, err
    }

    // 4. 返回结果 (脚手架自动包装为 JSON-RPC 响应)
    return result, nil
}
```

### 3.2 dto/ — 数据传输对象

**命名规范**：`XxxReq` / `XxxResp`，文件按模块命名

```go
package dto

// GetItemListReq 获取列表请求
type GetItemListReq struct {
    Lang     string `json:"lang"`      // 语言标识 (rpcutil 自动处理)
    Page     int    `json:"page"`      // 页码
    PageSize int    `json:"pageSize"`  // 每页数量
    Search   string `json:"search"`    // 搜索关键词
}

// GetItemListResp 获取列表响应
type GetItemListResp struct {
    Total int         `json:"total"`
    List  []ItemBrief `json:"list"`
}
```

> [!WARNING]
> DTO 只做数据传输，**不包含业务逻辑**。所有字段必须有 `json` tag。

### 3.3 service/ — 业务逻辑层

**工厂模式 + struct 方法**：

```go
package service

type ItemService struct{}

func NewItemService() *ItemService {
    return &ItemService{}
}

func (s *ItemService) GetList(args dto.GetItemListReq) (*dto.GetItemListResp, error) {
    // 业务逻辑
    return &dto.GetItemListResp{Total: 0, List: nil}, nil
}
```

**规则**：
- 不直接处理 RPC 请求/响应
- 可被多个 Handler 复用
- 依赖注入 adapter（如数据库、缓存）

### 3.4 model/ — 数据模型

```go
package model

type Item struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Status    int       `json:"status"`
    CreatedAt time.Time `json:"createdAt"`
}
```

**与 DTO 的区别**：

| | Model | DTO |
|--|-------|-----|
| 位置 | `model/` | `dto/` |
| 用途 | 内部实体，对应存储结构 | 接口传输，对外暴露 |

### 3.5 common/ — 通用定义

```go
package common

const (
    StatusActive   = 1
    StatusInactive = 0
)

const (
    ErrCodeInvalidParam = 40001
    ErrCodeNotFound     = 40004
)
```

### 3.6 core/adapter/ — 基础设施

封装外部服务调用，定义 interface 便于 mock 测试：

```go
type DatabaseAdapter interface {
    Query(sql string, args ...any) ([]map[string]any, error)
}
```

---

## 4. 公共工具包 — `pkg/`

### 4.1 rpcutil — RPC 参数解析

**所有 Handler 必须使用**：

```go
import "template/pkg/rpcutil"

// 解析参数 + 自动设置 i18n 语言
if err := rpcutil.ParseParams(req, &args); err != nil {
    return nil, err
}
```

| 函数 | 说明 |
|------|------|
| `ParseParams(req, &args)` | 解析参数 + 设置语言 (核心) |
| `SetLanguage(req)` | 仅设置语言 (ParseParams 已内部调用) |

### 4.2 files — 文件操作封装

**优先使用此包，禁止直接调用 `os` 包**：

| 函数 | 说明 |
|------|------|
| `Exists(path)` | 判断存在 |
| `IsDir(path)` / `IsFile(path)` | 类型判断 |
| `ReadFile(path)` | 读取文件 → `[]byte` |
| `WriteFile(path, data, perm)` | 写入 (自动创建父目录) |
| `ReadJSON(path, &v)` | 读取 JSON 并反序列化 |
| `WriteJSON(path, v, indent)` | 序列化并写入 JSON |
| `EnsureDir(path)` | 确保目录存在 |
| `Remove(path)` | 删除文件/目录 |
| `Copy(src, dst)` | 复制文件 |
| `ListDir(path)` | 列出目录 |
| `GetFileSize(path)` | 获取大小 |

---

## 5. 国际化 — `i18n/`

**INI 格式**，按语言分文件：

```ini
# zh-CN.ini
InvalidParams="参数无效"
OperationSuccess="操作成功"
ItemNotFound="条目不存在"

# en.ini
InvalidParams="Invalid parameters"
OperationSuccess="Operation successful"
ItemNotFound="Item not found"
```

**使用方式**：

```go
import "github.com/DemonZack/simplejrpc-go/core/gi18n"

msg := gi18n.Instance().T("ItemNotFound")
return nil, fmt.Errorf(msg)
```

> [!IMPORTANT]
> 新增翻译 Key 时，`zh-CN.ini` 和 `en.ini` 必须同步更新。Key 使用 PascalCase。

---

## 6. 环境配置 — `config.json`

```json
{
  "test": {
    "logger": {
      "path": "logs/",
      "file": "{Y-m-d}.log",
      "level": "info",
      "stdout": true,
      "rotateBackupLimit": 7,
      "rotateExpire": "1d"
    }
  },
  "prod": {
    "logger": {
      "level": "error",
      "stdout": false
    }
  }
}
```

通过 `core.InitContainer(config.WithConfigEnvFormatterOptionFunc("test"))` 加载对应环境配置。

---

## 7. 新增模块速查

```
1. dto/item.go        → 定义 GetItemListReq / GetItemListResp
2. model/item.go      → 定义 Item 实体 (如需)
3. service/item.go    → 实现 NewItemService().GetList()
4. handler/rpc/       → 编写 Server.GetItemList() Handler
5. server.go          → ds.RegisterHandle("getItemList", s.GetItemList)
6. i18n/zh-CN.ini     → 添加翻译 Key
6. i18n/en.ini        → 添加翻译 Key
```
