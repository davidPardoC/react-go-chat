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

func NewConflictError(messages ...string) *Error {
	if len(messages) < 1 {
		messages[0] = "Conflict"
	}
	return &Error{StatusCode: http.StatusConflict, Message: messages[0]}
}
