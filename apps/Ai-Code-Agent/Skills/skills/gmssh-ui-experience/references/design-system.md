# GMSSH 视觉设计系统规范

> 本规范定义了 GMSSH 子应用的全局样式体系。所有前端开发 **必须** 基于此设计系统构建，禁止使用硬编码颜色值。

---

## 1. 架构概览

设计系统由两层组成，协同工作：

| 层级 | 文件 | 作用 | 使用方式 |
|------|------|------|----------|
| **CSS 变量层** | [theme.css](theme.css) | 定义全局 Design Tokens（颜色、字号） | 通过 `var(--jm-xxx)` 在 CSS 和 TS 中引用 |
| **Naive UI 主题层** | [dark.ts](dark.ts) | 将 CSS 变量映射到 Naive UI 组件 | 通过 `NConfigProvider :theme-overrides` 注入 |

```
theme.css (Design Tokens)
    ↓ var(--jm-xxx)
dark.ts (Naive UI Overrides)
    ↓ NConfigProvider
Vue 组件 (最终渲染)
```

> [!CAUTION]
> **禁止在组件中硬编码颜色值**（如 `color: #5772FF`）。必须使用 `var(--jm-primary-1)` 或引用 `dark.ts` 中的 token。

---

## 2. CSS 变量层 — Design Tokens

文件：[theme.css](theme.css)，挂载在 `:root[data-theme='dark']` 下。

### 2.1 主色 (Primary)

| 变量 | 值 | 用途 |
|------|----|------|
| `--jm-primary-1` | `#5772FF` | **核心主色**，按钮、链接、高亮 |
| `--jm-primary-2` | `#708BFF` | Hover 态、次级强调 |
| `--jm-primary-3` | `#A8BDFF` | 选中态背景、弱高亮 |
| `--jm-primary-4` | `#000000` | 深色锚点 |

### 2.2 中性色 (Accent / Gray Scale)

从深到浅的灰阶，构成暗色主题的骨架：

| 变量 | 值 | 用途 |
|------|----|------|
| `--jm-accent-1` | `#222222` | 卡片/表格背景 |
| `--jm-accent-2` | `#393939` | 边框、分割线 |
| `--jm-accent-3` | `#555555` | 次级边框、禁用态 |
| `--jm-accent-4` | `#727272` | Placeholder 文字 |
| `--jm-accent-5` | `#afafaf` | 辅助文字 |
| `--jm-accent-6` | `#d2d2d2` | 次要正文 |
| `--jm-accent-7` | `#ffffff` | **主文字色 / 白色** |

> [!TIP]
> 每个颜色都有对应的 RGB 变量（如 `--jm-accent-1-rgb: 34, 34, 34`），便于在需要 `rgba()` 透明度时使用：`rgba(var(--jm-accent-1-rgb), 0.5)`。

### 2.3 状态色 (Semantic)

| 变量 | 值 | 用途 |
|------|----|------|
| `--jm-error-color` | `#f04d3c` | 错误、删除操作 |
| `--jm-success-color` | `#3bad5b` | 成功、确认 |
| `--jm-warning-color` | `#f4bf50` | 警告、注意 |

### 2.4 字号 (Typography)

| 变量 | 值 | 用途 |
|------|----|------|
| `--jm-font-size-12` | `12px` | 标签、表格、辅助文字 |
| `--jm-font-size-14` | `14px` | **正文默认** |
| `--jm-font-size-16` | `16px` | 小标题 |
| `--jm-font-size-18` | `18px` | 页面标题 |

### 2.5 场景色 (Legacy / Specific)

| 变量 | 值 | 用途 |
|------|----|------|
| `--jm-theme` | `rgba(0,0,0,0.3)` | 半透明主题背景（弹窗、输入框） |
| `--jm-bg-color` | `#101010` | 全局页面背景 |
| `--jm-system-hover-color` | `rgba(255,255,255,0.08)` | 系统菜单 Hover 态 |

---

## 3. Naive UI 主题层 — 组件覆盖

