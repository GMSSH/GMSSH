package config

// WithConfigOptionFunc defines a function type for configuring a Config instance.
// These functions are used to apply optional configuration settings.
type WithConfigOptionFunc func(*Config)

// Formatter is an interface that defines methods for accessing and parsing configuration values.
// It provides type-safe access to configuration data with various return formats.
// Note: Error messages currently need internationalization support (i18n).
type Formatter interface {
	// String returns the configuration value as string.
	// Returns error if value doesn't exist or has wrong type.
	String() (string, error)

	// StringWithOutErr returns the configuration value as string.
	// Returns empty string if value doesn't exist or has wrong type.
	StringWithOutErr() string

	// StringWithDefault returns the configuration value as string.
	// Returns the default value (df) if value doesn't exist or has wrong type.
	StringWithDefault(df string) string

	// Int returns the configuration value as integer.
	// Returns error if value doesn't exist or has wrong type.
	Int() (int, error)

	// Float64 returns the configuration value as float64.
	// Returns error if value doesn't exist or has wrong type.
	Float64() (v float64, err error)

	// IntWithOutErr returns the configuration value as integer.
	// Returns 0 if value doesn't exist or has wrong type.
	IntWithOutErr() int

	// IntWithDefault returns the configuration value as integer.
	// Returns the default value (df) if value doesn't exist or has wrong type.
	IntWithDefault(df int) int

	// Bool returns the configuration value as boolean.
	// Returns error if value doesn't exist or has wrong type.
	Bool() (bool, error)

	// BoolWithOutErr returns the configuration value as boolean.
	// Returns false if value doesn't exist or has wrong type.
	BoolWithOutErr() bool

	// Map returns the configuration value as map[string]any.
	// Returns error if value doesn't exist or has wrong type.
	Map() (v map[string]any, err error)

	// MapWithOutErr returns the configuration value as map[string]any.
	// Returns empty map if value doesn't exist or has wrong type.
	MapWithOutErr() map[string]any

	// List returns the configuration value as []any slice.
	// Returns error if value doesn't exist or has wrong type.
	List() (v []any, err error)

	// ListWithOutErr returns the configuration value as []any slice.
	// Returns empty slice if value doesn't exist or has wrong type.
	ListWithOutErr() []any

	// GetValue selects a nested configuration section using dot notation.
	// Returns a new Formatter instance for the specified section.
	// Example: GetValue("server.port") accesses the port value under server section.
	GetValue(section string) Formatter
}
