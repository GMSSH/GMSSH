# service - 业务逻辑层

本目录是应用的**核心业务层**，所有业务逻辑的实现都应放在这里。

## 职责

- 实现具体的业务流程和规则
- 编排对 `model` 层和 `core/adapter` 的调用
- 处理业务异常和错误
- 数据格式转换 (Model ↔ DTO)

## 开发规范

### 文件组织

按业务模块划分文件：

```
service/
├── user.go         # 用户相关业务
├── file.go         # 文件管理业务
└── config.go       # 配置管理业务
```

### 代码结构

建议每个模块使用 struct + 方法的方式组织：

```go
package service

import "template/internal/model"

// UserService 用户业务
type UserService struct {
    // 可注入适配器依赖
}

func NewUserService() *UserService {
    return &UserService{}
}

// GetUser 获取用户信息
func (s *UserService) GetUser(id int) (*model.User, error) {
    // 业务逻辑实现
    return nil, nil
}
```

## 注意事项

- **不要直接处理 RPC 请求/响应**，那是 `handler` 的职责
- **可被多个 handler 复用**（如 RPC 和未来的 HTTP 共用同一 service）
- 复杂业务流程建议添加详细注释
