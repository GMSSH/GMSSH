package gsock

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/sourcegraph/jsonrpc2"

	"github.com/DemonZack/simplejrpc-go/container/garray"
)

// JsonRpcSimpleServiceHandler implements IRpcServiceHandle for processing JSON-RPC 2.0 requests.
// It maintains a registry of method handlers and a middleware chain for request processing.
type JsonRpcSimpleServiceHandler struct {
	handlers    RpcServiceDispatcher // Map of API method names to their handler functions
	middlewares []RPCMiddleware      // Chain of middleware processors for request/response handling
}

// NewJsonRpcSimpleServiceHandler creates and initializes a new JsonRpcSimpleServiceHandler instance.
// Returns: Pointer to the newly created handler instance
func NewJsonRpcSimpleServiceHandler() *JsonRpcSimpleServiceHandler {
	return &JsonRpcSimpleServiceHandler{
		handlers: make(RpcServiceDispatcher),
	}
}

// multiply is an example handler method that demonstrates parameter parsing and processing.
// It multiplies two integers provided in the request parameters.
// req: The incoming request object
// Returns: Product of the two integers or error if parameters are invalid
func (h *JsonRpcSimpleServiceHandler) multiply(req *Request) (any, error) {
	var args struct {
		A, B int
	}
	if err := json.Unmarshal(*req.RawRequest().Params, &args); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	return args.A * args.B, nil
}

// registerServer is a placeholder method for server registration functionality.
// Currently unimplemented.
func (h *JsonRpcSimpleServiceHandler) registerServer(req *Request) (any, error) {
	// TODO: Implement server registration logic
	return nil, nil
}

// RegisterHandle adds a new method handler to the service's handler registry.
// api: The method name to register
// hand: The handler function to execute for this method
// middlewares: Optional middleware specific to this handler
func (h *JsonRpcSimpleServiceHandler) RegisterHandle(
	api string,
	hand func(req *Request) (any, error),
	middlewares ...RPCMiddleware,
) {
	if len(middlewares) > 0 {
		h.middlewares = append(h.middlewares, middlewares...)
	}
	h.handlers[api] = hand
}

// Ping implements a simple health check endpoint.
// Returns: Constant "pong" response
func (h *JsonRpcSimpleServiceHandler) Ping(req *Request) (any, error) {
	return "pong", nil
}

// ProcessRequest executes the request processing middleware chain.
// req: The request object to process
func (h *JsonRpcSimpleServiceHandler) ProcessRequest(req *Request) {
	for _, middleware := range h.middlewares {
		middleware.ProcessRequest(req)
	}
}

