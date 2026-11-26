# 跨平台编译与兼容性规范

> 本规范定义了 GMSSH 后端 (Go) 开发与编译的硬性约束，确保单一二进制文件在主流 Linux 发行版上"即插即用"。

---

## 1. 目标环境矩阵

### 1.1 支持的架构与发行版

| 架构 | 发行版 | 最低版本 | 内核版本 | glibc 版本 |
|------|--------|----------|----------|------------|
| x86_64 (amd64) | Ubuntu | 18.04 LTS | 4.15 | 2.27 |
| x86_64 (amd64) | CentOS | 7.x | 3.10 | **2.17** (最低基线) |
| x86_64 (amd64) | Debian | 9 (Stretch) | 4.9 | 2.24 |
| ARM64 (aarch64) | Ubuntu | 20.04 LTS | 5.4 | 2.31 |
| ARM64 (aarch64) | CentOS | 8.x+ | 4.18 | 2.28 |

> [!IMPORTANT]
> CentOS 7 的 glibc 2.17 是整个矩阵的**最低公分母**，这直接决定了必须采用纯静态编译策略。

### 1.2 Go 版本要求

- **最低版本**: Go 1.24+ (要求 Linux Kernel ≥ 3.2)
- **Toolchain**: 项目已锁定 `go1.24.5`

---

## 2. 编译策略

### 2.1 黄金编译命令

```bash
# ━━━ x86_64 ━━━
GOOS=linux GOARCH=amd64 GOAMD64=v1 CGO_ENABLED=0 \
  go build -trimpath -ldflags="-s -w" -o <app_name>_linux_amd64

# ━━━ ARM64 ━━━
GOOS=linux GOARCH=arm64 GOARM64=v8.0 CGO_ENABLED=0 \
  go build -trimpath -ldflags="-s -w" -o <app_name>_linux_arm64
```

### 2.2 关键参数解释

| 参数 | 作用 | 必要性 |
|------|------|--------|
| `CGO_ENABLED=0` | 禁用 CGO，生成纯静态二进制 | **强制** |
| `GOAMD64=v1` | 兼容 2013 年前 CPU (无 AVX2) | **强制** (amd64) |
| `GOARM64=v8.0` | 兼容所有 ARMv8 芯片 (鲲鹏/飞腾/树莓派) | **强制** (arm64) |
| `-trimpath` | 移除编译路径信息，增强安全性 | 推荐 |
| `-ldflags="-s -w"` | 去除符号表和调试信息，减小体积 | 推荐 |

> [!CAUTION]
> **严禁使用 `-extldflags '-static'`**：在 `CGO_ENABLED=0` 下此参数无意义且可能引发警告。它仅在 CGO 开启时才有效。

---

## 3. 依赖管控

### 3.1 CGO 替代方案速查

当业务需要以下能力时，**禁止引入 CGO 依赖**，必须使用纯 Go 替代：

| 需求 | ❌ 禁用 (需 CGO) | ✅ 替代 (Pure Go) |
|------|------------------|-------------------|
| SQLite | `mattn/go-sqlite3` | `modernc.org/sqlite` |
| 图像处理 | C 绑定的 ImageMagick | `disintegration/imaging` |
| 加密/TLS | OpenSSL 绑定 | Go 标准库 `crypto/*` |
| DNS 解析 | 系统 CGO resolver | `net.DefaultResolver` (纯 Go 模式) |
| 压缩 | C 绑定 zlib | `compress/*` 标准库 |

### 3.2 第三方库准入规则

引入任何新依赖前，必须满足以下**全部条件**：

1. **Pure Go** — 仓库中不含 `.c`、`.h`、`.s` 汇编文件
2. **无 `import "C"` 指令** — `grep -r 'import "C"'` 结果为空
3. **多架构 CI** — 项目 CI 覆盖 amd64 + arm64 两种架构
4. **活跃维护** — 最近 12 个月内有 commit 记录

校验命令：

```bash
# 检查依赖树中是否存在 CGO 引用
go list -deps ./... | xargs -I{} go list -f '{{.Name}} {{.CgoFiles}}' {} 2>/dev/null | grep -v '\[\]'
```

### 3.3 当前项目已验证的安全依赖

