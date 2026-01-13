
# 开发规范

---

## 一、开发规范

## 1.1 项目结构建议

为提升开发效率与代码可维护性，推荐按照如下目录组织外置应用项目：

```
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

* 您可以使用我们提供的命令行生成脚手架的方式

```bash
# 安装脚手架工具
pip3 install gmscaffold

# 使用 CLI 快速生成项目结构
gmcli gm_app create_app
```

* 或参考：[GitHub 脚手架仓库](https://github.com/GMSSH/app-sample-py.git)。
```bash
git clone https://github.com/GMSSH/app-sample-py.git
```
---

## 1.2 项目关键文件说明

| 文件名                | 描述                                |
| ------------------ | --------------------------------- |
| `config.yaml`      | 配置中心，包含日志等级、服务端口、健康检查路径等          |
| `install.sh`       | 初始化脚本，如安装 Redis、环境检测等             |
| `main.py`          | 程序入口，负责服务注册与启动                    |
| `Makefile`         | 可定义常用命令（构建、运行等）                   |
| `requirements.txt` | 推荐依赖管理方式之一，可替代为 Poetry 使用更现代的依赖控制 |

---

## 1.3 服务注册机制（重点

外置应用必须在启动时完成服务注册，否则 GA 服务中心将无法识别与转发请求。

#### 注册协议

* 通信方式：**Unix-Socket + JSON-RPC**
* 通信路径：`/.__gmssh/tmp/rpc.sock`
* 推荐使用提供的 SDK：[`simplejrpc`](https://pypi.org/project/simplejrpc/)

#### python示例代码：


<tabs>
    <tab title="python">

```python

from simplejrpc.client import GmRequest

def register_server():
    request = GmRequest()  # 内部使用 /.__gmssh/tmp/rpc.sock
    response = request.send_request(
        method="register_server",
        params={
            "port": "",
            "type": "socket",
            "healthPath": "ping",
            "healthTimeout": 5,
            "metaData": {
                "orgName": "wmm",
                "appName": "redis",
                "version": "1.0.0"
            }
        }
    )
    print("[recv] >", response)
```
</tab>
</tabs>


#### 参数说明：

| 字段名                    | 描述                         | 示例值               | 必填 |
| ---------------------- | -------------------------- | ----------------- | -- |
| `method`               | 方法名（固定为 `register_server`） | `register_server` | ✅  |
| `params.port`          | 服务端口，Socket 模式可留空          | `""`              | ❌  |
| `params.type`          | 服务类型（固定为 `socket`）         | `socket`          | ✅  |
| `params.healthPath`    | 健康检查路径                     | `ping`            | ✅  |
| `params.healthTimeout` | 健康检查超时时间（秒）                | `5`               | ✅  |
| `metaData.orgName`     | 开发组织名称                     | `wmm`             | ✅  |
| `metaData.appName`     | 应用名称                       | `redis`           | ✅  |
| `metaData.version`     | 应用版本                       | `1.0.0`           | ✅  |

---

## 1.4 服务中心心跳检测机制

* 注册成功后，GA 服务中心将接管应用健康监控；
* 默认每 **5 秒** 发起一次健康检测；
* 健康检查失败（如无响应）将自动下线服务；
* 注意：**请勿删除 `/.__gmssh/tmp/rpc.sock` 文件，否则通信将中断**。

---

## 1.5 服务开发方式选项

您可以选择以下任意方式开发外置服务：

| 模式        | 描述                             | 是否需注册 |
| --------- |--------------------------------| ----- |
| SDK 模式    | 推荐使用 `simplejrpc`，封装了与GA服务交互方法 | ✅     |
| HTTP 服务模式 | 使用自定义 HTTP 服务，自行实现注册逻辑         | ✅     |

无论选择哪种模式，注册行为必须通过 **Unix-Socket + JSON-RPC** 完成。

如需手动实现注册，请确保：
* 使用符合 JSON-RPC 标准的格式

## 1.6非SDK实现 {: id="6"}

> 这是一个demo级别的客户端，使用 Unix-socket+jsonrpc进行交互的

<tabs>
    <tab title="python">

```python

