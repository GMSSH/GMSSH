# JSON-RPC 接口规范

> 本规范定义了 GMSSH 外置应用后端与前端之间的通信协议标准，传输层支持 HTTP 和 Unix Socket 两种方式。

---

## 1. 通信架构

### 1.1 调用链路

```
┌──────────────┐         ┌──────────────┐    Unix Socket    ┌──────────────┐
│  Iframe 前端  │──HTTP──►│  GMSSH 宿主   │ ◄──────────────► │  外置应用后端  │
│  (SDK 调用)   │         │  (API 网关)   │   JSON-RPC       │  (Go 进程)    │
└──────────────┘         └──────────────┘                   └──────────────┘
```

### 1.2 外置应用访问接口

**接口地址**:

```
/api/call/{组织名}/{应用名}/{接口名}
```

**示例**: `/api/call/test/myapp/hello`

| 路径参数 | 说明 | 示例 |
|----------|------|------|
| `{组织名}` | 应用所属组织标识 | `test` |
| `{应用名}` | 应用名称 | `myapp` |
| `{接口名}` | 后端注册的 RPC 方法名 | `hello` |

### 1.3 传输方式

| transport | 说明 | 适用场景 |
|-----------|------|----------|
| `socket` | Unix Socket 通信 | 后端为常驻 Go 进程 (推荐) |
| `http` | HTTP 通信 | 后端为 HTTP 服务 |

---

## 2. 健康检查 (ping 接口)

> [!IMPORTANT]
> `ping` 接口是 GMSSH 平台的**强制要求**。未实现此接口的应用无法在平台上正常调试和运行。

### 2.1 平台检查机制

GMSSH 宿主会对所有外置应用后端进行持续的健康状态探测：

| 参数 | 值 | 说明 |
|------|------|------|
| 检查频率 | 每 **5 秒** 一次 | 平台自动发起，无需应用触发 |
| 超时时间 | **3 ~ 5 秒** | 超时视为服务不可用 |
| 期望响应 | HTTP 200 / RPC 正常返回 | 任何错误响应均判定为异常 |

### 2.2 请求与响应

**请求**: `/api/call/{组织名}/{应用名}/ping`

```json
{
  "version": "1.0.0",
  "transport": "socket",
  "params": {}
}
```

**响应**:

```json
{
  "code": 200000,
  "data": {
    "code": 200,
    "data": "pong",
    "msg": "操作成功"
  },
  "meta": {
    "close": 0,
    "endpoint": "/api/call/*"
  },
  "msg": "Successful operation"
}
```

### 2.3 脚手架默认实现

模板项目已在 `internal/handler/rpc/server.go` 中内置了 `ping` 接口：

```go
// Ping 心跳检测 — GMSSH 平台健康检查接口 (勿删除)
func (s *Server) Ping(req *gsock.Request) (any, error) {
    return "pong", nil
}
```

路由注册 (`RegisterHandles`):

```go
ds.RegisterHandle("ping", s.Ping)  // ⚠️ 必须保留，平台依赖此接口
```

> [!CAUTION]
> **禁止删除或重命名 `ping` 接口**。方法名必须为 `"ping"`，返回值必须为 `"pong"`。修改任何一项都会导致平台判定服务异常，应用将无法正常调试。

### 2.4 排查健康检查失败

| 现象 | 可能原因 | 解决方案 |
|------|---------|---------| 
| 应用启动后立即显示异常 | Socket 文件路径错误 | 检查 `main.go` 中的 `sock` 路径与平台配置一致 |
| 间歇性健康检查失败 | Handler 执行时间 > 5s | `ping` 中不要添加任何耗时逻辑 |
| 启动正常但后续异常 | 进程崩溃后 Socket 残留 | 启动前清理旧 `.sock` 文件 |

---

## 3. 请求格式

### 3.1 标准请求结构

