package config

import "context"

// Adapter defines the interface for configuration adapters
// that can load configuration data from different sources
type Adapter interface {
	// Available checks if the configuration source is accessible
	// ctx: Context for cancellation and timeouts
	// resource: Optional resource identifiers
	Available(ctx context.Context, resource ...string) (ok bool)

	// Data loads and returns the configuration data
	// ctx: Context for cancellation and timeouts
	// Returns:
	//   data: The loaded configuration as key-value pairs
	//   err: Any error that occurred
	Data(ctx context.Context) (data map[string]any, err error)
}
