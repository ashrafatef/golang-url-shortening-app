package errors

import (
	"net/http"

	"github.com/ashrafatef/urlshortening/validations"
)

type ValidationError struct {
	Fields     []validations.ValidationError `json:"fields"`
	Message    string                        `json:"message"`
	StatusCode int                           `json:"status_code"`
}

func NewValidationError(fields []validations.ValidationError) *ValidationError {
	return &ValidationError{
		Fields:     fields,
		Message:    "Bad Request",
		StatusCode: http.StatusBadRequest,
	}
}

func (v ValidationError) Error() string {
	return v.Message
}