文件：[dark.ts](dark.ts)，导出 `theme` 对象，通过 `NConfigProvider` 注入。

### 3.1 使用方式

```vue
<script setup lang="ts">
import { theme } from '@/theme/dark'
</script>

<template>
  <n-config-provider :theme-overrides="theme">
    <router-view />
  </n-config-provider>
</template>
```

### 3.2 Common (全局基础)

```typescript
common: {
  primaryColor: '#5772FF',        // 主色 = --jm-primary-1
  primaryColorHover: '#708BFF',   // Hover = --jm-primary-2
  bodyColor: '#101010',           // 页面背景 = --jm-bg-color
  cardColor: '#101010',           // 卡片背景
  textColor1: 'var(--jm-accent-7)', // 主文字 = 白色
}
```

### 3.3 已覆盖的组件清单

| 组件 | 关键定制 | 设计意图 |
|------|----------|----------|
| **Button** | 全类型 (primary/warning/info/success/error) 统一映射到 `--jm-primary-1`，三档尺寸 (28/34/40px) | 品牌一致性，运维场景紧凑布局 |
| **Input** | 高度 34px，`--jm-theme` 背景，`--jm-accent-4` placeholder | 暗色沉浸式输入体验 |
| **Select** | peers 嵌套定制 `InternalSelection` + `InternalSelectMenu` | 下拉菜单完全融入暗色主题 |
| **DataTable** | `--jm-accent-1` 背景，12px 字体，10px 内边距 | 高信息密度的运维数据展示 |
| **Card** | `--jm-theme` 背景，`--jm-accent-2` 边框 | 半透明科技感卡片 |
| **Dialog** | 状态色图标 (error/warning/success/info) | 与系统反馈色保持一致 |
| **Message** | 8px 圆角，状态色图标 | 轻量通知，风格统一 |
| **Modal / Drawer** | `--jm-theme` 背景 + 阴影 | 沉浸式浮层 |
| **Switch** | `--jm-accent-3` 关闭态，`--jm-primary-1` 开启态 | 直观的开关状态对比 |
| **Tooltip** | 深灰半透明 `rgba(57,57,57,0.95)` | 不遮挡内容，低干扰 |
| **Tag / Skeleton / Dropdown / Popconfirm** | 统一使用 accent 灰阶 | 全局视觉一致性 |

---

## 4. 开发规范

### 4.1 CSS 书写规则

```css
/* ✅ 正确 — 使用 Design Token */
.my-card {
  background: var(--jm-accent-1);
  border: 1px solid var(--jm-accent-2);
  color: var(--jm-accent-7);
  font-size: var(--jm-font-size-14);
}

/* ✅ 正确 — 使用 RGB 变量实现透明度 */
.overlay {
  background: rgba(var(--jm-primary-1-rgb), 0.15);
}

/* ❌ 错误 — 硬编码颜色值 */
.my-card {
  background: #222222;
  color: white;
}
```

### 4.2 组件开发规则

1. **优先使用 Naive UI 组件**，其样式已由 `dark.ts` 统一覆盖
2. **自定义组件** 必须引用 `--jm-*` CSS 变量
3. **新增组件覆盖** 应添加到 `dark.ts` 而非内联 style
4. **尺寸参考**：按钮三档 `28/34/40px`，输入框 `34px`，表格字体 `12px`

### 4.3 颜色速查

```
主色:     --jm-primary-1  (#5772FF)   → 按钮、链接、高亮
悬停:     --jm-primary-2  (#708BFF)   → Hover 态
背景:     --jm-bg-color   (#101010)   → 页面底色
卡片:     --jm-accent-1   (#222222)   → 卡片、表格行
边框:     --jm-accent-2   (#393939)   → 分割线、边框
文字:     --jm-accent-7   (#ffffff)   → 主文字
弱文字:   --jm-accent-4   (#727272)   → Placeholder
错误:     --jm-error-color (#f04d3c)  → 删除、报错
成功:     --jm-success-color (#3bad5b) → 确认、完成
警告:     --jm-warning-color (#f4bf50) → 提醒、注意
```