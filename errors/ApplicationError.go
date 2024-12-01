package errors

type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewApplicationError(message string, statusCode int) *ApplicationError {
	return &ApplicationError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func (e *ApplicationError) Error() string {
	return e.Message
}