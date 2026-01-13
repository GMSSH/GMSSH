package gerror

import "fmt"

// localCode is an internal implementation of the Code interface.
// It's used for representing error codes within the package.
type localCode struct {
	code    int    // Error code (typically an integer)
	message string // Brief description of the error
	detail  any    // Additional error details (can be any type)
	i18n    string // Internationalization template for translation
}

// Code returns the integer error code.
func (c localCode) Code() int {
	return c.code
}

// Message returns the brief error description.
func (c localCode) Message() string {
	return c.message
}

// I18n returns the internationalization template string.
// This is used for localized error messages.
func (c localCode) I18n() string {
	return c.i18n
}

// Detail returns additional error information.
// This field is designed as an extension point for error codes.
func (c localCode) Detail() any {
	return c.detail
}

// String formats the error code into a human-readable string.
// Format includes code, message and detail (if available):
// - With detail: "code:message detail"
// - Without detail: "code:message"
// - Message only: "code"
func (c localCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}

// Error implements the error interface by calling String().
// This allows localCode to be used as a standard error.
func (c localCode) Error() string {
	return c.String()
}

// HttpError represents an HTTP error response.
// It's simpler than localCode and used specifically for HTTP responses.
type HttpError struct {
	Code    int    // HTTP status code
	Message string // Error message to be returned to client
}

// Error implements the error interface for HttpError.
// Returns just the message without the code (unlike localCode).
func (c HttpError) Error() string {
	return c.Message
}
