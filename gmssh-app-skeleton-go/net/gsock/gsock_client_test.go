package gsock

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/sourcegraph/jsonrpc2"
)

func TestMultiJsonRPCRequest(t *testing.T) {
	socketPath := "/tmp/rpc.sock"
	count := 0

	wg := &sync.WaitGroup{}
	for count < 100 {
		wg.Add(1)
		go func(wg *sync.WaitGroup, index int) {
			defer wg.Done()

			nt := time.Now()
			simpleClient := NewRpcSimpleClient(socketPath)
			var result any

			opts := []jsonrpc2.CallOption{WithSimpleIDClientOpt(index)}
			err := simpleClient.Request(
				context.TODO(), "ping", nil, &result, opts...)
			if err != nil {
				fmt.Printf("[%d] send sockets failed err : %#v  cost: %#v \n", index, err, time.Since(nt).Microseconds())
				return
			}
			fmt.Printf("[%d] request successfully res : %#v cost: %#v \n", index, result, time.Since(nt).Microseconds())
		}(wg, count)
		count++
	}

	wg.Wait()
}

func TestJsonRPCPingClient(t *testing.T) {
	socketPath := "/tmp/rpc.sock"

	simpleClient := NewRpcSimpleClient(socketPath)

	var result any
	err := simpleClient.Request(
		context.TODO(), "ping", nil, &result,
	)
	if err != nil {
		t.Fatalf("send sockets failed err : %#v", err)
		return
	}

	fmt.Printf("[*] request successfully res:%#v \n", result)
}

func TestJsonRPCClient(t *testing.T) {
	socketPath := "/tmp/rpc.sock"

	simpleClient := NewRpcSimpleClient(socketPath)

	var result any
	err := simpleClient.Request(
		context.TODO(), "multiply", map[string]any{"A": 10, "B": 5}, &result,
	)
	if err != nil {
		t.Fatalf("send sockets failed err : %#v", err)
		return
	}

	fmt.Printf("[*] request successfully res:%#v \n", result)
}
