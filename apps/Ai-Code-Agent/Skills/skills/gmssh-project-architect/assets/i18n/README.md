# i18n - 国际化语言文件

本目录存放应用的国际化 (i18n) 翻译资源文件，使用 **INI 格式**。

## 文件清单

| 文件 | 语言 | 说明 |
|------|------|------|
| `zh-CN.ini` | 简体中文 | 中文翻译 |
| `en.ini` | English | 英文翻译 |

## 文件格式

采用 `Key="Value"` 格式，每行一个翻译条目：

```ini
InvalidParams="参数无效"
UserNotFound="用户不存在"
OperationSuccess="操作成功"
```

## 使用方式

在 Go 代码中通过 `gi18n` 包调用：

```go
import "github.com/DemonZack/simplejrpc-go/core/gi18n"

// 获取当前语言的翻译文本
msg := gi18n.Instance().T("InvalidParams")
```

语言切换由 `pkg/rpcutil.SetLanguage()` 根据前端请求参数自动完成，开发者通常无需手动设置。

## 添加新语言

1. 创建对应的 `.ini` 文件（如 `ja.ini`）
2. 将所有已有的 Key 翻译为对应语言
3. 确保 Key 名称与其他语言文件完全一致

## 命名规范

- **Key 命名**: 使用 PascalCase，如 `InvalidParams`、`UserNotFound`
- **语义化**: Key 应描述消息用途，而非具体文案
- **完整性**: 所有语言文件必须包含相同的 Key 集合
