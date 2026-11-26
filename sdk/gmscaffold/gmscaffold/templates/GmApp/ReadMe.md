# {{ description }}


# 应用后端目录仓库规范

## 📁 应用后端仓库目录结构

```text
example/
├── app/
│   ├── consts/           # 常量定义（如状态码、配置项等）
│   ├── i18n/             # 国际化模块（默认支持中英文）
│   ├── middlewares/      # 中间件（如鉴权、请求拦截等）
│   ├── schemas/          # 入参校验（Pydantic 或自定义规则）
│   ├── services/         # 核心业务逻辑处理
│   ├── utils/            # 工具模块（如日志、RPC调用等）
│   └── __init__.py
├── config.yaml           # 应用配置文件（日志、端口、元信息等）
├── install.sh            # 安装脚本（如 Redis、依赖库等）
├── main.py               # 应用入口，含服务注册与启动逻辑
├── Makefile              # 编译/打包等命令定义
├── requirements.txt      # Python 依赖列表（可选用 Poetry 管理）
├── .gitignore            # Git 忽略列表
├── README.md             # 项目说明文档
```
## 📁 makefile 构建脚本

```makefile
# 项目构建Makefile

# 定义构建目标目录
DIST_DIR := dist

# 构建后端应用
build:
	@echo "开始构建后端应用..."
	@chmod +x build.sh
	@./build.sh
	@echo "后端应用构建完成!"

# 清理构建结果
clean:
	@echo "清理构建结果..."
	@rm -rf $(DIST_DIR)
	@echo "清理完成!"

# 帮助信息
help:
	@echo "可用命令:"
	@echo "  make build - 构建后端应用"
	@echo "  make clean     - 清理构建结果"
	@echo "  make help      - 显示此帮助信息"

.PHONY: build clean help
```

## 🛠️ 构建指令

```bash
# 构建
make build

# 清理构建目录
make clean
```


### 制品项目结构示例规范

```text
example/
├── tmp/                    # 临时文件目录（必须）
├── cache/                  # 缓存目录（必须）
├── logs/                   # 日志目录（必须）
├── data/                   # 应用数据目录（必须）
│   └── config.yaml         # 应用配置文件（可选）
├── app/                    # 应用主目录（必须）
│   ├── www/                # 前端资源（必须）
│   │   ├── index.html
│   │   ├── static/
│   │   └── icon.png
│   ├── bin/                # 后端程序（必须）
│   │   ├── main
│   │   └── install.sh
│   └── app.json            # 应用元信息文件（必须）
```

## 🧾 app.json 元信息规范

```json
{
  "name": "official/example",
  "title": "示例应用",
  "icon": "./icon.png",
  "version": "1.0.0",
  "sys_req": ["centos", "ubuntu"],
  "app_req": ["nginx", "redis"],
  "conflict": ["old-webhook"],
  "icon_menu": [
    {
      "name": "打开所在文件夹",
      "icon": "./open.png",
      "call_func": "openWithPath"
    }
  ]
}
```

### 字段说明

| 字段名 | 类型 | 必填 | 示例 | 描述 |
|--------|------|------|------|------|
| name | string | ✅ | "xiaoming/webhook" | 应用唯一标识（建议加组织前缀） |
| title | string | ✅ | "Webhook 管理器" | 应用显示名称 |
| icon | string | ✅ | "./icon.png" | 图标路径（建议64x64 PNG/SVG） |
| version | string | ✅ | "1.0.0" | 语义化版本号 |
| sys_req | string[] | 否 | ["centos"] | 系统依赖要求 |
| app_req | string[] | 否 | ["nginx"] | 应用依赖要求 |
| conflict | string[] | 否 | ["old-webhook"] | 冲突应用列表 |
| icon_menu | object[] | 否 | - | 自定义快捷菜单 |

## 🧩 icon_menu 示例

```json
{
  "name": "打开所在文件夹",
  "icon": "./open.png",
  "call_func": "openWithPath"
}
```
