package config

import (
	"context"
	"errors"
	"strings"
)

var (
	// GmCfg is the global configuration instance that can be accessed throughout the application
	GmCfg *Config
)

// Config represents the main configuration handler that combines both:
// - An adapter for loading configuration from various sources
// - A formatter for parsing and accessing the configuration data
type Config struct {
	adapter   Adapter   // Interface for configuration source (file, env, etc.)
	formatter Formatter // Interface for parsing and accessing config data
}

const (
	// DefaultConfigFileName is the default name for configuration files
	DefaultConfigFileName = "config"

	// DefaultSectionMinNum is the minimum number of sections required for recursive parsing
	// Sections are dot-separated keys (e.g., "a.b.c" has 3 sections)
	DefaultSectionMinNum = 2
)

// NewWithAdapter creates a new Config instance with specified adapter and formatter
// This allows custom configuration sources and parsing logic
func NewWithAdapter(adapter Adapter, formatter Formatter) *Config {
	return &Config{
		adapter:   adapter,
		formatter: formatter,
	}
}

// New creates a default Config instance with file adapter and JSON formatter
// It automatically loads and parses the configuration file
func New() (*Config, error) {
	ctx := context.Background()

	// Create default file adapter
	adapterFile, err := NewAdapterFile()
	if err != nil {
		return nil, err
	}

	// Verify configuration is available
	if !adapterFile.Available(ctx) {
		return nil, errors.New("configuration item not found")
	}

	// Load configuration data
	content, err := adapterFile.Data(ctx)
	if err != nil {
		return nil, errors.New("configuration file parsing format error")
	}

	// Create formatter with loaded content
	formatter := &ConfigFormatter{
		content: content,
	}

	return &Config{
		adapter:   adapterFile,
		formatter: formatter,
	}, nil
}

// CloneWithFormatter creates a new Config instance sharing the same adapter
// but with a different formatter. Useful for parsing the same config differently.
func (c *Config) CloneWithFormatter(formatter Formatter) *Config {
	c.formatter = formatter
	return c
}

// SetFormatter replaces the current formatter with a new one
func (c *Config) SetFormatter(formatter Formatter) {
	c.formatter = formatter
}

// GetAdapter returns the current configuration adapter
func (c *Config) GetAdapter() Adapter {
	return c.adapter
}

// MustData returns the raw configuration data, ignoring any errors
// Note: Use with caution as it may panic if data is not available
func (c *Config) MustData(ctx context.Context) map[string]any {
	data, _ := c.adapter.Data(ctx)
	return data
}

// Cfg returns the current formatter for accessing configuration values
func (c *Config) Cfg() Formatter {
	return c.formatter
}

// ConfigFormatter implements the Formatter interface for JSON configuration
// It supports dot notation for nested keys (e.g., "a.b.c")
// Note: The configuration must strictly follow the expected format
type ConfigFormatter struct {
	content any      // The parsed configuration data
	value   []string // Current key path (split by dots)
}

// GetValue sets up the formatter to access a specific configuration path
// The section parameter uses dot notation for nested keys (e.g., "server.port")
// Warning: Chained calls will overwrite previous path selections
func (c *ConfigFormatter) GetValue(section string) Formatter {
	// Split the dot-separated path into individual keys
	c.value = strings.Split(section, ".")
	return c
}

// isRecursion checks if the current path requires recursive lookup
func (c *ConfigFormatter) isRecursion() bool {
	return len(c.value) >= DefaultSectionMinNum
}

// getValue retrieves a configuration value of specified type from a map
func getValue[T any](k string, content any) (v T, err error) {
	comma, ok := content.(map[string]any)
	if !ok {
		err = errors.New("configuration item not found")
		return
	}
	v, ok = comma[k].(T)
	if !ok {
		err = errors.New("configuration item not found")
		return
	}
	return
}

// getAnyValue retrieves any configuration value from a map
func getAnyValue(k string, content any) (v any, err error) {
	comma, ok := content.(map[string]any)
	if !ok {
		err = errors.New("configuration item not found")
		return
	}
	v, ok = comma[k]
	if !ok {
		err = errors.New("configuration item not found")
		return
	}
	return
}

