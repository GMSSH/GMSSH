package ghttp

import (
	"context"

	"github.com/DemonZack/simplejrpc-go/core/gerror"
)

// Response constants define the standard JSON response field names
const (
	ResponseDataKey     = "data"     // Key for primary response data
	ResponseCodeKey     = "code"     // Key for status/error code
	ResponseMessageKey  = "msg"      // Key for status/error message
	ResponseMetaKey     = "meta"     // Key for metadata
	ResponseCloseKey    = "close"    // Key for connection close flag
	ResponseEndpointKey = "endpoint" // Key for endpoint/API name
	ResponseExtraKey    = "extra"    // Key for additional custom data
)

// IResponse defines the interface for building HTTP/WebSocket responses
type IResponse interface {
	// Sets the response status code
	Code(code int) IResponse

	// Sets the response message
	Message(msg string) IResponse

	// Sets both response data and metadata
	Data(body any, meta map[string]any) IResponse

	// Sets only the response body data
	Body(body any) IResponse

	// Marks the response as an event with API endpoint
	Event(api string) IResponse

	// Sets connection close flag
	End(close int) IResponse

	// Finalizes and returns the JSON response
	JSON(codes ...ResponseOptionFunc) IResponse

	// Adds extra custom data to the response
	Extra(extra any) IResponse

	// Sets both code and message from an Exception
	CodeMessage(code gerror.Exception) IResponse

	// Returns the complete response as a map
	GetResponse() map[string]any
}

// ResponseOptionFunc defines functions for configuring responses
type ResponseOptionFunc func(IResponse)

// WithResponseCode creates a ResponseOptionFunc that sets the response code
func WithResponseCode(code int) ResponseOptionFunc {
	return func(iResponse IResponse) {
		iResponse.Code(code)
	}
}

// BaseResponse provides core response building functionality
type BaseResponse struct {
	data     map[string]any // Primary response fields
	metaData map[string]any // Metadata fields
}

// NewBaseResponse creates a new BaseResponse instance
func NewBaseResponse() *BaseResponse {
	return &BaseResponse{
		data:     make(map[string]any),
		metaData: make(map[string]any),
	}
}

// CodeMessage sets both code and message from an Exception
func (r *BaseResponse) CodeMessage(code gerror.Exception) IResponse {
	r.data[ResponseCodeKey] = code.Code()
	r.data[ResponseMessageKey] = code.I18n()
	return r
}

// Code sets the response status code
func (r *BaseResponse) Code(code int) IResponse {
	r.data[ResponseCodeKey] = code
	return r
}

// Message sets the response message
func (r *BaseResponse) Message(msg string) IResponse {
	r.data[ResponseMessageKey] = msg
	return r
}

// Data sets both the response body and metadata
func (r *BaseResponse) Data(body any, meta map[string]any) IResponse {
	r.data[ResponseDataKey] = body
	for k, v := range meta {
		r.metaData[k] = v
	}
	return r
}

// Body sets only the response body data
func (r *BaseResponse) Body(body any) IResponse {
	r.data[ResponseDataKey] = body
	return r
}

// Event marks the response as an API event
func (r *BaseResponse) Event(api string) IResponse {
	r.metaData[ResponseEndpointKey] = api
	r.metaData[ResponseCloseKey] = 0 // Default to keep connection open
	return r
}

// End sets the connection close flag
func (r *BaseResponse) End(close int) IResponse {
	r.metaData[ResponseCloseKey] = close
	return r
}

// Extra adds custom data to the response
func (r *BaseResponse) Extra(extra any) IResponse {
	r.metaData[ResponseExtraKey] = extra
	return r
}

// JSON finalizes and returns the JSON response
func (r *BaseResponse) JSON(opts ...ResponseOptionFunc) IResponse {
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// GetResponse constructs the final response map
func (r *BaseResponse) GetResponse() map[string]any {
	response := make(map[string]any)
	// Copy standard fields
	response[ResponseDataKey] = r.data[ResponseDataKey]
	response[ResponseCodeKey] = r.data[ResponseCodeKey]
	response[ResponseMessageKey] = r.data[ResponseMessageKey]
	response[ResponseMetaKey] = r.metaData
	return response
}

// HTTPResponse represents an HTTP-specific response
type HTTPResponse struct {
	*BaseResponse
	Context context.Context // Request context
}

// WsResponse represents a WebSocket-specific response
type WsResponse struct {
	*BaseResponse
}

// NewWsResponse creates a new WebSocket response
func NewWsResponse() *WsResponse {
	return &WsResponse{
		BaseResponse: NewBaseResponse(),
	}
}

// JSON implements WebSocket-specific JSON response
func (w *WsResponse) JSON(codes ...ResponseOptionFunc) IResponse {
	// WebSocket-specific JSON handling
	// TODO: Implement WebSocket-specific logic
	return w
}

// Response provides a generic response adapter
type Response struct {
	adapter IResponse // Underlying response implementation
}

// NewDefaultHttpResponse creates a default HTTP response
func NewDefaultHttpResponse() IResponse {
	return &Response{
		adapter: NewBaseResponse(),
	}
}

// NewResponseWithAdapter creates a response with custom adapter
func NewResponseWithAdapter(adapter IResponse) IResponse {
	return &Response{
		adapter: adapter,
	}
}

// All methods below delegate to the underlying adapter
func (r *Response) CodeMessage(code gerror.Exception) IResponse {
	return r.adapter.CodeMessage(code)
}

func (r *Response) Code(code int) IResponse {
	return r.adapter.Code(code)
}

func (r *Response) Message(msg string) IResponse {
	return r.adapter.Message(msg)
}

func (r *Response) Data(body any, meta map[string]any) IResponse {
	return r.adapter.Data(body, meta)
}

func (r *Response) Body(body any) IResponse {
	return r.adapter.Body(body)
}

func (r *Response) Event(api string) IResponse {
	return r.adapter.Event(api)
}

func (r *Response) End(close int) IResponse {
	return r.adapter.End(close)
}

func (r *Response) Extra(extra any) IResponse {
	return r.adapter.Extra(extra)
}

func (r *Response) JSON(codes ...ResponseOptionFunc) IResponse {
	return r.adapter.JSON(codes...)
}

func (r *Response) GetResponse() map[string]any {
	return r.adapter.GetResponse()
}
