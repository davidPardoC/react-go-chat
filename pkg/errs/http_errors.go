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

func NewNotFoundError(messages ...string) *Error {
	if len(messages) < 1 {
		messages[0] = "Not Found"
	}
	return &Error{StatusCode: http.StatusNotFound, Message: messages[0]}
}

func NewInternalServerError(messages ...string) *Error {
	if len(messages) < 1 {
		messages[0] = "Internal server error"
	}
	return &Error{StatusCode: http.StatusInternalServerError, Message: messages[0]}
}
