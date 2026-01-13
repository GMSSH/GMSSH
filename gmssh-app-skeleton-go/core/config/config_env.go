package config

import (
	"fmt"
	"strings"
)

// EnvString represents an environment identifier (test/prod/dev)
// Used to prefix configuration keys with environment-specific values
type EnvString string

const (
	// ENV_TEST represents the test environment
	ENV_TEST EnvString = "test"
	
	// ENV_PROD represents the production environment
	ENV_PROD EnvString = "prod"
	
	// ENV_DEV represents the development environment
	ENV_DEV  EnvString = "dev"
)

// NewEnvString creates a new EnvString from a string value
// Validates and normalizes the environment string
func NewEnvString(val string) EnvString {
	return EnvString(val)
}

// String returns the normalized string representation of the environment
// Defaults to production environment if invalid
func (e EnvString) String() string {
	switch strings.ToLower(string(e)) {
	case string(ENV_TEST):
		return string(ENV_TEST)
	case string(ENV_PROD):
		return string(ENV_PROD)
	case string(ENV_DEV):
		return string(ENV_DEV)
	default:
		return string(ENV_PROD)
	}
}

// concatEnvVal combines an environment prefix with a configuration key
// Returns a string in format "env.key"
func concatEnvVal(env EnvString, val string) string {
	return fmt.Sprintf("%s.%s", env.String(), val)
}
