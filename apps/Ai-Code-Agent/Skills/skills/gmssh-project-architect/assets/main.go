package main

import (
	"fmt"
	"github.com/DemonZack/simplejrpc-go/core"
	"template/internal/handler/rpc"
)

func main() {
	sock := "../../tmp/app.sock"

	// 创建 RPC Server 并注册服务
	server := rpc.NewServer()

	// 启动服务
	if err := server.Start(sock); err != nil {
		core.Container.Log().Error(fmt.Sprintf("启动RPC服务器失败: %v", err))
		return
	}

	core.Container.Log().Info(fmt.Sprintf("Site Hub 服务启动成功, socket: %s", sock))
}
