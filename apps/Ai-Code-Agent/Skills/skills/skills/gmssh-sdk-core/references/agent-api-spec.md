# GMSSH SDK 完整 API 参考

> 本文档定义了 `window.$gm` 提供的全部前端 API，涵盖 UI 反馈、应用生命周期、系统交互、终端控制及文件系统操作。
> 所有调用均通过 `gmAppSdk.js` 注入的全局对象 `window.$gm` 发起。

---

## 1. 消息提示 (`$gm.message`)

非阻塞式状态反馈，统一调用宿主的 Message 渲染。

```javascript
$gm.message.success("操作成功");
$gm.message.error("发生错误");
$gm.message.loading("加载中...");
$gm.message.destroyAll(); // 销毁所有消息
```

### 方法列表

| 方法 | 签名 | 说明 |
|------|------|------|
| `info` | `(content: string \| (() => VNodeChild), option?: MessageOption) => MessageReactive` | 常规信息 |
| `success` | 同上 | 成功信息 |
| `warning` | 同上 | 警告信息 |
| `error` | 同上 | 错误信息 |
| `loading` | 同上 | 加载状态 |
| `create` | 同上 | 自定义类型 |
| `destroyAll` | `() => void` | 销毁所有弹出消息 |

### MessageOption

| 属性 | 类型 | 说明 |
|------|------|------|
| `duration` | `number` | 自动消失时长 (ms) |
| `closable` | `boolean` | 是否显示关闭图标 |
| `showIcon` | `boolean` | 是否展示图标 (默认 `true`) |
| `keepAliveOnHover` | `boolean` | 鼠标悬停时不销毁 |
| `icon` | `() => VNodeChild` | 自定义图标渲染函数 |
| `render` | `MessageRenderMessage` | 消息渲染函数 |
| `type` | `'info' \| 'success' \| 'warning' \| 'error' \| 'loading' \| 'default'` | 消息类型 |
| `onAfterLeave` | `() => void` | 消失动画结束回调 |
| `onClose` | `() => void` | 点击关闭图标回调 |
| `onLeave` | `() => void` | 开始消失回调 |

---

## 2. 对话框 (`$gm.dialog`)

模态确认框，**涉及删除、执行脚本等敏感操作必须使用**，确保"人在回路"。

```javascript
$gm.dialog.warning({
  title: "确认删除",
  content: "确定要删除此文件吗？此操作不可逆。",
  positiveText: "确定删除",
  negativeText: "取消",
  maskClosable: false,
  onPositiveClick: () => { /* 执行删除 */ },
});
```

### 方法列表

| 方法 | 说明 |
|------|------|
| `create(options)` | 自定义对话框 |
| `info(options)` | 信息对话框 |
| `success(options)` | 成功对话框 |
| `warning(options)` | 警告对话框 |
| `error(options)` | 错误对话框 |
| `destroyAll()` | 关闭所有对话框 |

### DialogOptions

