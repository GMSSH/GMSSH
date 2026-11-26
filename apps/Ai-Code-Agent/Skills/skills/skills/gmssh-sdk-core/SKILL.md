---
name: gmssh-sdk-core
description: 负责与 GMSSH 系统底座进行深度集成。封装了基于 gmAppSdk.js 的 UI 反馈、系统控制及 ga_main 代理接口请求逻辑。当涉及调用终端、文件操作或系统对话框时激活。
metadata:
  version: "1.0.0"
  global_variable: "window.$gm"
  api_prefix: "/api/call/"
---

# GMSSH SDK 核心集成指令

## 1. 基础调用规范
- **单例模式**：所有操作必须通过全局变量 `window.$gm` 进行。
- **环境隔离**：严禁引入外部 CDN 资源，必须确保 `gmAppSdk.js` 在 HTML 中以 `<script src="gmAppSdk.js"></script>` 方式本地引入。
- **异步处理**：对于 `$gm.request`、`$gm.chooseFile` 等具有回调或返回值的操作，必须使用 `Promise` 或 `async/await` 进行封装。

## 2. UI 与反馈准则
- **轻量反馈**：普通通知使用 `$gm.message`。
- **重型交互**：涉及删除、更新、执行脚本等操作，必须调用 `$gm.dialog` 并配置 `positiveText` 和 `negativeText` 确保人类在回路（Human-in-the-loop） 。
- **状态管理**：应用启动时需注册 `$gm.childDestroyedListener` 以确保应用关闭时及时释放后端占用的句柄或资源 [1, 4]。

## 3. Agent 服务通信 (ga_main)
- **请求封装**：通过 `$gm.request()` 发送请求。
- **路由规则**：路径遵循 `/api/call/{group}/{app}/{path}` 格式。
- **响应处理**：强制校验 `code === 200000`。若失败，需提取 `msg` 并通过系统 Message 告知用户。

## 4. 终端与文件集成
- **终端交互**：使用 `openShell` 时，必须按照 `_GMC` 规范编写脚本，并在脚本结束时发送 `EndStak` 指令以自动关闭终端 [5, 4]。
- **文件预览**：对于文本类文件，优先调用 `$gm.openCodeEditor` 而非自建编辑器。

## 5. 参考指引
- 详细 UI 方法见(references/sdk-ui-methods.md)。
- Agent 接口定义见 [Agent API 说明](references/agent-api-spec.md)。