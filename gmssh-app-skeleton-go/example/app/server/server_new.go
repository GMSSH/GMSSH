package server

import (
	rpc "github.com/DemonZack/simplejrpc-go"
	"github.com/DemonZack/simplejrpc-go/net/gsock"
)

type AppServer struct{}

func NewAppServer() *AppServer {
	return &AppServer{}
}

func (s *AppServer) Run() {
	mockSockPath := "zack.sock"

	ds := rpc.NewDefaultServer(
		gsock.WithJsonRpcSimpleServiceHandler(gsock.NewJsonRpcSimpleServiceHandler()),
		gsock.WithJsonRpcSimpleServiceMiddlewares(),
	)

	hand := &CustomHandler{}
	ds.RegisterHandle("hello", hand.Hello, []gsock.RPCMiddleware{hand}...)

	err := ds.StartServer(mockSockPath)
	if err != nil {
		panic(err)
	}
}
