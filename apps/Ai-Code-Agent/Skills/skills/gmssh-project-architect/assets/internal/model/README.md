# model - 数据模型定义

本目录存放**业务实体和数据模型**的结构体定义。

## 与 DTO 的区别

| 类型 | 位置 | 用途 |
|------|------|------|
| **Model** | `model/` | 内部业务实体，与存储结构对应 |
| **DTO** | `dto/` | 接口传输对象，对外暴露 |

## 示例

```go
package model

import "time"

// User 用户模型
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Status    int       `json:"status"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

// Config 配置模型
type Config struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
```

## 开发规范

- 一个业务模块一个文件，如 `user.go`、`config.go`
- Model 可包含简单的数据转换方法
- 不应依赖 `handler` 或 `service` 层的代码
