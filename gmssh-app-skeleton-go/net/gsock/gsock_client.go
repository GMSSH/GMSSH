package gsock

import (
	"context"
	"net"

	"github.com/sourcegraph/jsonrpc2"
)

// RPCClient provides JSON-RPC 2.0 client functionality over Unix domain sockets.
// It manages connection lifecycle and request/response handling through a configurable adapter.
//
// Fields:
//   - sockPath:  Filesystem path to the Unix domain socket (e.g., "/tmp/rpc.sock")
//   - adapter:   Protocol adapter implementing the ClientAdapter interface (defaults to JSON-RPC)
//   - idCounter: Atomic counter for generating unique request IDs
//   - keepLive:  Flag controlling whether connections should be kept alive after requests
type rpcClient struct {
	sockPath  string        // Path to the Unix domain socket
	adapter   ClientAdapter // Protocol adapter (defaults to JSON-RPC)
	idCounter int64         // Atomic counter for generating request IDs
	keepLive  bool          // Connection persistence flag
}

// NewRpcKeepLivClient creates a new RPC client with connection persistence disabled.
// Connections will remain open after requests (caller must manage cleanup).
//
// Parameters:
//   - socketPath: Absolute filesystem path to the Unix domain socket
//
// Returns:
//   - *rpcClient: Initialized client instance ready for RPC calls
//
// Note:
//
//	Uses JsonRpcSimpleClient as the default adapter
func NewRpcSimpleClient(socketPath string) *rpcClient {
	return &rpcClient{
		sockPath:  socketPath,
		adapter:   &JsonRpcSimpleClient{},
		idCounter: 0,
	}
}

// NewRpcSimpleClient creates a new RPC client with connection persistence enabled.
// The client will automatically close connections after each request.
//
// Parameters:
//   - socketPath: Absolute filesystem path to the Unix domain socket
//
// Returns:
//   - *rpcClient: Initialized client instance with keepalive disabled
//
// Note:
//
//	Uses JsonRpcSimpleClient as the default adapter
func NewRpcKeepLiveClient(socketPath string) *rpcClient {
	return &rpcClient{
		sockPath:  socketPath,
		adapter:   &JsonRpcSimpleClient{},
		idCounter: 0,
		keepLive:  true,
	}
}

// Request executes a JSON-RPC 2.0 method call and handles response decoding.
// Automatically manages connection establishment, request ID generation, and error handling.
//
// Parameters:
//   - ctx:       Context for cancellation and timeout control
//   - method:    RPC method name to invoke
//   - params:    Input parameters (will be JSON-serialized)
//   - result:    Pointer to structure for response deserialization
//   - opts:      Optional JSON-RPC 2.0 call configurations
//
// Returns:
//   - error:     nil on success, or error describing failure:
//   - Connection errors
//   - Protocol errors
//   - Deserialization errors
//
// Example:
//
//	var response ResponseStruct
//	err := client.Request(
//	    context.Background(),
//	    "Service.Method",
//	    RequestParams{Field: "value"},
//	    &response,
//	)
func (c *rpcClient) Request(ctx context.Context, method string, params, result any, opts ...jsonrpc2.CallOption) error {
	conn, err := net.Dial("unix", c.sockPath)
	if err != nil {
		return err
	}
	if !c.keepLive {
		defer conn.Close()
	}

	// Generate monotonic request ID
	c.idCounter++
	idOpt := jsonrpc2.PickID(jsonrpc2.ID{Num: uint64(c.idCounter)})
	opts = append(opts, idOpt)

	client := c.adapter.NewConn(ctx, conn)
	return client.Request(ctx, method, params, result, opts...)
}
