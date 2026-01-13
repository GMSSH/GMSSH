package core

import (
	"go.uber.org/zap"

	"github.com/DemonZack/simplejrpc-go/core/config"
	"github.com/DemonZack/simplejrpc-go/core/glog"
	"github.com/DemonZack/simplejrpc-go/core/gvalid"
)

// Container is the global dependency container that provides centralized access to:
// - Application logging
// - Configuration management
// - Validation utilities
//
// This serves as the single source of truth for core application dependencies.
// While global variables are generally discouraged, this pattern is acceptable for:
// - Framework-level dependencies
// - Singleton components
// - Core infrastructure that needs app-wide access
//
// Usage Guidelines:
// 1. Initialize once at application startup using InitContainer()
// 2. Access through the provided interface methods
// 3. For testing, use Clone() to create isolated instances
var Container IContainer

// IContainer defines the interface for the dependency container.
// It provides access to fundamental application services while
// maintaining loose coupling between components.
type IContainer interface {
	// Log returns the core zap.Logger instance for structured logging
	Log() *zap.Logger

	// GLog returns a wrapped GLogger with enhanced functionality
	GLog() *glog.GLogger

	// Cfg returns the raw configuration manager
	Cfg() *config.Config

	// CfgFmt returns the configuration formatter for type-safe config access
	CfgFmt() config.Formatter

	// Valid returns the struct validation walker instance
	Valid() *gvalid.StructWalker

	// Clone creates a new container instance with optional overrides.
	// This is particularly useful for:
	// - Testing scenarios
	// - Request-specific customization
	// - Environment-specific configurations
	Clone(opts ...WithContainerFunc) IContainer
}
