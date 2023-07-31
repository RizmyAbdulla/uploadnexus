package errors

type HttpError struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data,omitempty"`
}

func NewHttpError(code int, error string, data interface{}) *HttpError {
	return &HttpError{
		Code:  code,
		Error: error,
		Data:  data,
	}
}
