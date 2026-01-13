package gvalid

// ValidatorErrorMessage provides base functionality for validation error messages.
// It is designed to be embedded in other validator types to provide consistent
// error message handling capabilities.
//
// Key Features:
// - Supports custom error messages
// - Provides standardized validation error creation
// - Implements common validator helper methods
type ValidatorErrorMessage struct {
	message string // Custom error message template
}

// Valid is a placeholder method that satisfies validator interfaces.
// Concrete validator implementations should override this with actual validation logic.
//
// Returns:
//   - bool: Always returns false as this is just a base implementation
func (v *ValidatorErrorMessage) Valid(value any) bool {
	// Base implementation always returns false
	// Real validators should override this method
	return false
}

// Message sets a custom error message template for the validator.
// This message will be used when creating validation errors.
//
// Parameters:
//   - message: The custom error message template
//
// Example:
//
//	validator.Message("Custom error: field %s is invalid")
func (v *ValidatorErrorMessage) Message(message string) {
	v.message = message
}

// NewValidationError creates a standardized validation error with optional
// custom message formatting. If a custom message was set via Message(),
// it will override the default format string.
//
// Parameters:
//   - field: The name of the field being validated
//   - format: Default error message format string
//   - args: Arguments for the format string
//
// Returns:
//   - *ValidationError: A new validation error instance
//
// Example:
//
//	err := v.NewValidationError("username", "invalid format")
func (v *ValidatorErrorMessage) NewValidationError(field, format string, args ...any) *ValidationError {
	// Create base validation error
	err := NewValidationError(field, format, args...)

	// Apply custom message if set
	return WithErrorMessage(err, v.message)
}
