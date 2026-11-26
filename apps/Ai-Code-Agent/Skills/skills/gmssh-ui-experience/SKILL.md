---
name: gmssh-ui-experience
description: 负责 GMSSH 子应用的前端开发与交互设计。专注于构建类似 ChatGPT/Google/Notion 风格的现代科技感美学，支持暗色和亮色模式，强制使用 Vue 3 + Naive UI 组件库，专注于构建具有现代科技美感的界面，并确保所有资源均为本地化依赖。
metadata:
  version: "1.0.0"
  style_reference: "Modern Tech / Minimalist"
  ui_framework: "Vue 3"
  component_library: "Naive UI"
  package_manager: "pnpm"
---

# 前端交互与美学开发指令

## 1. 组件开发准则
- **核心组件库**：强制使用 [Naive UI](https://www.naiveui.com/)。禁止随意混合使用其他 UI 框架（如 Element Plus 或 Ant Design）。
- **组件调用**：优先使用 Naive UI 的按需引入模式，确保包体积精简。
- **反馈拦截**：Naive UI 的 `useMessage` 和 `useDialog` 必须与 `window.$gm` 提供的系统级 API 进行封装联动，确保通知样式全系统统一。

## 2. 视觉美学 (Naive UI 定制)
- **主题驱动**：必须通过 `NConfigProvider` 注入自定义科技感主题。参考 [视觉规范文档](references/design-system.md)。
- **审美标准**：对标 Google/ChatGPT 风格，采用极简主义布局。利用 Naive UI 的 `NSpace`, `NCard`, `NDataTable` 构建高信息密度的专业工具界面。
- **动效控制**：利用 Naive UI 自带的过渡动画，并针对按钮、侧边栏等交互增加微动效，提升“现代感”。

## 3. 工程化与本地化约束
- **包管理**：强制使用 `pnpm`。
- **资源本地化**：所有 Naive UI 及周边依赖（如图标库 `xicons`）必须通过 pnpm 安装。严禁在代码中出现外部 CDN 链接。
- **自适应设计**：利用 Naive UI 的栅格系统 (`NGrid`) 充分考虑全屏模式、小窗口应用以及高分辨率屏幕的显示适配。

## 4. 执行流程 (Workflow)
1. **环境准备**：执行 `pnpm add naive-ui` 及相关依赖。
2. **主题配置**：根据 `design-system.md` 配置全局主题变量。
3. **交互实现**：按照 `sdk-best-practices.md` 桥接 Naive UI 组件与 GMSSH SDK。
4. **自检**：确认在断网环境下（仅使用本地资源）界面渲染完全正常。