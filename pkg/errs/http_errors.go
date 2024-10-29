package errs

import "net/http"

type Error struct {
	StatusCode int
	Message    string
	Err        error
}

func NewError(statusCode int, message string) *Error {
	return &Error{StatusCode: statusCode, Message: message}
}

func NewUnauthorizedError(message string) *Error {
	return &Error{StatusCode: http.StatusUnauthorized, Message: message}
}
