package gsock

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// RpcServerOptFunc defines functions for configuring an RPC server
type RpcServerOptFunc func(*rpcServer)

// rpcServer implements a JSON-RPC 2.0 server over Unix domain sockets
type rpcServer struct {
	service IRpcService // The underlying RPC service implementation
}

// WithServiceOptFunc creates a configuration option to set the RPC service
func WithServiceOptFunc(service IRpcService) RpcServerOptFunc {
	return func(rs *rpcServer) {
		rs.service = service
	}
}

// NewRpcServer creates a new RPC server instance with the given options
func NewRpcServer(opts ...RpcServerOptFunc) *rpcServer {
	server := &rpcServer{}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

// RegisterHandle adds a new handler for an RPC method
func (r *rpcServer) RegisterHandle(api string, hand func(req *Request) (any, error), middlewares ...RPCMiddleware) {
	r.service.RegisterHandle(api, hand, middlewares...)
}

// StartServer begins listening for RPC connections on a Unix domain socket
// It handles graceful shutdown on interrupt signals and cleans up the socket file
func (s *rpcServer) StartServer(socketPath string) error {
	// Clean up any existing socket file
	if err := os.Remove(socketPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove existing socket: %w", err)
	}

	// Create Unix domain socket listener
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return fmt.Errorf("failed to listen on Unix socket: %w", err)
	}

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	// Goroutine for handling shutdown signals
	go func() {
		<-sigChan
		log.Println("[*] Received shutdown signal, cleaning up...")

		// Clean up resources
		os.Remove(socketPath)
		listener.Close()
		cancel()
	}()

	log.Printf("JSON-RPC server listening on Unix socket: %s\n", socketPath)

	// Main connection acceptance loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			// Check if error is due to shutdown
			select {
			case <-ctx.Done():
				return nil // Graceful shutdown
			default:
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue // Temporary error, keep listening
				}
				return fmt.Errorf("connection accept error: %w", err)
			}
		}

		// Handle each connection in a separate goroutine
		go func(conn net.Conn) {
			// defer conn.Close()
			s.service.NewConn(ctx, conn)
		}(conn)
	}
}

// SetSocketPermissions sets file permissions for the Unix socket
// func (s *rpcServer) SetSocketPermissions(perm os.FileMode) RpcServerOptFunc {
//     return func(rs *rpcServer) {
//         rs.socketPerm = perm
//     }
// }

// WithLogger allows customizing the server logger
// func WithLogger(logger *log.Logger) RpcServerOptFunc {
//     return func(rs *rpcServer) {
//         rs.logger = logger
//     }
// }

// Connection limiting middleware could be added
// type connectionLimiter struct {
//     sem chan struct{}
// }
//
// func NewConnectionLimiter(max int) *connectionLimiter {
//     return &connectionLimiter{
//         sem: make(chan struct{}, max),
//     }
// }
