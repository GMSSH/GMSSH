# 项目介绍

```
example/
├── app/
│   ├── consts/           # 常量定义模块，例如状态码、配置项等
│   ├── i18n/             # 国际化模块，支持多语言
│   ├── middlewares/      # 中间件逻辑，如鉴权、请求拦截等
│   ├── schemas/          # 入参校验模块
│   ├── services/         # 核心业务逻辑处理模块
│   ├── utils/            # 工具函数模块，如日志、服务调用等
│   └── __init__.py
├── config.yaml           # 应用配置文件，包含日志、端口、元信息等
├── install.sh            # 安装脚本，如安装 Redis/依赖库等
├── main.py               # 应用主入口，包含服务注册与启动逻辑
├── makefile              # 编译、打包等命令定义
├── requirements.txt      # Python依赖清单（可替代为 Poetry 管理）
├── .gitignore            # Git 忽略文件列表
├── README.md             # 项目说明文档
```
---

# 📁 项目结构说明文档

该项目为 GMSSH 平台外置应用的标准工程结构，推荐结合脚手架。

---

## 📂 根目录结构

| 目录/文件              | 说明                        |
| ------------------ |---------------------------|
| `app/`             | 主应用逻辑模块目录                 |
| `config.yaml`      | 应用运行配置文件（端口、日志、服务元信息等）    |
| `install.sh`       | 服务安装脚本（例如 Redis，Mysql 包等） |
| `main.py`          | 应用启动入口，包含注册逻辑与服务启动逻辑      |
| `makefile`         | 执行make build 打包开发的外置应用    |
| `requirements.txt` | Python 依赖列表，可使用 pip 安装    |
| `.gitignore`       | Git 提交忽略配置                |
| `README.md`        | 项目简介、启动说明、开发指导等文档         |

---

## 📁 app 子目录说明

| 子目录            | 功能说明                              |
| -------------- | --------------------------------- |
| `consts/`      | 存放常量定义，如状态码、配置项（如 `SUCCESS = 200`） |
| `i18n/`        | 国际化支持     |
| `middlewares/` | 中间件逻辑，如日志拦截、请求校验、权限处理等            |
| `schemas/`     | 请求参数校验模型      |
| `services/`    | 业务处理逻辑模块，例如 Redis 启停管理等           |
| `utils/`       | 通用工具方法，如日志封装、JsonRPC 调用封装等        |

---

## 📄 config.yaml 配置样例

```yaml
app:
  name: redis-manager
  port: 8080
  log_level: INFO
  version: 1.0.0
meta:
  org: example-org
  author: your-name
```

---

## 📄 install.sh 示例

```bash
#!/bin/bash

# 初始化环境依赖
echo "Installing dependencies..."

# 安装 Redis 示例（按需启用）
# sudo apt install redis-server -y
# 或者
# sudo yum install redis -y

echo "Done."
```

---

## 📄 main.py 示例

```python
def start():
    print("Starting Redis Manager External App...")

if __name__ == "__main__":
    start()
```

---

## ✅ 启动命令示例

在项目根目录执行：

```bash
python main.py
```

---