// getRecursionValue recursively navigates through nested maps to find a value
// Note: Not thread-safe - caller should handle synchronization if needed
func getRecursionValue[T any](parts []string, content any) (v T, err error) {
	var data map[string]any
	var ok bool

	// Recursively navigate through nested maps
	for i, cfg := range parts {
		if i == len(parts)-1 {
			// Last part - return the actual value
			v, err = getValue[T](cfg, data)
			if !ok {
				err = errors.New("configuration item not found")
				return
			}
		} else {
			// Intermediate part - navigate deeper
			v1, err1 := getAnyValue(cfg, content)
			if err1 != nil {
				err = err1
				return
			}

			data, ok = v1.(map[string]any)
			content = data
		}
	}
	return
}

// String returns a string configuration value
func (c *ConfigFormatter) String() (v string, err error) {
	if !c.isRecursion() {
		// Simple case - single key lookup
		v, err = getValue[string](c.value[0], c.content)
		return
	}

	// Complex case - recursive lookup
	return getRecursionValue[string](c.value, c.content)
}

// StringWithOutErr returns string value or empty string on error
func (c *ConfigFormatter) StringWithOutErr() string {
	val, _ := c.String()
	return val
}

// StringWithDefault returns string value or default if not found
func (c *ConfigFormatter) StringWithDefault(df string) string {
	val, err := c.String()
	if err != nil {
		return df
	}
	return val
}

// Int returns an integer configuration value
// Note: JSON numbers are parsed as float64, so we convert
func (c *ConfigFormatter) Int() (v int, err error) {
	if !c.isRecursion() {
		v1, err := getValue[float64](c.value[0], c.content)
		return int(v1), err
	}

	v1, err := getRecursionValue[float64](c.value, c.content)
	return int(v1), err
}

// Float64 returns a float64 configuration value
func (c *ConfigFormatter) Float64() (v float64, err error) {
	if !c.isRecursion() {
		return getValue[float64](c.value[0], c.content)
	}
	return getRecursionValue[float64](c.value, c.content)
}

// IntWithOutErr returns integer value or 0 on error
func (c *ConfigFormatter) IntWithOutErr() int {
	val, _ := c.Int()
	return val
}

// IntWithDefault returns integer value or default if not found
func (c *ConfigFormatter) IntWithDefault(df int) int {
	val, err := c.Int()
	if err != nil {
		return df
	}
	return val
}

// Bool returns a boolean configuration value
func (c *ConfigFormatter) Bool() (v bool, err error) {
	if !c.isRecursion() {
		v, err = getValue[bool](c.value[0], c.content)
		return
	}
	return getRecursionValue[bool](c.value, c.content)
}

// BoolWithOutErr returns boolean value or false on error
func (c *ConfigFormatter) BoolWithOutErr() bool {
	val, _ := c.Bool()
	return val
}

// Map returns a map configuration value
func (c *ConfigFormatter) Map() (v map[string]any, err error) {
	if !c.isRecursion() {
		v, err = getValue[map[string]any](c.value[0], c.content)
		return
	}
	return getRecursionValue[map[string]any](c.value, c.content)
}

// MapWithOutErr returns map value or empty map on error
func (c *ConfigFormatter) MapWithOutErr() map[string]any {
	val, err := c.Map()
	if err != nil {
		return make(map[string]any)
	}
	return val
}

// List returns a slice configuration value
// Note: Does not support deep nested processing
func (c *ConfigFormatter) List() (v []any, err error) {
	if !c.isRecursion() {
		v, err = getValue[[]any](c.value[0], c.content)
		return
	}
	return getRecursionValue[[]any](c.value, c.content)
}

// ListWithOutErr returns slice value or empty slice on error
func (c *ConfigFormatter) ListWithOutErr() []any {
	val, err := c.List()
	if err != nil {
		return []any{}
	}
	return val
}