import json
import socket


def send_jsonrpc_request(method, params, request_id=1):
    # 连接到 Unix 套接字
    sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    sock.connect('/xxx/app.sock')
    
    # 构造 JSON-RPC 请求负载
    payload = json.dumps({
        "jsonrpc": "2.0",
        "method": method,
        "params": params,
        "id": request_id
    })
    payload_bytes = payload.encode('utf-8')
    content_length = len(payload_bytes)

    # 根据 VSCodeObjectCodec 格式构造消息
    header = f"Content-Length: {content_length}\r\n\r\n".encode('utf-8')
    message = header + payload_bytes

    # 发送请求
    sock.sendall(message)

    # 读取响应头
    header_data = b""
    # 持续读取直到遇到 "\r\n\r\n"（表明头部结束）
    while b"\r\n\r\n" not in header_data:
        chunk = sock.recv(1024)
        if not chunk:
            break
        header_data += chunk
    # 分离头部和已读取的部分数据
    header_section, _, remaining = header_data.partition(b"\r\n\r\n")
    # 解析 header 获取 Content-Length
    content_length_val = None
    for line in header_section.decode('utf-8').split("\r\n"):
        if line.lower().startswith("content-length:"):
            content_length_val = int(line.split(":", 1)[1].strip())
            break
    if content_length_val is None:
        sock.close()
        raise ValueError("响应中没有找到 Content-Length 头")

    # 根据 Content-Length 读取完整响应体
    body = remaining
    while len(body) < content_length_val:
        chunk = sock.recv(1024)
        if not chunk:
            break
        body += chunk

    sock.close()

    # 解析 JSON 响应
    response = json.loads(body.decode('utf-8'))
    return response
if __name__ == "__main__":
    response = send_jsonrpc_request("hello", {}, request_id=1)
    print("结果：", response)
```
</tab>
</tabs>

> 这是一个demo级别的服务端，使用 Unix-socket+jsonrpc

<tabs>
    <tab title="python">

```python
import socket
import json
import os
from typing import Dict, Any

