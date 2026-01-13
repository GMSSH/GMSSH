package gvalid

import (
	"fmt"
	"strings"
)

// WithErrorMessage updates the error message of a ValidationError if a new message is provided.
// This allows for customizing error messages after validation.
//
// Parameters:
//   - err: The original validation error (required)
//   - message: The new error message (optional)
//
// Returns:
//   - *ValidationError: The modified error (or original if no message provided)
//
// Example:
//
//	err := NewValidationError("email", "invalid format")
//	err = WithErrorMessage(err, "Email format is invalid")
func WithErrorMessage(err *ValidationError, message string) *ValidationError {
	if message != "" {
		err.Message = message
	}
	return err
}

// ValidationError represents a single validation error containing:
// - The field name that failed validation
// - A descriptive error message
type ValidationError struct {
	Field   string // Name of the field that failed validation
	Message string // Description of the validation failure
}

// NewValidationError creates a new ValidationError with formatted message.
//
// Parameters:
//   - field: Name of the invalid field
//   - format: Message format string (supports fmt.Sprintf formatting)
//   - args: Arguments for the format string
//
// Returns:
//   - *ValidationError: New validation error instance
//
// Example:
//
//	err := NewValidationError("age", "must be between %d and %d", 18, 99)
func NewValidationError(field, format string, args ...any) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: fmt.Sprintf(format, args...),
	}
}

// WithMessage returns just the error message without field name.
// Useful when only the message text is needed.
func (e *ValidationError) WithMessage() error {
	return fmt.Errorf("%v", e.Message)
}

// Error implements the error interface for ValidationError.
// Returns a string combining field name and error message.
// Format: "field: error message"
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors represents a collection of validation errors.
// This is used when multiple fields fail validation.
type ValidationErrors []*ValidationError

// Error implements the error interface for ValidationErrors.
// Joins all error messages with semicolons into a single string.
// Example: "email: invalid format; age: must be positive"
func (es ValidationErrors) Error() string {
	var sb strings.Builder
	for _, e := range es {
		sb.WriteString(e.Error())
		sb.WriteString("; ")
	}
	// Remove trailing semicolon and space
	return strings.TrimSuffix(sb.String(), "; ")
}
