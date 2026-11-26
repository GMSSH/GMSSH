# handler - 请求处理层

本目录是**接入层**，负责接收外部请求、解析参数、调用 `service` 层并返回响应。

## 子目录

| 目录 | 说明 |
|------|------|
| `rpc/` | RPC 服务 — JSON-RPC over Unix Socket |

## Handler 的职责

1. **解析请求参数** — 使用 `rpcutil.ParseParams()` 反序列化
2. **参数校验** — 基础的入参校验
3. **调用 Service** — 将业务逻辑委托给 `service` 层
4. **构建响应** — 组装返回结果

## 扩展说明

如果未来需要支持其他协议（如 HTTP、WebSocket），可在 `handler/` 下新建对应子目录：

```
handler/
├── rpc/        # JSON-RPC 处理器 (当前)
├── http/       # HTTP API 处理器 (扩展)
└── ws/         # WebSocket 处理器 (扩展)
```

## 注意事项

- Handler 中**不应包含复杂的业务逻辑**，应委托给 `service`
- 统一使用 `rpcutil.ParseParams()` 解析参数以确保语言设置正确
