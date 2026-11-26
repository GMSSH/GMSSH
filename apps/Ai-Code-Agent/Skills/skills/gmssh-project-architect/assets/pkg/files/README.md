# files - 文件操作工具集

提供一组**简洁易用的文件系统操作**封装函数，统一处理路径创建、错误处理等边界情况。

## API 概览

### 判断类

| 函数 | 说明 |
|------|------|
| `Exists(path)` | 判断文件或目录是否存在 |
| `IsDir(path)` | 判断路径是否为目录 |
| `IsFile(path)` | 判断路径是否为文件 |

### 读写类

| 函数 | 说明 |
|------|------|
| `ReadFile(path)` | 读取文件内容 (返回 `[]byte`) |
| `WriteFile(path, data, perm)` | 写入文件 (自动创建父目录) |
| `ReadJSON(path, v)` | 读取 JSON 文件并反序列化 |
| `WriteJSON(path, v, indent)` | 序列化为 JSON 并写入文件 |

### 操作类

| 函数 | 说明 |
|------|------|
| `EnsureDir(path)` | 确保目录存在，不存在则创建 |
| `Remove(path)` | 删除文件或目录 |
| `Copy(src, dst)` | 复制文件 |
| `ListDir(path)` | 列出目录内容 |
| `GetFileSize(path)` | 获取文件大小 (字节) |

## 使用示例

```go
import "template/pkg/files"

// 读写 JSON 配置
var config MyConfig
files.ReadJSON("/path/to/config.json", &config)

config.Version = "1.1.0"
files.WriteJSON("/path/to/config.json", config, true)

// 确保目录存在后写入文件
files.EnsureDir("/data/output")
files.WriteFile("/data/output/result.txt", []byte("done"), 0644)
```

## 设计特点

- `WriteFile` 会**自动创建不存在的父目录**，无需手动 `mkdir`
- 所有操作对错误进行了统一处理
- 相比直接使用 `os` 包，减少了重复的样板代码