| 属性 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `title` | `string \| (() => VNodeChild)` | — | 标题 |
| `content` | `string \| (() => VNodeChild)` | — | 内容主体 |
| `positiveText` | `string` | — | 确认按钮文字（不填则不显示） |
| `negativeText` | `string` | — | 取消按钮文字（不填则不显示） |
| `loading` | `boolean` | `false` | 确认按钮是否显示加载状态 |
| `maskClosable` | `boolean` | `true` | 点击遮罩是否可关闭 |
| `closable` | `boolean` | `true` | 是否显示关闭图标 |
| `closeOnEsc` | `boolean` | `true` | 按 Esc 是否关闭 |
| `draggable` | `boolean \| { bounds?: 'none' }` | `false` | 是否可拖拽 |
| `type` | `'error' \| 'success' \| 'warning'` | `'warning'` | 对话框类型 |
| `showIcon` | `boolean` | `true` | 是否显示图标 |
| `iconPlacement` | `'left' \| 'top'` | `'left'` | 图标位置 |
| `icon` | `() => VNodeChild` | — | 自定义图标渲染函数 |
| `action` | `() => VNodeChild` | — | 操作区域渲染函数 |
| `blockScroll` | `boolean` | `true` | 打开时禁用 body 滚动 |
| `autoFocus` | `boolean` | `true` | 自动聚焦第一个可聚焦元素 |
| `bordered` | `boolean` | `false` | 是否显示边框 |
| `transformOrigin` | `'mouse' \| 'center'` | `'mouse'` | 动画出现位置 |
| `positiveButtonProps` | `ButtonProps` | — | 确认按钮属性 |
| `negativeButtonProps` | `ButtonProps` | — | 取消按钮属性 |
| `style` | `string \| Object` | — | 对话框样式 |
| `class` | `any` | — | 对话框类名 |
| `contentClass` | `string` | — | 内容类名 |
| `contentStyle` | `Object \| string` | — | 内容样式 |
| `titleClass` | `string` | — | 标题类名 |
| `titleStyle` | `Object \| string` | — | 标题样式 |
| `actionClass` | `string` | — | 操作区域类名 |
| `actionStyle` | `Object \| string` | — | 操作区域样式 |
| `onPositiveClick` | `(e: MouseEvent) => boolean \| Promise<boolean> \| any` | — | 确认回调，返回 `false` 阻止关闭 |
| `onNegativeClick` | `(e: MouseEvent) => boolean \| Promise<boolean> \| any` | — | 取消回调，返回 `false` 阻止关闭 |
| `onClose` | `() => boolean \| Promise<boolean> \| any` | — | 关闭回调，返回 `false` 阻止关闭 |
| `onMaskClick` | `() => void` | — | 点击蒙层回调 |
| `onAfterEnter` | `() => void` | — | 出现动画结束回调 |
| `onAfterLeave` | `() => void` | — | 关闭动画结束回调 |

---

## 3. 应用生命周期

### 3.1 关闭当前应用

```javascript
$gm.closeApp();
```

> [!WARNING]
> 调用后应用**立即关闭**，请确保已保存重要数据。会触发 `childDestroyedListener` 回调。

### 3.2 关闭第三方应用

```javascript
$gm.closeOtherApp("组织名/应用名");
// 示例：$gm.closeOtherApp("xz/svn");
```

### 3.3 应用销毁监听

```javascript
$gm.childDestroyedListener(() => {
  // 清理 WebSocket、定时器等资源
  console.log("应用已销毁");
});
```

**触发时机**：`closeApp()` 调用 / 用户手动关闭 / 系统强制关闭。

### 3.4 系统通知监听

```javascript
$gm.mainNotificationListener((data) => {
  console.log("收到系统通知:", data);
});
```

### 3.5 刷新桌面

```javascript
$gm.updateDesktop(); // 安装/卸载应用后调用
```

---

## 4. 文件与目录选择

### 4.1 选择文件夹

```javascript
$gm.chooseFolder((path) => {
  console.log("选中目录:", path);
}, "/www"); // 第二个参数：默认打开路径
```

### 4.2 选择文件

```javascript
$gm.chooseFile((file) => {
  console.log("选中文件:", file);
}, "/www");
```

### 4.3 打开文件管理器

```javascript
$gm.openFolder("/www"); // 在系统文件管理器中打开指定目录
```

---

## 5. 命令执行

### 5.1 静默执行 (`execShell`)

后台执行命令，无 UI，返回 Promise。

