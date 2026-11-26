---
name: gmssh-project-architect
description: 负责 GMSSH 外置应用 Go 后端的架构设计与代码生成。基于 simplejrpc-go 脚手架，遵循严格的分层架构（handler → service → model），提供 JSON-RPC over Unix Socket 通信能力。当涉及后端接口开发、业务逻辑编写或数据模型定义时激活。
metadata:
  version: "1.0.0"
  language: "Go 1.24+"
  framework: "simplejrpc-go"
  transport: "Unix Socket + JSON-RPC"
  scaffold_path: "backend/"
---

# Go 后端架构开发指令

## 1. 脚手架结构 (必须遵循)

后端代码 **必须** 按照 `backend/` 目录下的分层架构组织。详见 [后端架构规范](references/backend-architecture.md)。

```
backend/
├── main.go              # 入口：创建 Server、启动 Socket 监听
├── config.json          # 环境配置 (test/prod)
├── go.mod               # 模块定义 (module template)
├── i18n/                # 国际化 INI 文件 (zh-CN.ini / en.ini)
├── internal/            # 核心业务代码 (私有)
│   ├── common/          # 常量、枚举、错误码
│   ├── core/adapter/    # 外部服务适配器
│   ├── dto/             # 请求/响应结构体 (XxxReq / XxxResp)
│   ├── handler/rpc/     # RPC 路由注册与处理器 (server.go)
│   ├── model/           # 业务实体与数据模型
│   └── service/         # 核心业务逻辑
├── pkg/                 # 可复用工具 (公开)
│   ├── files/           # 文件操作封装
│   └── rpcutil/         # RPC 参数解析 + i18n 设置
└── test/                # 测试代码
```

## 2. 分层调用规则 (强制)

```
handler → service → model → core/adapter
```

- **上层调用下层**，禁止反向依赖
- `common/` 和 `dto/` 为公共层，可被任意层引用
- Handler 不含业务逻辑，Service 不处理 RPC 请求

## 3. 新增业务模块流程

每当需要新增一个业务功能时，**必须** 按以下顺序操作：

1. **`dto/`** — 定义 `XxxReq` / `XxxResp` 结构体（JSON tag 必填）
2. **`model/`** — 定义数据模型（如需持久化）
3. **`service/`** — 实现业务逻辑 (`NewXxxService()` 工厂模式)
4. **`handler/rpc/`** — 编写 Handler 函数（签名：`func(req *gsock.Request) (any, error)`）
5. **`server.go`** — 在 `RegisterHandles` 中注册路由（camelCase 方法名）
6. **`i18n/`** — 添加中英文翻译 Key

## 4. 核心约束

- **参数解析**：所有 Handler 必须使用 `rpcutil.ParseParams(req, &args)` 解析参数
- **国际化**：错误消息使用 `gi18n.Instance().T("Key")`，翻译文件在 `i18n/`
- **文件操作**：使用 `pkg/files` 封装，禁止直接调用 `os` 包
- **ping 接口**：禁止删除或修改，平台健康检查依赖此接口
- **跨平台编译**：参考 [跨平台规范](references/cross-platform.md)

## 5. 参考文档

- [后端架构完整规范](references/backend-architecture.md) — 分层详解、代码模板、配置说明
- [JSON-RPC 通信协议](references/jsonrpc-spec.md) — 请求/响应格式、错误码
- [跨平台编译指南](references/cross-platform.md) — CGO、交叉编译