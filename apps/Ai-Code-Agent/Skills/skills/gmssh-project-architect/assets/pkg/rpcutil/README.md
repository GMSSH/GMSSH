# rpcutil - RPC 请求解析工具

提供 RPC 请求的**参数解析与国际化语言设置**的统一入口。

## API

### `ParseParams(req, args)`

**核心函数** — 解析 RPC 请求参数并自动设置国际化语言。

```go
func ParseParams(req *gsock.Request, args any) error
```

**功能**:
1. 从请求中提取 `lang` 字段并设置当前语言环境
2. 将请求参数反序列化到目标结构体
3. 参数为空时返回国际化错误信息

**使用示例**:

```go
import "template/pkg/rpcutil"

func (s *Server) GetList(req *gsock.Request) (any, error) {
    var args struct {
        Lang string `json:"lang"`
        Page int    `json:"page"`
        Size int    `json:"size"`
    }
    if err := rpcutil.ParseParams(req, &args); err != nil {
        return nil, err
    }
    // args.Page, args.Size 已可用
}
```

### `SetLanguage(req)`

单独设置语言环境（`ParseParams` 已内部调用，通常无需直接使用）。

```go
func SetLanguage(req *gsock.Request)
```

## 设计说明

- 所有 RPC Handler **必须使用** `ParseParams` 解析参数，以确保国际化机制正常工作
- 错误消息 `InvalidParams` 定义在 `i18n/` 目录的语言文件中
