package errors

import (
	"fmt"
	"net/http"
)

// Error defines our custom error type
type Error struct {
	stacktrace error
	msg        string
	code       ErrorCode
}

// ErrorCode defines our internal error type
type ErrorCode uint8

// defines our internal error codes
// TODO implement internal errors for each HTTP response code
const (
	ErrorUnknown ErrorCode = iota
	ErrorNotFound
	ErrorInvalidArgument
	ErrorUnauthorized
	ErrorServerFault
)

// httpResponses defines our HTTP error responses in the case of an internal error
var httpResponses = map[ErrorCode]http.ConnState{
	ErrorUnknown:         http.StatusInternalServerError,
	ErrorNotFound:        http.StatusNotFound,
	ErrorInvalidArgument: http.StatusBadRequest,
	ErrorUnauthorized:    http.StatusUnauthorized,
	ErrorServerFault:     http.StatusInternalServerError,
}

// WrapError wraps the error & throws up stack
func WrapError(stacktrace error, code ErrorCode, msg string, arg ...interface{}) error {
	return &Error{
		code:       code,
		stacktrace: stacktrace,
		msg:        fmt.Sprintf(msg, arg...),
	}
}

// UnwrapError unwraps our error
func (e *Error) UnwrapError() error {
	return e.stacktrace
}

// NewError creates a new error
func NewError(code ErrorCode, msg string, arg ...interface{}) error {
	return WrapError(nil, code, msg, arg...)
}

// Error returns our error message
func (e *Error) Error() string {
	if e.stacktrace != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.stacktrace)
	}
	return e.msg
}

// ErrorCode returns our error code
func (e *Error) ErrorCode() ErrorCode {
	return e.code
}
