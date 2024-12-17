package errors

import "net/http"

type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewApplicationError(message string) *ApplicationError {
	return &ApplicationError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func (e *ApplicationError) Error() string {
	return e.Message
}
