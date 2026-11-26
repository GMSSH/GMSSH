# dto - 数据传输对象

本目录存放 **DTO (Data Transfer Object)**，即 RPC 接口的**请求参数和响应结果**的结构体定义。

## 命名规范

- 请求结构体: `XxxReq` 或 `XxxRequest`
- 响应结构体: `XxxResp` 或 `XxxResponse`
- 文件按业务模块命名，如 `user.go`、`file.go`

## 示例

```go
package dto

// GetUserReq 获取用户信息请求
type GetUserReq struct {
    Lang   string `json:"lang"`   // 语言标识 (由 rpcutil 自动处理)
    UserID int    `json:"userId"` // 用户ID
}

// GetUserResp 获取用户信息响应
type GetUserResp struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

## 使用方式

在 `handler` 层中通过 `rpcutil.ParseParams` 解析请求参数：

```go
func (s *Server) GetUser(req *gsock.Request) (any, error) {
    var args dto.GetUserReq
    if err := rpcutil.ParseParams(req, &args); err != nil {
        return nil, err
    }
    // 调用 service 层...
}
```

## 注意事项

- DTO 只用于数据传输，**不应包含业务逻辑**
- 所有字段必须添加 `json` tag
- `lang` 字段用于国际化，由框架自动提取，无需业务处理
