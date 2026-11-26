本系统文档定义了 GMSSH 前端 SDK (`window.$gm`) 的 UI 交互与窗口管理方法。所有调用必须确保在 `window.$gm` 存在的前提下执行。

## 1. 消息提示 ($gm.message)

用于提供非阻塞式的状态反馈。

### 基础用法

JavaScript

```
$gm.message.success("数据保存成功");
```

### 方法列表

- `success(content, option?)`: 显示成功信息。
- `error(content, option?)`: 显示错误信息。
- `warning(content, option?)`: 显示警告信息。
- `info(content, option?)`: 显示常规信息。
- `loading(content, option?)`: 显示加载中状态（通常在异步请求开始时调用）。
- `destroyAll()`: 销毁当前页面所有弹出的信息。

### MessageOption 属性

| **属性**           | **类型**  | **说明**                 |
| ------------------ | --------- | ------------------------ |
| `duration`         | `number`  | 自动消失时长 (ms)        |
| `closable`         | `boolean` | 是否显示关闭图标         |
| `showIcon`         | `boolean` | 是否展示图标 (默认 true) |
| `keepAliveOnHover` | `boolean` | 鼠标悬停时是否不销毁     |

------

## 2. 对话框 ($gm.dialog)

用于需要用户确认或承载复杂内容的模态窗口。涉及敏感操作（如删除、执行脚本）必须使用对话框确保“人在回路” 。



### 基础用法

JavaScript

```
$gm.dialog.warning({
    title: "确认删除",
    content: "确定要删除此数据库表吗？此操作不可逆。",
    positiveText: "确定删除",
    negativeText: "取消",
    onPositiveClick: () => { /* 执行删除逻辑 */ }
});
```

### 方法列表

- `create(options)`: 创建自定义对话框。
- `info/success/warning/error(options)`: 创建带预设图标的对话框。
- `destroyAll()`: 关闭所有打开的对话框。

### DialogOptions 关键属性

- `title`: 标题（支持字符串或渲染函数）。
- `content`: 对话框内容主体。
- `positiveText`: 确认按钮文字（不填则不显示）。
- `negativeText`: 取消按钮文字。
- `loading`: 设为 `true` 时，确认按钮显示加载状态。
- `maskClosable`: 点击遮罩层是否允许关闭（默认 true）。

------

## 3. 应用与窗口管理

### 应用退出

- `$gm.closeApp()`: 关闭当前运行的子应用，并触发销毁监听。
- `$gm.closeOtherApp(name)`: 关闭指定的第三方应用（格式：`组织名/应用名`）。

### 生命周期监听

- `$gm.childDestroyedListener(callback)`: 监听当前应用被关闭的事件。最佳实践：在此回调中清理 WebSocket 链接或定时器 。

  

  

- `$gm.mainNotificationListener(callback)`: 监听来自主应用的通知消息。

### 环境交互

- `$gm.updateDesktop()`: 立即刷新桌面图标和布局。在安装、卸载应用或修改桌面配置后调用。

------

## 4. 状态码与异常处理

所有通过 SDK 发起的请求响应必须校验 `code` 字段。

- **200000**: 操作成功。
- **非 200000**: 操作失败，应将 `msg` 字段的内容通过 `$gm.message.error` 展示。