```json
{
  "version": "1.0.0",
  "transport": "socket",
  "params": {
    "lang": "zh-CN"
  }
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `version` | string | ✅ | 外置应用版本号，使用 GMSSH SDK 会自动填入 |
| `transport` | string | ✅ | 通信类型：`"http"` 或 `"socket"` |
| `params` | object | ❌ | 传递给外置应用的参数 |

### 3.2 保留参数

每个请求的 `params` 中可包含以下框架级保留字段：

| 字段 | 类型 | 说明 |
|------|------|------|
| `lang` | string | 客户端语言标识，由 `rpcutil.ParseParams` 自动提取设置 i18n |

### 3.3 真实示例 — Hello 请求

```json
{
  "version": "1.0.0",
  "transport": "socket",
  "params": {
    "lang": "zh-CN"
  }
}
```

> [!NOTE]
> `params` 中的内容完全由业务自定义，脚手架不做额外限制。上例仅传递了 `lang` 语言标识。

---

## 4. 响应格式

### 4.1 统一响应结构

所有 API 接口遵循统一的响应格式：

```json
{
  "code": 200000,
  "data": "<响应数据>",
  "meta": {},
  "msg": "Successful operation"
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| `code` | number | 状态码，成功为 `200000`，失败为非 `200000` |
| `data` | any | 响应数据，类型由具体接口决定 |
| `meta` | object | 元数据信息 |
| `msg` | string | 响应消息（英文） |

> [!IMPORTANT]
> 判断请求是否成功的**唯一依据**是 `code === 200000`，不要依赖 `msg` 字段判断。

### 4.2 真实示例 — Hello 成功响应

后端 Handler：

```go
func (s *Server) Hello(req *gsock.Request) (any, error) {
    return "Hello", nil
}
```

> [!IMPORTANT]
> **开发者只需关心 Handler 的返回值**。脚手架会自动将返回值包装为完整的 JSON-RPC 响应结构。
> 上面的 Handler 仅返回了字符串 `"Hello"`，脚手架自动将其适配为下方的完整响应：

前端收到的完整响应：

```json
{
  "code": 200000,
  "data": {
    "code": 200,
    "data": "Hello",
    "msg": "操作成功"
  },
  "meta": {
    "close": 0,
    "endpoint": "/api/call/*"
  },
  "msg": "Successful operation"
}
```

**映射关系**：

| Handler 返回值 | 响应中的位置 | 说明 |
|---------------|------------|------|
| 第一个返回值 (`"Hello"`) | `data.data` | Handler 的返回数据 |
| 第二个返回值 (`nil`) | — | `nil` 表示成功，`code` 自动设为 `200000` |
| `data.code` / `data.msg` | 自动填充 | 由脚手架根据 i18n 和状态自动设置 |
| 外层 `code` / `meta` / `msg` | 自动填充 | 由 GMSSH 宿主自动包装 |

### 4.3 真实示例 — Hello_ERROR 错误响应

当 Handler 返回 `error` 时，外层 `code` 为非 `200000`，前端可据此判断请求失败：

后端 Handler：

```go
func (s *Server) Hello_ERROR(req *gsock.Request) (any, error) {
    return nil, errors.New("error msg")
}
```

前端收到的错误响应：

```json
{
  "code": 500000,
  "data": null,
  "meta": {
    "close": 0,
    "endpoint": "/api/call/*"
  },
  "msg": "error msg"
}
```

> [!TIP]
> 开发者只需 `return nil, errors.New("错误信息")` 即可返回错误，脚手架会自动将 `error.Error()` 映射为响应中的 `msg` 字段。

### 4.4 Handler 返回值总结

| 场景 | Handler 返回 | 外层 `code` | 前端判断 |
|------|-------------|------------|----------|
| 成功 | `return data, nil` | `200000` | `code === 200000` |
| 失败 | `return nil, errors.New("...")` | 非 `200000` | `code !== 200000` |

---

## 5. 标准错误码

### 5.1 协议保留错误码

| 错误码 | 含义 | 说明 |
|--------|------|------|
| `-32700` | Parse error | JSON 解析失败 |
| `-32600` | Invalid Request | 请求格式不符合规范 |
| `-32601` | Method not found | 方法名未注册 |
| `-32602` | Invalid params | 参数无效 |
| `-32603` | Internal error | 服务端内部错误 |

### 5.2 业务错误码

业务错误建议通过自定义响应结构区分，而非使用 JSON-RPC error：

```go
// 方式一：返回 error (框架自动映射为错误响应，code != 200000)
return nil, errors.New("user not found")

// 方式二：返回业务状态码 (推荐复杂业务使用，code 仍为 200000，业务层自行判断)
return map[string]any{
    "code":    40004,
    "message": "user not found",
    "data":    nil,
}, nil
```

---

## 6. Handler 开发规范

### 6.1 函数签名

所有 RPC Handler 必须遵循统一签名：

```go
func (s *Server) MethodName(req *gsock.Request) (any, error)
```

### 6.2 标准处理流程

```go
func (s *Server) CreateItem(req *gsock.Request) (any, error) {
    // 1️⃣ 解析参数 (自动处理 i18n)
    var args dto.CreateItemReq
    if err := rpcutil.ParseParams(req, &args); err != nil {
        return nil, err
    }

    // 2️⃣ 参数校验
    if args.Name == "" {
        return nil, fmt.Errorf(gi18n.Instance().T("NameRequired"))
    }

    // 3️⃣ 调用 Service 层
    result, err := service.NewItemService().Create(args)
    if err != nil {
        return nil, err
    }

    // 4️⃣ 返回结果
    return result, nil
}
```

### 6.3 路由注册

在 `RegisterHandles` 中集中注册，方法名使用 **camelCase**：

```go
func (s *Server) RegisterHandles(ds interface {
    RegisterHandle(name string, handler func(*gsock.Request) (any, error), middlewares ...gsock.RPCMiddleware)
}) {
    // 基础接口 (平台必须)
    ds.RegisterHandle("ping", s.Ping)

    // 业务接口 — 按模块分组注释
    ds.RegisterHandle("getItemList", s.GetItemList)
    ds.RegisterHandle("createItem", s.CreateItem)
    ds.RegisterHandle("deleteItem", s.DeleteItem)
}
```

---

## 7. 中间件

`simplejrpc-go` 支持为路由绑定中间件，通过实现 `RPCMiddleware` 接口：

```go
type AuthMiddleware struct{}

// ProcessRequest 请求前置处理 (鉴权、日志等)
func (m *AuthMiddleware) ProcessRequest(req *gsock.Request) {
    rpcutil.SetLanguage(req)
    core.Container.Log().Info(fmt.Sprintf("收到请求: %s", req.RawRequest().Method))
}

// ProcessResponse 响应后置处理 (日志、数据包装等)
func (m *AuthMiddleware) ProcessResponse(resp any) (any, error) {
    return resp, nil
}

// 注册时绑定中间件
auth := &AuthMiddleware{}
ds.RegisterHandle("secureMethod", s.SecureMethod, auth)
```

---

## 8. 前端调用示例

### 8.1 通过 SDK 调用

```javascript
// GMSSH 前端 SDK 自动处理 Socket 通信
const result = await window.$gm.request('/rpc', {
    method: 'POST',
    data: {
        method: 'getItemList',
        params: {
            lang: currentLang,
            page: 1,
            size: 20
        }
    }
});
```

### 8.2 方法命名对照

| 前端调用 method | 后端 Handler | 说明 |
|----------------|-------------|------|
| `"ping"` | `Server.Ping` | 健康检查 (平台必须) |
| `"hello"` | `Server.Hello` | 示例接口 |
| `"getItemList"` | `Server.GetItemList` | 获取列表 |
| `"createItem"` | `Server.CreateItem` | 创建条目 |

---

## 9. 接口文档模板

新增接口时，建议在代码注释中包含以下信息：

```go
// GetItemList 获取条目列表
//
// Method: getItemList
//
// Params:
//   - lang   string  (可选) 语言标识
//   - page   int     (必填) 页码，从 1 开始
//   - size   int     (必填) 每页数量，最大 100
//   - search string  (可选) 搜索关键词
//
// Result:
//   - total  int           总数
//   - list   []ItemResp    条目数组
//
// Errors:
//   - InvalidParams  参数无效
func (s *Server) GetItemList(req *gsock.Request) (any, error) {
    // ...
}
```
