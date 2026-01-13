
# Development Guidelines

---

## 1. Development Guidelines

## 1.1 Recommended Project Structure

To improve development efficiency and code maintainability, it is recommended to organize external application projects as follows:

```
example/
├── app/
│   ├── consts/           # Constants (e.g., status codes, configuration items)
│   ├── i18n/             # Internationalization module (default: Chinese and English)
│   ├── middlewares/      # Middlewares (e.g., authentication, request interception)
│   ├── schemas/          # Input validation (Pydantic or custom rules)
│   ├── services/         # Core business logic
│   ├── utils/            # Utility modules (e.g., logging, RPC calls)
│   └── __init__.py
├── config.yaml           # Application configuration (logging, port, metadata, etc.)
├── install.sh            # Installation script (e.g., Redis, dependencies)
├── main.py               # Application entry point with service registration and startup logic
├── Makefile              # Build/package command definitions
├── requirements.txt      # Python dependency list (can use Poetry alternatively)
├── .gitignore            # Git ignore list
├── README.md             # Project documentation
```

* You can use our CLI tool to generate the scaffold:

```bash
# Install the scaffold tool
pip3 install gmscaffold

# Use CLI to quickly generate project structure
gmcli gm_app create_app
```

* Or refer to: [GitHub Scaffold Repository](https://github.com/GMSSH/app-sample-py.git).
```bash
git clone https://github.com/GMSSH/app-sample-py.git
```
---

## 1.2 Key Project Files

| Filename           | Description                                                       |
| ------------------ | ----------------------------------------------------------------- |
| `config.yaml`      | Configuration center, includes log level, service port, health check path, etc. |
| `install.sh`       | Initialization script, e.g., Redis installation, environment detection |
| `main.py`          | Entry point, responsible for service registration and startup      |
| `Makefile`         | Define common commands (build, run, etc.)                          |
| `requirements.txt` | Recommended dependency management; can be replaced with Poetry for modern dependency control |

---

## 1.3 Service Registration Mechanism (Important)

External applications must complete service registration at startup; otherwise, the GA Service Center cannot recognize or forward requests.

#### Registration Protocol

* Communication method: **Unix-Socket + JSON-RPC**
* Communication path: `/.__gmssh/tmp/rpc.sock`
* Recommended SDK: [`simplejrpc`](https://pypi.org/project/simplejrpc/)

#### Python Example:


<tabs>
    <tab title="python">

```python

from simplejrpc.client import GmRequest

def register_server():
    request = GmRequest()  # Uses /.__gmssh/tmp/rpc.sock internally
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


#### Parameter Description:

| Field Name             | Description                              | Example           | Required |
| ---------------------- | ---------------------------------------- | ----------------- | -------- |
| `method`               | Method name (fixed as `register_server`) | `register_server` | ✅        |
| `params.port`          | Service port, can be empty in Socket mode | `""`              | ❌        |
| `params.type`          | Service type (fixed as `socket`)          | `socket`          | ✅        |
| `params.healthPath`    | Health check path                         | `ping`            | ✅        |
| `params.healthTimeout` | Health check timeout (seconds)            | `5`               | ✅        |
| `metaData.orgName`     | Developer organization name               | `wmm`             | ✅        |
| `metaData.appName`     | Application name                          | `redis`           | ✅        |
| `metaData.version`     | Application version                       | `1.0.0`           | ✅        |

---

## 1.4 Service Center Heartbeat Detection

* After successful registration, the GA Service Center takes over application health monitoring;
* Default health check interval: every **5 seconds**;
* Failed health checks (e.g., no response) will automatically take the service offline;
* Note: **Do not delete the `/.__gmssh/tmp/rpc.sock` file, otherwise communication will be interrupted**.

---

## 1.5 Service Development Options

You can choose any of the following methods to develop external services:

| Mode         | Description                                              | Registration Required |
| ------------ | -------------------------------------------------------- | --------------------- |
| SDK Mode     | Recommended: use `simplejrpc`, encapsulates GA service interaction methods | ✅                     |
| HTTP Service | Use custom HTTP service, implement registration logic yourself | ✅                     |

Regardless of the mode chosen, registration must be completed via **Unix-Socket + JSON-RPC**.

If implementing registration manually, ensure:
* Use JSON-RPC compliant format

## 1.6 Non-SDK Implementation {: id="6"}

> This is a demo-level client using Unix-socket + JSON-RPC for interaction

<tabs>
    <tab title="python">

```python

import json
import socket


def send_jsonrpc_request(method, params, request_id=1):
    # Connect to Unix socket
    sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    sock.connect('/xxx/app.sock')
    
    # Construct JSON-RPC request payload
    payload = json.dumps({
        "jsonrpc": "2.0",
        "method": method,
        "params": params,
        "id": request_id
    })
    payload_bytes = payload.encode('utf-8')
    content_length = len(payload_bytes)

    # Construct message according to VSCodeObjectCodec format
    header = f"Content-Length: {content_length}\r\n\r\n".encode('utf-8')
    message = header + payload_bytes

    # Send request
    sock.sendall(message)

    # Read response header
    header_data = b""
    # Keep reading until encountering "\r\n\r\n" (indicating header end)
    while b"\r\n\r\n" not in header_data:
        chunk = sock.recv(1024)
        if not chunk:
            break
        header_data += chunk
    # Separate header and already read data
    header_section, _, remaining = header_data.partition(b"\r\n\r\n")
    # Parse header to get Content-Length
    content_length_val = None
    for line in header_section.decode('utf-8').split("\r\n"):
        if line.lower().startswith("content-length:"):
            content_length_val = int(line.split(":", 1)[1].strip())
            break
    if content_length_val is None:
        sock.close()
        raise ValueError("Content-Length header not found in response")

    # Read complete response body according to Content-Length
    body = remaining
    while len(body) < content_length_val:
        chunk = sock.recv(1024)
        if not chunk:
            break
        body += chunk

    sock.close()

    # Parse JSON response
    response = json.loads(body.decode('utf-8'))
    return response
if __name__ == "__main__":
    response = send_jsonrpc_request("hello", {}, request_id=1)
    print("Result:", response)
```
</tab>
</tabs>

> This is a demo-level server using Unix-socket + JSON-RPC

<tabs>
    <tab title="python">

```python
import socket
import json
import os
from typing import Dict, Any

class JSONRPCServer:
    def __init__(self, socket_path: str):
        self.socket_path = socket_path  # Socket file path
        self.methods: Dict[str, Any] = {}  # Store registered methods
        
    def register_method(self, name: str, func: callable):
        """Register a JSON-RPC method handler"""
        self.methods[name] = func
    
    def handle_request(self, request: Dict) -> Dict:
        """Process JSON-RPC request and return response"""
        try:
            # Validate JSON-RPC version
            if request.get("jsonrpc") != "2.0":
                return {
                    "jsonrpc": "2.0",
                    "error": {"code": -32600, "message": "Invalid request"},
                    "id": request.get("id")
                }
            
            method = request.get("method")
            # Check if method exists
            if method not in self.methods:
                return {
                    "jsonrpc": "2.0",
                    "error": {"code": -32601, "message": "Method not found"},
                    "id": request.get("id")
                }
            
            # Call registered method
            params = request.get("params", {})
            result = self.methods[method](params)
            
            # Return success response
            return {
                "jsonrpc": "2.0",
                "result": result,
                "id": request.get("id")
            }
            
        except Exception as e:
            # Return error response
            return {
                "jsonrpc": "2.0",
                "error": {"code": -32603, "message": str(e)},
                "id": request.get("id")
            }
    
    def start(self):
        """Start server and listen for connections"""
        # Delete socket file if it already exists
        try:
            os.unlink(self.socket_path)
        except OSError:
            if os.path.exists(self.socket_path):
                raise
        
        # Create Unix domain socket
        sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        sock.bind(self.socket_path)
        sock.listen(1)
        
        print(f"Server started, listening on {self.socket_path}")
        
        try:
            while True:
                connection, _ = sock.accept()
                try:
                    # Read header until double CRLF
                    header_data = b""
                    while b"\r\n\r\n" not in header_data:
                        chunk = connection.recv(1024)
                        if not chunk:
                            break
                        header_data += chunk
                    
                    # Separate header and possibly already read body data
                    header_section, _, remaining = header_data.partition(b"\r\n\r\n")
                    
                    # Parse Content-Length
                    content_length = 0
                    for line in header_section.decode('utf-8').split("\r\n"):
                        if line.lower().startswith("content-length:"):
                            content_length = int(line.split(":", 1)[1].strip())
                            break
                    
                    # Read remaining body if needed
                    body = remaining
                    while len(body) < content_length:
                        chunk = connection.recv(1024)
                        if not chunk:
                            break
                        body += chunk
                    
                    # Parse JSON request
                    request = json.loads(body.decode('utf-8'))
                    
                    # Process request
                    response = self.handle_request(request)
                    
                    # Prepare response
                    response_json = json.dumps(response)
                    response_bytes = response_json.encode('utf-8')
                    response_headers = f"Content-Length: {len(response_bytes)}\r\n\r\n".encode('utf-8')
                    
                    # Send response
                    connection.sendall(response_headers + response_bytes)
                    
                except json.JSONDecodeError:
                    error_response = {
                        "jsonrpc": "2.0",
                        "error": {"code": -32700, "message": "Parse error"},
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

# Create a server instance with socket path
server = JSONRPCServer("/xxx/tmp.socket")

# Register an example method
def hello(params):
    name = params.get("name", "World")
    return f"Hello, {name}!"

server.register_method("greet", hello)

# Start server
server.start()

```
</tab>
</tabs>


## 1.7 Service Description

#### **Process Startup Mechanism**
- **Startup Method**:
  ```bash
  nohup ./main > nohup.out 2>&1 &
  ```  
  - ✅ **Background execution** (`&`)
  - ✅ **Terminal signal blocking** (`nohup` ignores `SIGHUP`)
  - ✅ **Complete logging** (`stdout` + `stderr` redirected to `nohup.out`)

- **Process Metadata Storage**:
  - 📌 **PID, PPID, startup time** (prevents PID reuse misjudgment)

---  

#### **Process Termination Mechanism**

**Termination Strategy (Force Priority)**:
  - **Send `SIGKILL` (-9) directly**, immediately release resources
    ```bash
    kill -9 ${PID}  # Unconditional termination, avoids zombie processes
    ```



## 1.8 Socket Specification

When developing external applications using Unix Socket + JSON-RPC, pay special attention to the socket file location. Standard specifications require:

* **Location**: Must be placed in the `/.__gmssh/plugin/organization_name/app_name/tmp/app.sock` directory.

* **Path Description**:

  * **Organization Name**: The name of the application developer or author.
  * **App Name**: The name of the application you are developing, e.g., `redis`, `mysql`, etc.
  * **Naming Convention**: The socket file must be named `app.sock` for uniformity and standardization.

* **Optimization Requirements**: Ensure the socket file is placed in the application's `tmp` directory and stored according to this path. This helps with unified management and improves application maintainability.

#### Example Path:

```
/.__gmssh/plugin/johndoe/redis/tmp/app.sock
```

This ensures the application's socket file can be correctly identified and used by the GA service. For details, see [Request Response](请求响应.md)

> Note: Application names must be unique within the same organization
