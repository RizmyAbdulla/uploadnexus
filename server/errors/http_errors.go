package errors

import "github.com/ArkamFahry/uploadnexus/server/constants"

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

func NewBadRequestError(data interface{}) *HttpError {
	return NewHttpError(constants.StatusBadRequest, "bad request error", data)
}

func NewNotFoundError(data interface{}) *HttpError {
	return NewHttpError(constants.StatusNotFound, "not found error", data)
}

func NewInternalServerError(data interface{}) *HttpError {
	return NewHttpError(constants.StatusInternalServerError, "internal server error", data)
}
