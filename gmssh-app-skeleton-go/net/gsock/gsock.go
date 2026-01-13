package gsock

import (
	"context"
	"net"

	"github.com/sourcegraph/jsonrpc2"
)

// RPCMiddleware defines the interface for request/response processing middleware
// Middlewares can inspect/modify requests and responses in the processing pipeline
type RPCMiddleware interface {
	// ProcessRequest is called before handler execution
	ProcessRequest(req *Request)

	// ProcessResponse is called after handler execution
	// Can modify or replace the response
	ProcessResponse(resp any) (any, error)
}

// IRpcHandler provides method registration capabilities for RPC services
type IRpcHandler interface {
	// RegisterHandle binds a handler function to an API endpoint
	// Middlewares are executed in registration order
	RegisterHandle(api string, hand func(req *Request) (any, error), middlewares ...RPCMiddleware)
}

// IRpcServer combines handler registration with server lifecycle management
type IRpcServer interface {
	IRpcHandler

	// StartServer begins listening on the specified Unix domain socket
	// Returns error if server fails to start
	StartServer(socketPath string) error
}

// RpcServiceDispatcher maps API endpoints to their handler functions
type RpcServiceDispatcher map[string]func(req *Request) (any, error)

// IRpcService defines the core RPC service interface combining:
// - Handler registration
// - Connection management
// - Request processing
type IRpcService interface {
	IRpcHandler

	// NewConn creates a managed JSON-RPC 2.0 connection
	NewConn(ctx context.Context, conn net.Conn) *jsonrpc2.Conn

	// Handle processes incoming JSON-RPC requests
	Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (any, error)
}

// IRpcServiceHandle provides a simplified handler-only interface
type IRpcServiceHandle interface {
	IRpcHandler
	Handle(req *Request) (any, error)
}

// IRpcClient defines the client interface for making RPC calls
type IRpcClient interface {
	// Request executes a JSON-RPC 2.0 method call
	// Parameters:
	//   - ctx: Context for cancellation/timeout
	//   - method: RPC method name
	//   - params: Input parameters
	//   - result: Pointer to struct for response decoding
	//   - opts: Additional call options
	Request(ctx context.Context, method string, params, result any, opts ...jsonrpc2.CallOption) error
}

// ClientAdapter creates client connections for different protocols
type ClientAdapter interface {
	// NewConn establishes a new protocol-specific client connection
	NewConn(ctx context.Context, conn net.Conn) IRpcClient
}
