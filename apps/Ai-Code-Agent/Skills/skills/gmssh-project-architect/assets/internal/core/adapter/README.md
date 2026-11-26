# adapter - 外部服务适配器

本目录用于封装**对外部系统/第三方服务的调用**，隔离外部依赖变化对业务层的影响。

## 适用场景

- 调用外部 HTTP API
- 封装数据库操作
- 对接消息队列、缓存等中间件
- 与 GMSSH 宿主系统的交互

## 开发规范

1. **定义接口**: 每个适配器先定义 interface
2. **实现结构体**: 编写具体实现
3. **命名约定**: 文件名与适配的外部服务对应，如 `database.go`、`cache.go`

## 示例

```go
package adapter

// DatabaseAdapter 数据库适配器接口
type DatabaseAdapter interface {
    Query(sql string, args ...any) ([]map[string]any, error)
    Execute(sql string, args ...any) (int64, error)
}

// mysqlAdapter MySQL 适配器实现
type mysqlAdapter struct {
    dsn string
}

func NewMySQLAdapter(dsn string) DatabaseAdapter {
    return &mysqlAdapter{dsn: dsn}
}
```
