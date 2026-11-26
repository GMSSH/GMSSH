# GMSSH 外置应用后端脚手架 (Template)

GMSSH 外置应用的 Go 后端标准工程模板，基于 `simplejrpc-go` 框架，提供开箱即用的 RPC 服务骨架。

## 技术栈

- **语言**: Go 1.24+
- **RPC 框架**: [simplejrpc-go](https://github.com/DemonZack/simplejrpc-go) (Unix Socket + JSON-RPC)
- **日志**: zap (uber-go)
- **国际化**: ini 配置文件

## 目录结构

```
template/
├── main.go              # 程序入口
├── config.json          # 环境配置 (test/prod)
├── go.mod               # Go 模块定义
├── i18n/                # 国际化语言文件
├── internal/            # 核心业务代码 (私有)
│   ├── common/          # 通用常量与工具
│   ├── core/            # 核心适配器与基础设施
│   │   └── adapter/     # 外部服务适配器
│   ├── dto/             # 数据传输对象
│   ├── handler/         # 请求处理层
│   │   └── rpc/         # RPC 路由注册与处理器
│   ├── model/           # 数据模型定义
│   └── service/         # 业务逻辑层
├── logs/                # 运行日志输出
├── pkg/                 # 可复用工具库 (公开)
│   ├── files/           # 文件操作工具集
│   └── rpcutil/         # RPC 请求解析工具
└── test/                # 测试代码
```

## 快速开始

### 1. 克隆并初始化

```bash
# 克隆模板
git clone <template-repo-url> my-app
cd my-app

# 修改模块名 (将 go.mod 中的 "template" 替换为你的应用名)
# 例如: module gmssh-apps/my-app
```

### 2. 注册 RPC 接口

在 `internal/handler/rpc/server.go` 的 `RegisterHandles` 方法中添加路由：

```go
func (s *Server) RegisterHandles(ds interface {
    RegisterHandle(name string, handler func(*gsock.Request) (any, error), middlewares ...gsock.RPCMiddleware)
}) {
    ds.RegisterHandle("ping", s.Ping)
    ds.RegisterHandle("myMethod", s.MyMethod)  // 新增接口
}
```

### 3. 编写业务逻辑

在 `internal/service/` 中实现业务，在 `internal/handler/` 中对接 RPC。

### 4. 编译运行

```bash
go build -o app .
./app
```

## 配置说明

`config.json` 支持 `test` 和 `prod` 两个环境配置，主要包含日志选项：

| 字段 | 说明 | 示例值 |
|------|------|--------|
| `path` | 日志输出目录 | `logs/` |
| `file` | 日志文件名格式 | `{Y-m-d}.log` |
| `level` | 日志级别 | `info` / `error` |
| `stdout` | 是否输出到控制台 | `true` / `false` |
| `rotateBackupLimit` | 日志保留天数 | `7` |
| `rotateExpire` | 日志过期时间 | `1d` |

## 开发规范

- **分层架构**: `handler` → `service` → `model`，严禁跨层调用
- **参数解析**: 统一使用 `pkg/rpcutil.ParseParams()` 解析请求参数
- **国际化**: 错误提示使用 `gi18n.Instance().T("key")`, 语言文件放在 `i18n/` 目录
- **文件操作**: 使用 `pkg/files` 中的工具函数，而非直接调用 `os` 包
