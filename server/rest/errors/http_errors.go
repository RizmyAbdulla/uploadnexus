package errors

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewHttpError(code int, message string, error string) *HttpError {
	return &HttpError{
		Code:    code,
		Message: message,
		Error:   error,
	}
}
