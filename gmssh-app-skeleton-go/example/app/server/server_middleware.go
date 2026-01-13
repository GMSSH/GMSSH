package server

import "github.com/DemonZack/simplejrpc-go/net/gsock"

type CustomMiddleware struct{}

func (c *CustomMiddleware) ProcessRequest(req *gsock.Request) {
	println("[*] ProcessRequest. ", req)
}

func (c *CustomMiddleware) ProcessResponse(resp any) (any, error) {
	println("[*] ProcessResponse. ", resp)
	return resp, nil
}
