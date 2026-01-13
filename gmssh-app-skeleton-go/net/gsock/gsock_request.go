package gsock

import (
	"context"

	"github.com/sourcegraph/jsonrpc2"
)

// RequestOptFunc defines functions for configuring Request objects
type RequestOptFunc func(*Request)

// Request wraps a JSON-RPC 2.0 request with additional context and functionality
type Request struct {
	ctx context.Context   // Context for cancellation/timeout
	req *jsonrpc2.Request // Underlying JSON-RPC request
}

// RawRequest returns the underlying JSON-RPC 2.0 request object
// Useful for accessing low-level request details when needed
func (r *Request) RawRequest() *jsonrpc2.Request {
	return r.req
}

// Method returns the RPC method name being called
// This is a convenience method that delegates to the underlying request
func (r *Request) Method() string {
	return r.req.Method
}

// Context returns the request's context
// The context carries deadlines, cancellation signals, and other request-scoped values
func (r *Request) Context() context.Context {
	return r.ctx
}

// WithRequestCtxOption creates a RequestOptFunc that sets the request context
// This is typically used to propagate cancellation signals and deadlines
func WithRequestCtxOption(ctx context.Context) RequestOptFunc {
	return func(r *Request) {
		r.ctx = ctx
	}
}

// WithRequestReqOption creates a RequestOptFunc that sets the underlying JSON-RPC request
// This is used when wrapping an existing JSON-RPC request
func WithRequestReqOption(req *jsonrpc2.Request) RequestOptFunc {
	return func(r *Request) {
		r.req = req
	}
}

// MakeRequest constructs a new Request instance with the provided options
// This follows the functional options pattern for flexible request creation
//
// Example:
//
//	req := MakeRequest(
//	    WithRequestCtxOption(ctx),
//	    WithRequestReqOption(jsonReq),
//	)
func MakeRequest(opts ...RequestOptFunc) *Request {
	r := &Request{
		ctx: context.Background(), // Default context
	}

	// Apply all configuration options
	for _, opt := range opts {
		opt(r)
	}
	return r
}
