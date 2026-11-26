# common - 通用常量与工具

本目录存放**通用的常量、枚举值、错误码和内部工具函数**，可被 `internal` 下的任意层引用。

## 适合放置的内容

- **常量定义**: 业务状态码、错误码、通用配置键
- **枚举类型**: 业务状态枚举、类型枚举
- **工具函数**: 仅限内部使用的小型辅助函数
- **错误定义**: 自定义错误类型和错误消息

## 示例

```go
package common

// 业务状态码
const (
    StatusActive   = 1
    StatusInactive = 0
)

// 错误码
const (
    ErrCodeInvalidParam = 40001
    ErrCodeNotFound     = 40004
    ErrCodeInternal     = 50000
)
```

## 注意事项

- ⚠️ 如果工具函数需要在应用外部复用，应放到 `pkg/` 目录
- 避免在此处放置业务逻辑代码
- 常量命名使用 PascalCase，保持语义清晰
