package server

import "github.com/DemonZack/simplejrpc-go/net/gsock"

type CustomHandler struct{}

func (c *CustomHandler) Hello(req *gsock.Request) (any, error) {
	return "Hello World", nil
}

func (c *CustomHandler) ProcessRequest(req *gsock.Request) {
	println("[*] ProcessRequest. ", req)
}

func (c *CustomHandler) ProcessResponse(resp any) (any, error) {
	println("[*] ProcessResponse. ", resp)
	return resp, nil
}
