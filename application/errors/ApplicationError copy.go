package errors

import "net/http"

type NotFoundError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

func (e *NotFoundError) Error() string {
	return e.Message
}
