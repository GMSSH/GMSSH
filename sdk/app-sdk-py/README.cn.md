[![Latest Version](https://img.shields.io/pypi/v/simplejrpc.svg)](https://pypi.python.org/pypi/simplejrpc/)
[![Supported Python versions](https://img.shields.io/pypi/pyversions/simplejrpc.svg)](https://pypi.python.org/pypi/simplejrpc/)
[![Downloads](https://img.shields.io/pypi/dm/simplejrpc.svg)](https://pypi.org/project/simplejrpc/)

simplejrpc（app-sdk-py）
==========

[English](README.md) | **简体中文**

> **GMSSH（极目桌面）生态** · 官网 <https://www.gm.cn> · 文档 <https://doc.gm.cn> · 社区 <https://hub.gm.cn>
>
> GMSSH 应用 Python 后端 SDK：基于 Unix Socket + JSON-RPC 2.0 与桌面端通信。

## 1. 快速上手

### 1.1 安装 SDK

```bash
pip3 install simplejrpc
```

> 安装后即可在项目中导入并使用 SDK 功能。

---

## 2. 服务注册

### 2.1 注册 Socket 文件

SDK 通过 **Unix Socket 文件**（基于 JSON-RPC 2.0 协议）与后端进程通信，必须显式指定 `socket_path`，服务才会生成 `.sock` 文件并暴露通信能力。

```python
# main.py
from simplejrpc import ServerApplication

# 指定 Unix Socket 文件路径
socket_path = "/xxx/app.socket"

# 创建应用实例
app = ServerApplication(socket_path=socket_path)
```

> ✅ 提示：确保路径可写且不与其他服务冲突。

---

### 2.2 注册配置文件（日志系统）

初始化时指定日志配置文件路径即可启用日志；也可在应用启动后调用 `app.setup_logger(config_path)` 动态加载。

* 支持格式：**YAML**
* 未配置时默认使用 loguru 提供的简单日志

```python
# main.py
from simplejrpc import ServerApplication

sock_path = "/xxx/app.sock"
config_path = "/xxx/config.yaml"

# 初始化时带日志配置
app = ServerApplication(socket_path=sock_path, config_path=config_path)

# 或初始化后配置：
# app.setup_logger(config_path)
```

---

### 2.3 注册方法（业务路由）

注册方法后，函数即成为可通过 JSON-RPC 调用的接口。

#### ✅ 默认注册（用函数名）

未显式指定方法名时，系统默认使用函数名：

```python
from simplejrpc.response import jsonify


@app.route()  # 默认注册为 'hello'
async def hello():
    return jsonify(data="hello", msg="OK")
```

#### ✅ 显式注册（指定方法名）

用 `name` 参数指定方法名。建议与函数名保持一致，便于维护：

```python
from simplejrpc.response import jsonify


@app.route(name="hello")  # 显式指定方法名
async def hello():
    return jsonify(data="hello", msg="OK")
```

> ⚠️ 注意：`name` 与函数名不一致可能造成维护与调用混乱，建议保持一致。

---

## 3. 发起请求

### 3.1 请求与响应结构

#### ✅ 支持的参数形式

##### 📌 显式参数声明（推荐）

参数固定且命名明确时使用：

```python
@app.route(name="hello")
async def hello(lang, action):
    # 直接接收请求参数 lang 和 action
    ...
```

✅ 优点：类型清晰，便于校验与维护。

---

##### 📌 动态参数（\*args, \*\*kwargs）

参数数量可变或需通用处理时使用：

```python
@app.route(name="hello")
async def hello(*args, **kwargs):
    lang = kwargs.get("lang")
    action = kwargs.get("action")
    ...
```

---

### ✅ 标准响应格式

所有响应都应通过 `jsonify` 返回，包含 `code`、`data`、`msg` 字段：

```python
from simplejrpc.response import jsonify


@app.route(name="hello")
async def hello():
    return jsonify(code=400, data=True, msg="操作失败")
```

---

### ✅ 响应结构示例

#### 成功响应：

```json
{
    "jsonrpc": "2.0",
    "result": {
        "code": 200,
        "meta": {
            "endpoint": null,
            "close": 1
        },
        "data": [],
        "msg": "OK"
    },
    "id": 1
}
```

#### 错误响应：

```json
{
    "jsonrpc": "2.0",
    "result": {
        "code": 400,
        "meta": {
            "endpoint": null,
            "close": 1
        },
        "data": null,
        "msg": "expected value ['start', 'stop']"
    },
    "id": 1
}
```

---

### 3.2 启动服务

服务端示例：

```python
# main.py
import asyncio
from simplejrpc import ServerApplication
from simplejrpc.response import jsonify

socket_path = "/xxx/app.socket"
app = ServerApplication(socket_path)


@app.route(name="hello")
async def hello():
    return jsonify(data="hello", msg="OK")


if __name__ == "__main__":
    asyncio.run(app.run())
```

启动：

```bash
$ python3 main.py
```

---

### 3.3 调用接口（客户端测试）

用 SDK 的 `Request` 类测试：

```python
from simplejrpc import Request

# 必须与服务端 socket 路径一致
socket_path = "/xxx/app.sock"


def test_hello():
    method = "hello"
    params = {
        "lang": "zh-CN",
        "action": "start"
    }

    request = Request(socket_path)
    result = request.send_request(method, params)
    print("[recv] >", result)
```

---

## 4. 进阶功能

### 4.1 表单校验

SDK 集成 [`wtforms`](https://wtforms.readthedocs.io/en/3.1.x/)，提供简洁强大的校验体系并支持自定义校验器。

#### ✅ 示例：定义并使用表单类

```python
from simplejrpc.schemas import BaseForm, StrRangeValidator, simple
from simplejrpc.response import jsonify


# 自定义表单类，限制 'action' 只能为 "start" 或 "stop"
class TestForm(BaseForm):
    action = simple.StringField(
        validators=[StrRangeValidator(allows=["start", "stop"])]
    )


# 指定 socket 路径并注册应用
socket_path = "/xxx/app.socket"
app = ServerApplication(socket_path)


# 指定 'form' 即自动触发校验
@app.route(name="hello", form=TestForm)
async def hello(lang, action):
    return jsonify(data=[1, 2, 3], msg="OK")
```

---

#### ✅ 自定义校验器

继承 `BaseValidator` 实现自定义校验逻辑。

> 自定义校验器必须继承 `BaseValidator` 并实现 `validator(form, field)` 方法，只接收表单实例和字段实例两个参数。

```python
from simplejrpc import BaseValidator


class TestValidator(BaseValidator):
    def validator(self, form, field):
        # 自定义逻辑：值必须全大写
        if field.data and not field.data.isupper():
            raise ValueError("字段必须全大写")
```

在表单中使用：

```python
class TestForm(BaseForm):
    action = simple.StringField(
        validators=[
            StrRangeValidator(allows=["start", "stop"]),
            TestValidator()
        ]
    )
```

---

### 4.2 异常处理

框架内置异常处理系统：直接抛出异常，框架会格式化为标准 JSON-RPC 错误响应。

```python
from simplejrpc.exceptions import RPCException


@app.route(name="hello", form=TestForm)
async def hello(lang, action):
    raise RPCException("测试错误")  # 自动捕获并格式化
```

#### ✅ 自定义异常类

继承基础异常类扩展：

```python
class UnauthorizedError(RPCException):
    '''未授权'''

class ValidationError(RPCException):
    '''校验失败'''

class FileNotFoundError(RPCException):
    '''文件未找到'''

class ValueError(RPCException):
    '''值错误'''

class RuntimeError(RPCException):
    '''运行时错误'''
```

---

### 4.3 国际化（i18n）

SDK 用 `.ini` 文件支持多语言输出，由每次请求的 `lang` 参数驱动。

#### ✅ 初始化语言配置

项目根目录包含 `i18n/` 目录：

```bash
project/
├── main.py
└── i18n/
    ├── zh-CN.ini
    └── en.ini
```

用 `GI18n()` 初始化：

```python
GI18n(i18n_dir="i18n", lang="zh-CN")
```

> ⚠️ 每次请求应携带 `lang` 参数，框架据此设置语言。

---

#### ✅ 示例 1：基础翻译

**en.ini**

```ini
TEST_I18N = "test"
```

**Python 使用：**

```python
from simplejrpc import i18n

print(i18n.translate("TEST_I18N"))  # 输出: test
```

---

#### ✅ 示例 2：占位符翻译（带参数）

**en.ini**

```ini
TEST_I18N = "test{}"
```

**Python 使用：**

```python
from simplejrpc import i18n

print(i18n.translate_ctx("TEST_I18N", "i18n"))  # 输出: testi18n
```

---

#### ✅ 支持的语言

| 语言码       | 说明         | 地区/备注        |
| ------------ | ------------ | ---------------- |
| `en`         | 英语         | 默认            |
| `zh-CN`      | 简体中文     | 中国大陆        |
| `zh-TW`      | 繁体中文     | 台湾/香港       |
| `ja`         | 日语         | 日本            |
| `ru`         | 俄语         | 俄罗斯          |

---

### 4.4 中间件

框架支持中间件，在请求处理前后执行通用逻辑（日志、鉴权、性能分析等）。

```python
from simplejrpc import ServerApplication
from simplejrpc.interfaces import RPCMiddleware


class CustomMiddleware(RPCMiddleware):
    def process_request(self, request, context):
        print("[前置] 收到请求:", request)
        return request

    def process_response(self, response, context):
        print("[后置] 发出响应:", response)
        return response


# 注册中间件
app = ServerApplication("/xxx/app.sock")
app.middleware(CustomMiddleware())
```

---

## 反馈

在 [Gitee](https://gitee.com/gmssh_1/simplerpc.git) 提交 issue / fork 项目。

在 [GitHub](https://github.com/GMSSH/app-sdk-py.git) 提交 issue / fork 项目。
