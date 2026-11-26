[![Latest Version](https://img.shields.io/pypi/v/simplejrpc.svg)](https://pypi.python.org/pypi/simplejrpc/)
[![Supported Python versions](https://img.shields.io/pypi/pyversions/simplejrpc.svg)](https://pypi.python.org/pypi/simplejrpc/)
[![Downloads](https://img.shields.io/pypi/dm/simplejrpc.svg)](https://pypi.org/project/simplejrpc/)

simplejrpc
==========


# 1. Getting Started Quickly

### 1.1 Install SDK

Install the SDK using pip:

```bash
pip3 install simplejrpc
```

> Once installed, you can import and use the SDK functions in your project.

---

# 2. Service Registration

### 2.1 Register a Socket File

Since the SDK communicates with the backend process via a **Unix Socket file** (based on the JSON-RPC 2.0 protocol), you must explicitly specify the `socket_path` to ensure the service generates a `.sock` file and exposes communication capabilities.

```python
# main.py
from simplejrpc import ServerApplication

# Specify the Unix Socket file path
socket_path = "/xxx/app.socket"

# Create the application instance
app = ServerApplication(socket_path=socket_path)
```

> ✅ Tip: Make sure the path is writable and does not conflict with other services.

---

### 2.2 Register a Configuration File (Logging System)

To enable logging, specify the path to a logger config file during initialization. You can also call `app.setup_logger(config_path)` after the app starts to load the logger config dynamically.

* Supported format: **YAML**
* If not configured, a simple logger provided by loguru is used by default.

```python
# main.py
from simplejrpc import ServerApplication

sock_path = "/xxx/app.sock"
config_path = "/xxx/config.yaml"

# Initialize the app with logger configuration
app = ServerApplication(socket_path=sock_path, config_path=config_path)

# Alternatively, configure the logger after initialization:
# app.setup_logger(config_path)
```

---

### 2.3 Register Methods (Business Routes)

Registering methods exposes functions as interfaces accessible via JSON-RPC.

#### ✅ Default Registration (Using Function Name)

If the method name is not explicitly specified, the system will use the function name by default:

```python
from simplejrpc.response import jsonify


@app.route()  # Registered as 'hello' by default
async def hello():
    return jsonify(data="hello", msg="OK")
```

#### ✅ Explicit Registration (Specify Method Name)

You can also specify a method name using the `name` parameter. It’s recommended to keep it consistent with the function name for maintainability:

```python
from simplejrpc.response import jsonify


@app.route(name="hello")  # Explicitly set method name
async def hello():
    return jsonify(data="hello", msg="OK")
```

> ⚠️ Note: If `name` differs from the function name, it may cause confusion during maintenance or calls. Consistent naming is advised.

---

# 3. Making Requests

---

## 3.1 Request and Response Structure

### ✅ Supported Parameter Formats

The framework supports two ways of receiving request parameters:

#### 📌 Explicit Parameter Declaration (Recommended)

Use when parameters are fixed and clearly named:

```python
@app.route(name="hello")
async def hello(lang, action):
    # Directly receive request parameters lang and action
    ...
```

✅ Advantage: Clear typing, easier validation and maintenance.

---

#### 📌 Dynamic Parameter Handling (\*args, \*\*kwargs)

Use when the number of parameters is variable or needs generic processing:

```python
@app.route(name="hello")
async def hello(*args, **kwargs):
    lang = kwargs.get("lang")
    action = kwargs.get("action")
    ...
```

---

### ✅ Standard Response Format

All responses should be returned using `jsonify`, containing `code`, `data`, and `msg` fields:

```python
from simplejrpc.response import jsonify


@app.route(name="hello")
async def hello():
    return jsonify(code=400, data=True, msg="Operation failed")
```

---

### ✅ Example Response Structure

#### Success Response:

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

#### Error Response:

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

## 3.2 Start the Service

Example server code:

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

Start the service:

```bash
$ python3 main.py
```

---

## 3.3 Call Interface (Client Test)

Use the SDK's `Request` class for testing:

```python
from simplejrpc import Request

# Must match the socket path used by the server
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

# 4. Making Requests

---

## 4.1 Form Validation

The SDK integrates [`wtforms`](https://wtforms.readthedocs.io/en/3.1.x/) to provide a simple yet powerful validation system with support for custom validators.

### ✅ Example: Define and Use a Form Class

```python
from simplejrpc.schemas import BaseForm, StrRangeValidator, simple
from simplejrpc.response import jsonify


# Custom form class, restricts 'action' to "start" or "stop"
class TestForm(BaseForm):
    action = simple.StringField(
        validators=[StrRangeValidator(allows=["start", "stop"])]
    )


# Specify socket path and register app
socket_path = "/xxx/app.socket"
app = ServerApplication(socket_path)


# Specify 'form' to trigger validation automatically
@app.route(name="hello", form=TestForm)
async def hello(lang, action):
    return jsonify(data=[1, 2, 3], msg="OK")
```

---

### ✅ Custom Validators

You can create custom validation logic by extending `BaseValidator`.

> A custom validator class must inherit from `BaseValidator` and implement the `validator(form, field)` method. Only two parameters are accepted: the form instance and the field instance.

```python
from simplejrpc import BaseValidator


class TestValidator(BaseValidator):
    def validator(self, form, field):
        # Custom logic: value must be uppercase
        if field.data and not field.data.isupper():
            raise ValueError("Field must be uppercase")
```

Use the custom validator in your form:

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

## 4.2 Exception Handling

The framework includes a built-in exception handling system. Just raise exceptions, and the framework will format them into standard JSON-RPC error responses.

### ✅ Example

```python
from simplejrpc.exceptions import RPCException


@app.route(name="hello", form=TestForm)
async def hello(lang, action):
    raise RPCException("Test error")  # Automatically caught and formatted
```

### ✅ Custom Exception Classes

The framework provides a base exception class that can be extended:

```python
class UnauthorizedError(RPCException):
    """Unauthorized"""

class ValidationError(RPCException):
    """Validation failed"""

class FileNotFoundError(RPCException):
    """File not found"""

class ValueError(RPCException):
    """Value error"""

class RuntimeError(RPCException):
    """Runtime error"""
```

---

## 4.3 Internationalization (i18n)

The SDK supports multilingual output using `.ini` files, driven by the `lang` parameter in each request.

### ✅ Initialize Language Configuration

Your project root should include an `i18n/` directory:

```bash
project/
├── main.py
└── i18n/
    ├── zh-CN.ini
    └── en.ini
```

Initialize with `GI18n()`:

```python
GI18n(i18n_dir="i18n", lang="zh-CN")
```

> ⚠️ Each request should include a `lang` parameter, which the framework uses to set the language.

---

### ✅ Example 1: Basic Translation

**en.ini**

```ini
TEST_I18N = "test"
```

**Python usage:**

```python
from simplejrpc import i18n

print(i18n.translate("TEST_I18N"))  # Output: test
```

---

### ✅ Example 2: Placeholder Translation (Parameterized)

**en.ini**

```ini
TEST_I18N = "test{}"
```

**Python usage:**

```python
from simplejrpc import  i18n

print(i18n.translate_ctx("TEST_I18N", "i18n"))  # Output: testi18n
```

---

### ✅ Supported Languages

| Language Code | Description         | Region/Note      |
| ------------- | ------------------- | ---------------- |
| `en`          | English             | Default          |
| `zh-CN`       | Simplified Chinese  | Mainland China   |
| `zh-TW`       | Traditional Chinese | Taiwan/Hong Kong |
| `ja`          | Japanese            | Japan            |
| `ru`          | Russian             | Russia           |

---

## 4.4 Middleware Support

The framework supports middleware for handling common logic before and after request processing, such as logging, authentication, or performance analysis.

### ✅ Example Usage

```python
from simplejrpc import ServerApplication
from simplejrpc.interfaces import RPCMiddleware


class CustomMiddleware(RPCMiddleware):
    def process_request(self, request, context):
        print("[Before] Incoming request:", request)
        return request

    def process_response(self, response, context):
        print("[After] Outgoing response:", response)
        return response


# Register middleware
app = ServerApplication("/xxx/app.sock")
app.middleware(CustomMiddleware())
```

---

Would you like this translated version exported into a PDF or Markdown file?


Feedback
--------

Open a ticket / fork the project on [Gitee](https://gitee.com/gmssh_1/simplerpc.git).

Open a ticket / fork the project on [Github](https://github.com/GMSSH/app-sdk-py.git).

Here is the fully translated version of your documentation from Chinese to English:

---