| 依赖 | 版本 | 用途 | CGO |
|------|------|------|-----|
| `simplejrpc-go` | v1.0.4 | JSON-RPC 框架 | ❌ |
| `gorilla/websocket` | v1.5.3 | WebSocket 通信 | ❌ |
| `uber-go/zap` | v1.27.0 | 结构化日志 | ❌ |
| `ini.v1` | v1.67.0 | INI 配置文件解析 | ❌ |
| `lumberjack.v2` | v2.2.1 | 日志轮转 | ❌ |

---

## 4. 运行时兼容性约束

### 4.1 文件系统

```go
// ✅ 正确 — 动态路径
home, _ := os.UserHomeDir()
dataDir := filepath.Join(home, ".gmssh", "data")

// ❌ 错误 — 硬编码路径
dataDir := "/root/.gmssh/data"
```

**规则清单**:
- 使用 `path/filepath` 处理所有路径操作（不要使用 `path` 包，它是 URL 路径专用）
- 使用 `os.UserHomeDir()` 替代硬编码 `/root` 或 `/home/xxx`
- 临时文件使用 `os.MkdirTemp()` / `os.CreateTemp()`
- 文件权限始终显式指定（`0644` 普通文件 / `0755` 可执行文件和目录）

### 4.2 进程与信号

```go
// 标准优雅退出模板
ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
defer stop()

// 启动服务...

<-ctx.Done()
// 清理资源: 关闭 Socket、释放数据库句柄、刷写日志缓冲
```

**要求**:
- 必须处理 `SIGTERM` 和 `SIGINT` — GMSSH 宿主通过这些信号管理应用生命周期
- 退出时必须清理 Unix Socket 文件，避免残留导致下次启动失败
- 避免使用 `SIGUSR1/2` 等非标准信号，CentOS 7 的信号处理行为可能不一致

### 4.3 网络与 Socket

- Unix Socket 路径长度限制在 **104 字节** (Linux sunpath 限制为 108)
- Socket 监听前检查并清理残留的 `.sock` 文件
- 连接超时必须设置合理上限，防止在低版本内核上无限阻塞

### 4.4 时间与时区

```go
// ✅ 正确 — 使用 UTC 或显式加载时区
loc, _ := time.LoadLocation("Asia/Shanghai")
now := time.Now().In(loc)

// ❌ 错误 — 依赖系统时区
now := time.Now() // 不同服务器可能是不同时区
```

---

## 5. Makefile 集成模板

建议在项目根目录维护统一的编译入口：

```makefile
APP_NAME := my-gmssh-app
VERSION  := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS  := -s -w -X main.Version=$(VERSION)

.PHONY: build-all build-amd64 build-arm64 check-deps clean

build-all: build-amd64 build-arm64

build-amd64:
	GOOS=linux GOARCH=amd64 GOAMD64=v1 CGO_ENABLED=0 \
	  go build -trimpath -ldflags="$(LDFLAGS)" -o dist/$(APP_NAME)_linux_amd64

build-arm64:
	GOOS=linux GOARCH=arm64 GOARM64=v8.0 CGO_ENABLED=0 \
	  go build -trimpath -ldflags="$(LDFLAGS)" -o dist/$(APP_NAME)_linux_arm64

check-deps:
	@echo "=== 检查 CGO 依赖 ==="
	@go list -deps ./... | xargs -I{} go list -f '{{.Name}} {{.CgoFiles}}' {} 2>/dev/null | grep -v '\[\]' || echo "✅ 无 CGO 依赖"

clean:
	rm -rf dist/
```

---

## 6. 常见问题排查

| 现象 | 原因 | 解决方案 |
|------|------|----------|
| `GLIBC_2.28 not found` | CGO 未禁用，动态链接了高版本 glibc | 确认 `CGO_ENABLED=0` |
| `illegal instruction` | 使用了高阶指令集 (AVX2等) | 确认 `GOAMD64=v1` |
| `bind: address already in use` | 上次退出未清理 Socket 文件 | 启动前 `os.Remove(sockPath)` |
| 日志时间不对 | 服务器时区非预期 | 显式加载 `time.Location` |
| `too many open files` | CentOS 7 默认 ulimit 较低 | 程序内 `syscall.Setrlimit` 或部署调整 |