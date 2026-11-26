[![Latest Version](https://img.shields.io/pypi/v/gmscaffold.svg)](https://pypi.python.org/pypi/gmscaffold/)
[![Supported Python versions](https://img.shields.io/pypi/pyversions/gmscaffold.svg)](https://pypi.python.org/pypi/gmscaffold/)
[![Downloads](https://img.shields.io/pypi/dm/gmscaffold.svg)](https://pypi.org/project/gmscaffold/)

gmscaffold
==========

[English](README.md) | [简体中文](README.cn.md)

> **GMSSH（极目桌面）生态** · 官网 <https://www.gm.cn> · 文档 <https://doc.gm.cn> · 社区 <https://hub.gm.cn>
>
> GMSSH 应用开发脚手架：一行命令生成标准化微应用工程结构。


## 1. Preparing the Development Environment

### 1.1 Introduction

A scaffolding tool helps you **quickly generate the basic structure of a project**. With a simple command, it can automatically create project directories, configuration files, and initialization code, saving you the hassle of manual setup.

In this system, you can use the command-line tool `gmcli` to quickly create a standardized application project structure, enabling you to start development efficiently.

---

### 1.2 Environment Setup

#### Install Python Environment

* **Install Python 3.10 or later**
  Visit the [official Python website](https://www.python.org/downloads/) to download and install the appropriate version for your operating system.

  > ⚠️ Currently, the scaffolding tool supports **Python 3.10 and above** only.

* **Install pip (Python package manager)**
  Most systems come with pip pre-installed. If not, please refer to the [official pip documentation](https://pip.pypa.io/en/stable/installation/) for installation instructions.

---

### 1.3 Install the Scaffolding Tool

Install the CLI tool via pip:

```bash
pip3 install gmscaffold
```

Once installed, you can use the `gmcli` command to create your project.

---

## 2. Generate a Project with the Scaffold

### 2.1 Create a New Scaffolded Project

Run the following command to generate the project structure:

```bash
gmcli gm_app create_app
```

> 💡 Tip: The project will be created in the **current working directory**, so make sure you `cd` into the desired path first.

---

#### Example Output:

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

### 2.2 Command Parameters

| Parameter           | Description                     |
| ------------------- | ------------------------------- |
| `gm_app create_app` | Create a new scaffolded project |

---

### 2.3 Project Structure

After execution, the following structure will be generated:

```
example/
├── app/
│   ├── consts/           # Constants (e.g., status codes, config keys)
│   ├── i18n/             # Internationalization (default: English & Chinese)
│   ├── middlewares/      # Middleware modules (e.g., auth, request handlers)
│   ├── schemas/          # Data validation schemas
│   ├── services/         # Business logic
│   ├── utils/            # Utility functions (e.g., logging, tools)
│   └── __init__.py
├── config.yaml           # Configuration file (logs, ports, etc.)
├── install.sh            # Optional install script (e.g., Redis, MySQL)
├── main.py               # Application entry point
├── makefile              # Build or packaging script
├── requirements.txt      # Python dependencies list
├── .gitignore            # Git ignore settings
├── ReadMe.md             # Project documentation
```

---

#### Key File Descriptions:

* `main.py`: Entry point of the application
* `install.sh`: Optional script to install external services (e.g., Redis, MySQL)
* `config.yaml`: Centralized configuration management (port, logs, DB, etc.)
* `requirements.txt`: Python dependencies list (can also use `poetry`)

---

### 2.4 Usage & Extension

The generated project serves as a general-purpose API service template and can be extended for various business scenarios:

* ✅ Can be used with [**simplejrpc SDK**](https://pypi.org/project/simplejrpc/) to quickly build JSON-RPC services
* ✅ Also supports standalone use with custom business logic and structure
