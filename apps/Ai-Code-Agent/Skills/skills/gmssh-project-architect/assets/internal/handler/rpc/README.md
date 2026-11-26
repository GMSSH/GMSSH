# rpc - RPC 服务注册与处理

本目录是应用的 **RPC 入口**，基于 `simplejrpc-go` 框架实现 JSON-RPC over Unix Socket 通信。

## 核心文件

| 文件 | 说明 |
|------|------|
| `server.go` | 服务启动、路由注册、处理器定义 |

## 关键结构

### Server

`Server` 结构体承载所有 RPC 处理函数，通过 `RegisterHandles` 方法批量注册路由。

### 内置接口

| 方法名 | 说明 |
|--------|------|
| `ping` | 心跳检测，返回 `"pong"` |

## 添加新接口

### 第一步：定义处理函数

```go
func (s *Server) MyMethod(req *gsock.Request) (any, error) {
    var args dto.MyMethodReq
    if err := rpcutil.ParseParams(req, &args); err != nil {
        return nil, err
    }
    // 调用 service 层
    return result, nil
}
```

### 第二步：注册路由

在 `RegisterHandles` 方法中添加：

```go
ds.RegisterHandle("myMethod", s.MyMethod)
```

## 启动流程

`Start()` 方法依次执行：

1. 创建 Socket 目录
2. 初始化 RPC Server
3. 设置配置路径和国际化路径
4. 初始化容器 (日志等核心组件)
5. 注册所有路由
6. 监听 Unix Socket 启动服务