// ProcessResponse executes the response processing middleware chain.
// rep: The response object to process
// Returns: Processed response or error if middleware fails
func (h *JsonRpcSimpleServiceHandler) ProcessResponse(rep any) (any, error) {
	out := rep
	array := garray.NewArray[RPCMiddleware](h.middlewares)
	for _, middleware := range array.Reverse() {
		var err error
		out, err = middleware.ProcessResponse(out)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// Handle processes an incoming JSON-RPC request by finding and executing the appropriate handler.
// req: The incoming request object
// Returns: Response object or error if processing fails
func (h *JsonRpcSimpleServiceHandler) Handle(req *Request) (any, error) {
	response := NewResponse()
	method := req.Method()
	response.SetEndpoint(method)

	handler, ok := h.handlers[method]
	if !ok {
		response.Code = http.StatusNotFound
		response.Message = http.StatusText(http.StatusNotFound)
		return response, nil
	}

	body, err := handler(req)
	response.Data = body
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
	}
	return response, nil
}

// JsonRpcSimpleServiceOptionFunc defines the signature for service configuration functions.
type JsonRpcSimpleServiceOptionFunc func(*JsonRpcSimpleService)

// JsonRpcSimpleService implements IRpcService for handling JSON-RPC 2.0 protocol.
type JsonRpcSimpleService struct {
	handler     IRpcServiceHandle // Core request handler implementation
	middlewares []RPCMiddleware   // Service-level middleware chain
}

// NewDefaultJsonRpcSimpleService creates a service instance with default configuration.
// hand: The handler implementation to use
// Returns: New service instance
func NewDefaultJsonRpcSimpleService(handler IRpcServiceHandle) *JsonRpcSimpleService {
	return &JsonRpcSimpleService{
		handler:     handler,
		middlewares: make([]RPCMiddleware, 0),
	}
}

// WithJsonRpcSimpleServiceHandler creates a configuration function to set the service handler.
// hand: The handler implementation to configure
// Returns: Configuration function
func WithJsonRpcSimpleServiceHandler(hand IRpcServiceHandle) JsonRpcSimpleServiceOptionFunc {
	return func(s *JsonRpcSimpleService) {
		s.handler = hand
	}
}

// WithJsonRpcSimpleServiceMiddlewares creates a configuration function to set service middleware.
// middlewares: Middleware chain to configure
// Returns: Configuration function
func WithJsonRpcSimpleServiceMiddlewares(middlewares ...RPCMiddleware) JsonRpcSimpleServiceOptionFunc {
	return func(s *JsonRpcSimpleService) {
		s.middlewares = append(s.middlewares, middlewares...)
	}
}

// NewJsonRpcSimpleService creates a new service instance with custom configuration.
// opts: Optional configuration functions
// Returns: Configured service instance
func NewJsonRpcSimpleService(opts ...JsonRpcSimpleServiceOptionFunc) *JsonRpcSimpleService {
	rpc := &JsonRpcSimpleService{}
	for _, opt := range opts {
		opt(rpc)
	}
	return rpc
}

// NewConn creates a new JSON-RPC 2.0 connection with context and codec support.
// ctx: Context for the connection
// conn: Underlying network connection
// Returns: New JSON-RPC 2.0 connection
func (r *JsonRpcSimpleService) NewConn(ctx context.Context, conn net.Conn) *jsonrpc2.Conn {
	return jsonrpc2.NewConn(
		ctx,
		jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{}),
		jsonrpc2.HandlerWithError(r.Handle),
	)
}

// ProcessRequest executes the service-level request middleware chain.
// req: Request object to process
func (r *JsonRpcSimpleService) ProcessRequest(req *Request) {
	for _, middleware := range r.middlewares {
		middleware.ProcessRequest(req)
	}
}

// RegisterHandle delegates handler registration to the underlying handler implementation.
// api: Method name to register
// hand: Handler function to execute
// middlewares: Optional middleware for this handler
func (r *JsonRpcSimpleService) RegisterHandle(
	api string,
	hand func(req *Request) (any, error),
	middlewares ...RPCMiddleware,
) {
	r.handler.RegisterHandle(api, hand, middlewares...)
}

// ProcessResponse executes the service-level response middleware chain in reverse order.
// rep: Response object to process
// Returns: Processed response or error if middleware fails
func (r *JsonRpcSimpleService) ProcessResponse(rep any) (any, error) {
	out := rep
	array := garray.NewArray[RPCMiddleware](r.middlewares)
	for _, middleware := range array.Reverse() {
		var err error
		out, err = middleware.ProcessResponse(out)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// Handle processes an incoming JSON-RPC request through the complete middleware and handler chain.
// ctx: Request context
// conn: JSON-RPC connection
// req: Incoming request
// Returns: Response object or error if processing fails
func (r *JsonRpcSimpleService) Handle(
	ctx context.Context,
	conn *jsonrpc2.Conn,
	req *jsonrpc2.Request,
) (any, error) {
	request := MakeRequest(
		WithRequestCtxOption(ctx),
		WithRequestReqOption(req),
	)
	r.ProcessRequest(request)

	response, err := r.handler.Handle(request)
	if err != nil {
		return response, err
	}

	return r.ProcessResponse(response)
}
