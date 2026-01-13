package simplejrpc

import (
	"github.com/DemonZack/simplejrpc-go/net/gsock"
)

// Server represents a JSON-RPC server with middleware support.
// It wraps the underlying IRpcServer implementation and provides additional functionality.
type Server struct {
	middlewares []gsock.RPCMiddleware // Global middleware that applies to all registered handlers
	service     gsock.IRpcServer      // Underlying JSON-RPC server implementation
}

// NewDefaultServer creates a new Server instance with default configuration.
// It accepts optional configuration functions that can modify the behavior of the JSON-RPC service.
// opts: Optional configuration functions for the JSON-RPC service
// Returns a pointer to the newly created Server instance
func NewDefaultServer(opts ...gsock.JsonRpcSimpleServiceOptionFunc) *Server {
	// Create a new RPC server with a JSON-RPC simple service as the underlying implementation
	service := gsock.NewRpcServer(
		gsock.WithServiceOptFunc(
			gsock.NewJsonRpcSimpleService(opts...),
		),
	)
	return &Server{
		service: service, // Initialize the server with the created RPC service
	}
}

// RegisterHandle registers a new handler function for a specific API endpoint.
// The handler will be wrapped with any provided middlewares, which will be executed
// in addition to the server's global middlewares.
// api: The API endpoint path (e.g., "/api/v1/method")
// hand: The handler function that processes the request and returns a response or error
// middlewares: Optional middlewares that apply only to this specific handler
func (s *Server) RegisterHandle(api string, hand func(req *gsock.Request) (any, error), middlewares ...gsock.RPCMiddleware) {
	s.service.RegisterHandle(api, hand, middlewares...)
}

// Middlewares returns the server's global middlewares.
// Returns: A slice of RPCMiddleware currently registered as global middlewares
func (s *Server) Middlewares() []gsock.RPCMiddleware {
	return s.middlewares
}

// Server returns the underlying IRpcServer implementation.
// This provides access to the core server functionality if needed.
// Returns: The underlying IRpcServer instance
func (s *Server) Server() gsock.IRpcServer {
	return s.service
}

// StartServer starts the JSON-RPC server listening on the specified Unix domain socket path.
// socketPath: The filesystem path where the Unix domain socket will be created
// Returns: An error if the server fails to start, nil otherwise
func (s *Server) StartServer(socketPath string) error {
	return s.service.StartServer(socketPath)
}
