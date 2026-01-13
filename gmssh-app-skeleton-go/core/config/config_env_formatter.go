package config

// EnvFormatter implements configuration formatting with environment variable support
// It wraps another Formatter and prepends environment prefixes to section names
type EnvFormatter struct {
	Formatter
	env EnvString
}

// NewEnvFormatter creates a new EnvFormatter instance
// env: The environment string to prepend
// formatter: The underlying formatter to wrap
func NewEnvFormatter(env EnvString, formatter Formatter) *EnvFormatter {
	return &EnvFormatter{
		Formatter: formatter,
		env:       env,
	}
}

func (e *EnvFormatter) concatEnvSection(section string) string {
	// Note: Adds environment prefix handling
	return concatEnvVal(e.env, section)
}

// GetValue retrieves a configuration value with environment prefix
// section: The configuration section name
// Returns: A Formatter for the requested section
func (e *EnvFormatter) GetValue(section string) Formatter {
	sec := e.concatEnvSection(section)
	return e.Formatter.GetValue(sec)
}

// WithConfigEnvFormatterOptionFunc creates a config option that wraps
// the configuration with environment prefix support
// env: The environment string to use as prefix
func WithConfigEnvFormatterOptionFunc(env string) WithConfigOptionFunc {
	return func(c *Config) {
		nFormatter := NewEnvFormatter(NewEnvString(env), c.Cfg())
		c.CloneWithFormatter(nFormatter)
	}
}
