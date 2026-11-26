[![Latest Version](https://img.shields.io/pypi/v/gmscaffold.svg)](https://pypi.python.org/pypi/gmscaffold/)
[![Supported Python versions](https://img.shields.io/pypi/pyversions/gmscaffold.svg)](https://pypi.python.org/pypi/gmscaffold/)
[![Downloads](https://img.shields.io/pypi/dm/gmscaffold.svg)](https://pypi.org/project/gmscaffold/)

gmscaffold
==========

[English](README.md) | **简体中文**

> **GMSSH（极目桌面）生态** · 官网 <https://www.gm.cn> · 文档 <https://doc.gm.cn> · 社区 <https://hub.gm.cn>

## 1. 准备开发环境

### 1.1 简介

脚手架工具帮你**快速生成项目的基础结构**：一条命令自动创建项目目录、配置文件和初始化代码，省去手工搭建。

在本体系中，使用命令行工具 `gmcli` 可以快速创建标准化的应用工程结构，高效进入开发。

---

### 1.2 环境准备

#### 安装 Python 环境

* **安装 Python 3.10 或更高版本**
  前往 [Python 官网](https://www.python.org/downloads/) 下载并安装适合你操作系统的版本。

  > ⚠️ 当前脚手架工具仅支持 **Python 3.10 及以上**。

* **安装 pip（Python 包管理器）**
  多数系统已预装 pip。若未安装，请参考 [pip 官方文档](https://pip.pypa.io/en/stable/installation/)。

---

### 1.3 安装脚手架工具

通过 pip 安装 CLI 工具：

```bash
pip3 install gmscaffold
```

安装完成后即可使用 `gmcli` 命令创建项目。

---

## 2. 用脚手架生成项目

### 2.1 创建新的脚手架项目

执行以下命令生成项目结构：

```bash
gmcli gm_app create_app
```

> 💡 提示：项目会创建在**当前工作目录**，请先 `cd` 到目标路径。

---

#### 示例输出：

```text
➜  ~ gmcli gm_app create_app

     _/_/_/  _/      _/    _/_/_/    _/_/_/  _/    _/
  _/        _/_/  _/_/  _/        _/        _/    _/
 _/  _/_/  _/  _/  _/    _/_/      _/_/    _/_/_/_/
_/    _/  _/      _/        _/        _/  _/    _/
 _/_/_/  _/      _/  _/_/_/    _/_/_/    _/    _/

Please enter the project name: example
Please enter the author name: xxx
Please enter contact email: xxx@163.com
Please enter the project version: 0.0.1
Please enter the project description: example
DEBUG:GmScaffold:[example] create_gm_app successfully
```

---

### 2.2 命令参数

| 参数                 | 说明               |
| -------------------- | ------------------ |
| `gm_app create_app`  | 创建新的脚手架项目 |

---

### 2.3 项目结构

执行完成后生成如下结构：

```
example/
├── app/
│   ├── consts/           # 常量（状态码、配置键等）
│   ├── i18n/             # 国际化（默认中英双语）
│   ├── middlewares/      # 中间件模块（鉴权、请求处理等）
│   ├── schemas/          # 数据校验模型
│   ├── services/         # 业务逻辑
│   ├── utils/            # 工具函数（日志等）
│   └── __init__.py
├── config.yaml           # 配置文件（日志、端口等）
├── install.sh            # 可选安装脚本（如 Redis、MySQL）
├── main.py               # 应用入口
├── makefile              # 构建或打包脚本
├── requirements.txt      # Python 依赖清单
├── .gitignore            # Git 忽略配置
├── ReadMe.md             # 项目文档
```

---

#### 关键文件说明：

* `main.py`：应用入口
* `install.sh`：可选的外部服务安装脚本（如 Redis、MySQL）
* `config.yaml`：集中配置管理（端口、日志、数据库等）
* `requirements.txt`：Python 依赖清单（也可用 `poetry`）

---

### 2.4 使用与扩展

生成的项目是通用 API 服务模板，可扩展到各种业务场景：

* ✅ 可配合 [**simplejrpc SDK**](https://pypi.org/project/simplejrpc/) 快速构建 JSON-RPC 服务
* ✅ 也支持独立使用，自定义业务逻辑与结构
