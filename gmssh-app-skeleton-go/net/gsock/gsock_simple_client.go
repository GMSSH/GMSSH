package gsock

import (
	"context"
	"net"

	"github.com/sourcegraph/jsonrpc2"
)

// JsonRpcSimpleClientHandler implements IRpcClient for making JSON-RPC 2.0 requests
type JsonRpcSimpleClientHandler struct {
	conn *jsonrpc2.Conn // The underlying JSON-RPC connection
}

// WithSimpleIDClientOpt creates a jsonrpc2.CallOption to set a numeric request ID
// This allows explicit control over request IDs rather than using auto-increment
func WithSimpleIDClientOpt(id int) jsonrpc2.CallOption {
	return jsonrpc2.PickID(jsonrpc2.ID{Num: uint64(id)})
}

// NewJsonRpcSimpleClientHandler creates a new client handler wrapping a JSON-RPC connection
func NewJsonRpcSimpleClientHandler(conn *jsonrpc2.Conn) *JsonRpcSimpleClientHandler {
	return &JsonRpcSimpleClientHandler{
		conn: conn,
	}
}

// Request executes a JSON-RPC 2.0 method call
// Implements the IRpcClient interface
func (c *JsonRpcSimpleClientHandler) Request(
	ctx context.Context,
	method string,
	params, result any,
	opts ...jsonrpc2.CallOption,
) error {
	return c.conn.Call(ctx, method, params, result, opts...)
}

// JsonRpcSimpleClient implements ClientAdapter for creating JSON-RPC 2.0 clients
type JsonRpcSimpleClient struct{}

// NewConn establishes a new JSON-RPC 2.0 client connection
// Implements the ClientAdapter interface
func (r *JsonRpcSimpleClient) NewConn(ctx context.Context, conn net.Conn) IRpcClient {
	// Create buffered connection with VSCode-style message codec
	jsonConn := jsonrpc2.NewConn(
		ctx,
		jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{}),
		jsonrpc2.HandlerWithError(r.Handle), // Empty handler for client connections
	)
	return NewJsonRpcSimpleClientHandler(jsonConn)
}

// Handle is a placeholder for client-side request handling
// For pure clients, this typically won't be used as clients make requests rather than handle them
func (r *JsonRpcSimpleClient) Handle(
	ctx context.Context,
	conn *jsonrpc2.Conn,
	req *jsonrpc2.Request,
) (any, error) {
	// Client connections don't typically handle incoming requests
	// This would be used if implementing bidirectional RPC
	return nil, nil
}

// Close provides clean connection shutdown
// func (c *JsonRpcSimpleClientHandler) Close() error {
//     return c.conn.Close()
// }

// Notification sends a JSON-RPC notification (request without response)
// func (c *JsonRpcSimpleClientHandler) Notification(
//     ctx context.Context,
//     method string,
//     params any,
// ) error {
//     return c.conn.Notify(ctx, method, params)
// }

// BatchRequest sends multiple requests in a single batch
// func (c *JsonRpcSimpleClientHandler) BatchRequest(
//     ctx context.Context,
//     calls []jsonrpc2.BatchCall,
// ) error {
//     return c.conn.BatchCall(ctx, calls)
// }
