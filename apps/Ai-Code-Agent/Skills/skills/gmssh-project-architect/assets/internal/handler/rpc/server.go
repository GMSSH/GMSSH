package rpc

import (
	"github.com/DemonZack/simplejrpc-go"
	"github.com/DemonZack/simplejrpc-go/core"
	"github.com/DemonZack/simplejrpc-go/core/config"
	"github.com/DemonZack/simplejrpc-go/core/gi18n"
	"github.com/DemonZack/simplejrpc-go/net/gsock"
	"github.com/DemonZack/simplejrpc-go/os/gpath"
	"os"
)

// Server RPC服务器
type Server struct {
}

// NewServer 创建RPC服务器实例
func NewServer() *Server {

	return &Server{}
}

// RegisterHandles 注册所有RPC处理函数
func (s *Server) RegisterHandles(ds interface {
	RegisterHandle(name string, handler func(*gsock.Request) (any, error), middlewares ...gsock.RPCMiddleware)
},
) {
	// 注册基础接口
	ds.RegisterHandle("ping", s.Ping)
}

// Start 启动RPC服务器
func (s *Server) Start(sockPath string) error {
	_ = os.MkdirAll(sockPath, 0o755)
	ds := simplejrpc.NewDefaultServer(
		gsock.WithJsonRpcSimpleServiceHandler(gsock.NewJsonRpcSimpleServiceHandler()),
	)
	gpath.GmCfgPath = "./"
	gi18n.Instance().SetPath("./i18n")
	core.InitContainer(config.WithConfigEnvFormatterOptionFunc("test"))

	s.RegisterHandles(ds)

	return ds.StartServer(sockPath)
}

// Ping 心跳检测
func (s *Server) Ping(req *gsock.Request) (any, error) {
	// core.Container.Log().Info("收到Ping请求")
	return "pong", nil
}
