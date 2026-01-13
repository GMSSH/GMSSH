package glog

import (
	"fmt"
	"runtime"

	"go.uber.org/zap"
)

// GLogger is an enhanced logger that automatically includes stack traces
// for error-level logs. It wraps a zap.Logger while maintaining all its
// original functionality and adding stack trace capabilities.
type GLogger struct {
	*zap.Logger // Embedded zap.Logger for core logging functionality
}

// NewGLogger creates a new wrapped logger instance that will automatically
// include stack traces in error logs. This provides better debugging
// information when errors occur.
//
// Parameters:
//   - logger: The base zap.Logger instance to wrap
//
// Returns:
//   - *GLogger: The enhanced logger instance
func NewGLogger(logger *zap.Logger) *GLogger {
	return &GLogger{
		Logger: logger,
	}
}

// Error logs an error message with the current goroutine's stack trace.
// The stack trace helps identify where the error originated.
// Uses a 64KB buffer which should be sufficient for most stack traces.
func (g *GLogger) Error(msg string, fields ...zap.Field) {
	var stack = make([]byte, 1<<16)          // 64KB buffer
	stackSize := runtime.Stack(stack, false) // Get current goroutine's stack
	msg = fmt.Sprintf("%s\nstack: %s", msg, string(stack[:stackSize]))
	g.Logger.Error(msg, fields...)
}

// ErrorWithStack logs an error message with full stack traces from all goroutines.
// This is useful for debugging complex concurrency issues.
// Uses a 1MB buffer to accommodate potentially large stacks.
func (g *GLogger) ErrorWithStack(msg string, fields ...zap.Field) {
	var stack = make([]byte, 1<<20)         // 1MB buffer
	stackSize := runtime.Stack(stack, true) // Get all goroutines' stacks
	msg = fmt.Sprintf("%s\nfull stack: %s", msg, string(stack[:stackSize]))
	g.Logger.Error(msg, fields...)
}

// Info logs an informational message without stack traces.
// Maintains the original zap.Logger Info functionality.
func (g *GLogger) Info(msg string, fields ...zap.Field) {
	g.Logger.Info(msg, fields...)
}

// Warn logs a warning message without stack traces.
// Maintains the original zap.Logger Warn functionality.
func (g *GLogger) Warn(msg string, fields ...zap.Field) {
	g.Logger.Warn(msg, fields...)
}

// Debug logs a debug message without stack traces.
// Maintains the original zap.Logger Debug functionality.
func (g *GLogger) Debug(msg string, fields ...zap.Field) {
	g.Logger.Debug(msg, fields...)
}

// Fatal logs a fatal error message with the current goroutine's stack trace
// then terminates the program (os.Exit(1)).
// Uses a 64KB buffer which should be sufficient for most stack traces.
func (g *GLogger) Fatal(msg string, fields ...zap.Field) {
	var stack = make([]byte, 1<<16)
	stackSize := runtime.Stack(stack, false)
	msg = fmt.Sprintf("%s\nstack: %s", msg, string(stack[:stackSize]))
	g.Logger.Fatal(msg, fields...)
}

// Panic logs a panic message with the current goroutine's stack trace
// then panics. Uses a 64KB buffer which should be sufficient for most stack traces.
func (g *GLogger) Panic(msg string, fields ...zap.Field) {
	var stack = make([]byte, 1<<16)
	stackSize := runtime.Stack(stack, false)
	msg = fmt.Sprintf("%s\nstack: %s", msg, string(stack[:stackSize]))
	g.Logger.Panic(msg, fields...)
}