```javascript
const result = await $gm.execShell({
  cmd: 'ls -la',
  cwd: '/www',       // 可选：工作目录
  timeout: 60000,    // 可选：超时 (ms)
});
console.log(result.stdout, result.stderr, result.exitCode);
```

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cmd` | `string` | ✅ | Shell 命令 |
| `cwd` | `string` | ❌ | 工作目录 |
| `timeout` | `number` | ❌ | 超时时间 (ms) |
| `name` | `string` | ❌ | 执行用户名 (特定环境) |

**返回值**：

| 字段 | 类型 | 说明 |
|------|------|------|
| `stdout` | `string` | 标准输出 |
| `stderr` | `string` | 错误输出 |
| `exitCode` | `number` | 退出码 (0 = 成功) |

> [!CAUTION]
> 注意命令注入风险，尽可能使用绝对路径，避免执行用户输入的未验证命令。

### 5.2 交互式终端 (`openShell`)

打开可见终端窗口，支持传入命令数组。

```javascript
$gm.openShell({ arr: ["cd /www\n", "ls\n"] });
```

### 5.3 脚本执行 (`exec_script`)

通过 REST API 提交 Base64 编码的脚本：

```javascript
const script = `#!/bin/bash\necho "Hello World"`;
const base64 = btoa(unescape(encodeURIComponent(script)));

const res = await $gm.request('/api/command/exec_script', {
  method: 'post',
  data: { script: base64 },
});
```

**错误处理示例**：

```javascript
if (res?.data.toString().match("BashError") !== null) {
  $gm.message.error(res.data.toString().replace('BashError ', ''));
}
```

---

## 6. GMC 终端交互协议

通过 `openShell` 打开终端后，可使用 `_GMC` 指令与宿主通信。

### 6.1 指令格式

```bash
echo _GMC:'{"type": "<指令类型>", "data": <数据>}'
```

### 6.2 指令列表

| 指令 type | 数据 | 说明 |
|-----------|------|------|
| `updateTitle` | `string` | 更新终端窗口标题 |
| `StartStak` | — | 标记脚本正式开始 |
| `EndStak` | — | 结束脚本并自动关闭终端 |
| `emitApp` | `{ app_name, type }` | 向外置应用发送事件通知 |

### 6.3 标准脚本模板

```bash
#!/bin/bash

# 1. 设置终端标题
echo _GMC:'{"type": "updateTitle", "data": "任务执行中..."}'

# 2. 定义清理钩子
cleanup() {
    echo _GMC:'{"type": "emitApp", "data": {"app_name": "org/app", "type": "flushed"}}'
    echo _GMC:'{"type": "EndStak"}'
}
trap cleanup EXIT

# 3. 标记开始
echo _GMC:'{"type": "StartStak"}'

# --- 业务逻辑 ---
echo "正在执行任务..."
# ...
# --- 业务逻辑结束 ---