class JSONRPCServer:
    def __init__(self, socket_path: str):
        self.socket_path = socket_path  # 套接字文件路径
        self.methods: Dict[str, Any] = {}  # 存储注册的方法
        
    def register_method(self, name: str, func: callable):
        """注册一个JSON-RPC方法处理器"""
        self.methods[name] = func
    
    def handle_request(self, request: Dict) -> Dict:
        """处理JSON-RPC请求并返回响应"""
        try:
            # 验证JSON-RPC版本
            if request.get("jsonrpc") != "2.0":
                return {
                    "jsonrpc": "2.0",
                    "error": {"code": -32600, "message": "无效请求"},
                    "id": request.get("id")
                }
            
            method = request.get("method")
            # 检查方法是否存在
            if method not in self.methods:
                return {
                    "jsonrpc": "2.0",
                    "error": {"code": -32601, "message": "方法未找到"},
                    "id": request.get("id")
                }
            
            # 调用注册的方法
            params = request.get("params", {})
            result = self.methods[method](params)
            
            # 返回成功响应
            return {
                "jsonrpc": "2.0",
                "result": result,
                "id": request.get("id")
            }
            
        except Exception as e:
            # 返回错误响应
            return {
                "jsonrpc": "2.0",
                "error": {"code": -32603, "message": str(e)},
                "id": request.get("id")
            }
    
    def start(self):
        """启动服务器并监听连接"""
        # 如果套接字文件已存在则删除
        try:
            os.unlink(self.socket_path)
        except OSError:
            if os.path.exists(self.socket_path):
                raise
        
        # 创建Unix域套接字
        sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        sock.bind(self.socket_path)
        sock.listen(1)
        
        print(f"服务器已启动，监听 {self.socket_path}")
        
        try:
            while True:
                connection, _ = sock.accept()
                try:
                    # 读取头部直到遇到双CRLF
                    header_data = b""
                    while b"\r\n\r\n" not in header_data:
                        chunk = connection.recv(1024)
                        if not chunk:
                            break
                        header_data += chunk
                    
                    # 分离头部和可能已读取的正文数据
                    header_section, _, remaining = header_data.partition(b"\r\n\r\n")
                    
                    # 解析Content-Length
                    content_length = 0
                    for line in header_section.decode('utf-8').split("\r\n"):
                        if line.lower().startswith("content-length:"):
                            content_length = int(line.split(":", 1)[1].strip())
                            break
                    
                    # 如果需要，读取剩余的正文
                    body = remaining
                    while len(body) < content_length:
                        chunk = connection.recv(1024)
                        if not chunk:
                            break
                        body += chunk
                    
                    # 解析JSON请求
                    request = json.loads(body.decode('utf-8'))
                    
                    # 处理请求
                    response = self.handle_request(request)
                    
                    # 准备响应
                    response_json = json.dumps(response)
                    response_bytes = response_json.encode('utf-8')
                    response_headers = f"Content-Length: {len(response_bytes)}\r\n\r\n".encode('utf-8')
                    
                    # 发送响应
                    connection.sendall(response_headers + response_bytes)
                    
                except json.JSONDecodeError:
                    error_response = {
                        "jsonrpc": "2.0",
                        "error": {"code": -32700, "message": "解析错误"},
                        "id": None
                    }
                    error_json = json.dumps(error_response)
                    connection.sendall(f"Content-Length: {len(error_json)}\r\n\r\n{error_json}".encode('utf-8'))
                finally:
                    connection.close()
        finally:
            sock.close()
            try:
                os.unlink(self.socket_path)
            except OSError:
                pass

# 创建一个服务端实例，指定套接字路径
server = JSONRPCServer("/xxx/tmp.socket")

# 注册一个示例方法
def hello(params):
    name = params.get("name", "World")
    return f"Hello, {name}!"

server.register_method("greet", hello)

# 启动服务器
server.start()

```
</tab>
</tabs>


## 1.7 服务说明

#### **进程启动机制**
- **启动方式**：
  ```bash
  nohup ./main > nohup.out 2>&1 &
  ```  
  - ✅ **后台运行**（`&`）
  - ✅ **终端信号屏蔽**（`nohup` 忽略 `SIGHUP`）
  - ✅ **完整日志记录**（`stdout` + `stderr` 重定向至 `nohup.out`）

- **进程元数据存储**：
  - 📌 **PID、PPID、启动时间**（防止 PID 复用误判）

---  

#### **进程终止机制**

**终止策略（强制优先）**：
  - **直接发送 `SIGKILL`（-9）**，立即释放资源
    ```bash
    kill -9 ${PID}  # 无条件终止，避免僵尸进程
    ```



## 1.8 Socket规范

在使用Unix Socket+JSON-RPC开发外置应用时，必须特别注意socket文件的位置。标准规范要求：

* **位置**: 必须放置于 `/.__gmssh/plugin/组织名/应用名/tmp/app.sock` 目录中。

* **路径说明**:

  * **组织名**: 指应用的开发者或作者的名称。
  * **应用名**: 指的是你正在开发的应用名称，例如 `redis`、`mysql` 等等。
  * **命名规范**: socket文件的名称必须为 `app.sock`，以便保持统一和规范。

* **优化要求**: 请确保socket文件放置在应用的 `tmp` 目录中，并且按照此路径进行存储。这有助于统一管理和提升应用的可维护性。

#### 示例路径:

```
/.__gmssh/plugin/johndoe/redis/tmp/app.sock
```

这样可以确保应用的socket文件能够被GA服务正确识别和使用，具体详见 [请求响应](请求响应.md)

> 需要注意同一个组织下应用名称必须是唯一的
