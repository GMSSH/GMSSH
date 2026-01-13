package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/DemonZack/simplejrpc-go/net/gsock"
)

func TestJRpcServerClient(t *testing.T) {
	socketPath := "../zack.sock"

	simpleClient := gsock.NewRpcSimpleClient(socketPath)

	var result any
	err := simpleClient.Request(
		context.TODO(), "hello", nil, &result,
	)
	if err != nil {
		t.Fatalf("send sockets failed err : %#v", err)
		return
	}

	fmt.Printf("[*] request successfully res:%#v \n", result)
}
