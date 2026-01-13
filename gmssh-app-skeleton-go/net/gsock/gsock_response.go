package gsock

import "net/http"

// Meta contains WebSocket metadata for message handling
type Meta struct {
	Endpoint string `json:"endpoint"` // The API endpoint/path this response corresponds to
	Close    int    `json:"close"`    // Flag indicating if connection should close (0=keep open, 1=close)
}

// Response represents a standardized WebSocket response format
type Response struct {
	Code    int    `json:"code"` // HTTP-style status code
	Data    any    `json:"data"` // Primary response payload
	Message string `json:"msg"`  // Human-readable status message
	Meta    *Meta  `json:"meta"` // WebSocket-specific metadata
}

// NewResponse creates a new Response with default values:
// - Code: 200 (StatusOK)
// - Message: "OK"
// - Initialized Meta struct
func NewResponse() *Response {
	return &Response{
		Meta:    &Meta{},
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}
}

// SetEndpoint specifies the API endpoint this response corresponds to
// This helps clients route responses to the correct handlers
func (r *Response) SetEndpoint(endpoint string) {
	r.Meta.Endpoint = endpoint
}

// SetClose controls the WebSocket connection close behavior
// Values:
//
//	0 - Keep connection open (default)
//	1 - Close connection after this message
func (r *Response) SetClose(close int) {
	r.Meta.Close = close
}

// Helper methods for common status codes

// WithSuccess configures a successful (200 OK) response
func (r *Response) WithSuccess(data any) *Response {
	r.Code = http.StatusOK
	r.Message = http.StatusText(http.StatusOK)
	r.Data = data
	return r
}

// WithError configures an error response
func (r *Response) WithError(code int, message string) *Response {
	r.Code = code
	r.Message = message
	r.Data = nil
	return r
}

// WithData sets both the response data and metadata
func (r *Response) WithData(data any, endpoint string) *Response {
	r.Data = data
	r.SetEndpoint(endpoint)
	return r
}