# 脚本退出自动触发 cleanup
```

### 6.4 前端监听回调

```javascript
$gm.mainGMCListener((key) => {
  if (key === "flushed") {
    // 终端脚本执行完毕，刷新 UI
    console.log("收到终端完成通知");
  }
});
```

### 6.5 完整执行流程

```
1. 将脚本内容写入服务器 .sh 文件
2. 通过 openShell 执行: bash /path/to/script.sh && exit
3. 脚本通过 _GMC 指令与前端通信
4. 脚本结束 → trap cleanup → EndStak 自动关闭终端
5. 前端 mainGMCListener 收到回调 → 更新 UI
```

---

## 7. 文件系统 REST API

通过 `$gm.request()` 调用，所有接口使用 POST 方法。

### 7.1 获取目录列表

**接口**: `/api/files/get_dir`

```javascript
const res = await $gm.request('/api/files/get_dir', {
  method: 'post',
  data: {
    path: "/www",
    search: "",        // 可选：搜索关键词
    sort: "name",      // 可选：name | size | time
    reverse: "False",  // 可选：是否倒序
    showRow: 9999,     // 可选：显示行数（桌面生效）
  },
});
```

**响应结构**：

```json
{
  "code": 200000,
  "data": {
    "LIST": [
      {
        "name": "/www/index.html",
        "filename": "index.html",
        "cwd": "/www",
        "fullname": "/www/index.html",
        "filesize": 1024,
        "size": 1024,
        "mtime": 1751468646,
        "accept": "755",
        "user": "root",
        "type": 0
      }
    ],
    "PATH": "/www"
  }
}
```

| `type` 值 | 文件类型 |
|-----------|----------|
| `0` | 文件 |
| `1` | 目录 |

### 7.2 创建文件

**接口**: `/api/files/create_file`

```javascript
await $gm.request('/api/files/create_file', {
  method: 'post',
  data: { path: "/www/demo.txt" },
});
```

### 7.3 删除文件

**接口**: `/api/files/delete_file`

```javascript
await $gm.request('/api/files/delete_file', {
  method: 'post',
  data: { path: "/www/demo.txt" },
});
```

### 7.4 创建目录

**接口**: `/api/files/create_dir`

```javascript
await $gm.request('/api/files/create_dir', {
  method: 'post',
  data: { path: "/www/new_folder" },
});
```

### 7.5 删除目录

**接口**: `/api/files/delete_dir`

```javascript
await $gm.request('/api/files/delete_dir', {
  method: 'post',
  data: { path: "/www/new_folder" },
});
```

### 7.6 读取文件内容

**接口**: `/api/files/read_file_body`

```javascript
const res = await $gm.request('/api/files/read_file_body', {
  method: 'post',
  data: {
    path: "/www/config.json",
    auto_create: "0",  // "1" = 文件不存在时自动创建
  },
});
console.log(res.data.data); // 文件内容字符串
```

**响应结构**：

```json
{
  "code": 200000,
  "data": {
    "path": "",
    "encoding": "utf-8",
    "data": "文件内容...",
    "only_read": false,
    "size": 1024,
    "st_mtime": "1753061659"
  }
}
```

### 7.7 写入文件内容

**接口**: `/api/files/write_file_body`

```javascript
await $gm.request('/api/files/write_file_body', {
  method: 'post',
  data: {
    path: "/www/config.json",
    data: '{"key": "value"}',
    encoding: "utf-8",  // 可选，默认 utf-8
  },
});
```

---

## 8. API 速查表

| 类别 | API | 说明 |
|------|-----|------|
| **消息** | `$gm.message.success/error/warning/info/loading(msg)` | 非阻塞提示 |
| **对话框** | `$gm.dialog.warning/error/info/success(options)` | 模态确认 |
| **生命周期** | `$gm.closeApp()` | 关闭当前应用 |
| | `$gm.closeOtherApp(name)` | 关闭第三方应用 |
| | `$gm.childDestroyedListener(cb)` | 销毁回调 |
| | `$gm.updateDesktop()` | 刷新桌面 |
| **文件选择** | `$gm.chooseFolder(cb, path)` | 选择文件夹 |
| | `$gm.chooseFile(cb, path)` | 选择文件 |
| | `$gm.openFolder(path)` | 打开文件管理器 |
| **命令** | `$gm.execShell({ cmd, cwd, timeout })` | 静默执行 |
| | `$gm.openShell({ arr })` | 打开终端 |
| **脚本** | `$gm.request('/api/command/exec_script', ...)` | 执行 Base64 脚本 |
| **文件 API** | `$gm.request('/api/files/get_dir', ...)` | 获取目录列表 |
| | `$gm.request('/api/files/create_file', ...)` | 创建文件 |
| | `$gm.request('/api/files/delete_file', ...)` | 删除文件 |
| | `$gm.request('/api/files/create_dir', ...)` | 创建目录 |
| | `$gm.request('/api/files/delete_dir', ...)` | 删除目录 |
| | `$gm.request('/api/files/read_file_body', ...)` | 读取文件 |
| | `$gm.request('/api/files/write_file_body', ...)` | 写入文件 |

> [!IMPORTANT]
> 所有 `$gm.request()` 响应必须校验 `code === 200000`。失败时提取 `msg` 通过 `$gm.message.error()` 告知用户。
