# test - 测试代码

本目录存放**集成测试和端到端测试**代码。

## 测试分类

| 类型 | 位置 | 说明 |
|------|------|------|
| **单元测试** | 各 `*_test.go` 文件旁 | 与源码同目录 |
| **集成测试** | `test/` (本目录) | 跨模块/跨层测试 |
| **端到端测试** | `test/` (本目录) | 完整业务流测试 |

## 命名规范

- 测试文件: `xxx_test.go`
- 测试函数: `TestXxx(t *testing.T)`
- 辅助文件: `testdata/` 目录存放测试数据

## 运行测试

```bash
# 运行所有测试
go test ./...

# 运行指定测试
go test ./test/ -run TestXxx -v

# 查看覆盖率
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 编写建议

- 集成测试应模拟完整的 RPC 调用流程
- 使用 `t.Cleanup()` 确保测试资源清理
- 测试数据放在 `testdata/` 子目录，Go 工具链会自动忽略该目录
