package core

import (
	"go.uber.org/zap"

	"github.com/DemonZack/simplejrpc-go/core/config"
	"github.com/DemonZack/simplejrpc-go/core/glog"
	"github.com/DemonZack/simplejrpc-go/core/gvalid"
)

// WithContainerFunc defines a function type for configuring the container
type WithContainerFunc func(*container)

// WithContainerValidOption returns a WithContainerFunc that sets the validator
func WithContainerValidOption(valid *gvalid.StructWalker) WithContainerFunc {
	return func(c *container) {
		c.valid = valid
	}
}

// WithContainerLoggerOption returns a WithContainerFunc that sets the logger
func WithContainerLoggerOption(logger *zap.Logger) WithContainerFunc {
	return func(c *container) {
		c.logger = logger
	}
}

// WithContainerConfigOption returns a WithContainerFunc that sets the config
func WithContainerConfigOption(config *config.Config) WithContainerFunc {
	return func(c *container) {
		c.config = config
	}
}

// container is the core dependency container implementation
type container struct {
	logger *zap.Logger          // Application logger
	config *config.Config       // Application configuration
	valid  *gvalid.StructWalker // Validation component
}

// Clone creates a new container instance with optional overrides
func (c *container) Clone(opts ...WithContainerFunc) IContainer {
	newContainer := *c // Create a copy
	for _, opt := range opts {
		opt(&newContainer)
	}
	return &newContainer
}

// NewContainer creates a new container instance with optional configurations
func NewContainer(opts ...WithContainerFunc) IContainer {
	c := &container{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Log returns the zap.Logger instance
func (c *container) Log() *zap.Logger {
	return c.logger
}

// GLog returns a wrapped GLogger instance
func (c *container) GLog() *glog.GLogger {
	return glog.NewGLogger(c.logger)
}

// Cfg returns the configuration instance
func (c *container) Cfg() *config.Config {
	return c.config
}

// CfgFmt returns the configuration formatter
func (c *container) CfgFmt() config.Formatter {
	return c.config.Cfg()
}

// Valid returns the validation walker instance
func (c *container) Valid() *gvalid.StructWalker {
	return c.valid
}

// InitContainer initializes and returns the global container with default dependencies
func InitContainer(opts ...config.WithConfigOptionFunc) IContainer {
	// Initialize configuration
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	// Apply configuration options
	for _, opt := range opts {
		opt(cfg)
	}

	// Initialize logger from config
	loggerConfig, err := cfg.Cfg().GetValue("logger").Map()
	if err != nil {
		panic(err)
	}

	logConfig, err := glog.LoadConfig(loggerConfig)
	if err != nil {
		panic(err)
	}

	logger, err := glog.NewLogger(logConfig)
	if err != nil {
		panic(err)
	}

	// Initialize validation components
	visitor := gvalid.NewValidatorVisitor()

	// Register default validators
	visitor.RegisterValidator("required", &gvalid.RequiredValidator{})
	visitor.RegisterValidator("min_length", &gvalid.MinLengthValidator{})
	visitor.RegisterValidator("range", &gvalid.RangeValidator{})

	// Create struct walker with "validate" tag
	walker := gvalid.NewStructWalker(visitor, "validate")

	// Initialize global container
	Container = NewContainer(
		WithContainerLoggerOption(logger),
		WithContainerConfigOption(cfg),
		WithContainerValidOption(walker),
	)
	return Container
}

// GetValueStringFormConfigWithOutErr safely gets a string config value (returns empty string on error)
func GetValueStringFormConfigWithOutErr(section string) string {
	return Container.CfgFmt().GetValue(section).StringWithOutErr()
}

// GetValueStringFormConfigWithErr gets a string config value (panics on error)
func GetValueStringFormConfigWithErr(section string) string {
	val, err := Container.CfgFmt().GetValue(section).String()
	if err != nil {
		panic(err)
	}
	return val
}
