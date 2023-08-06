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

func NewBadRequestError(error string, data interface{}) *HttpError {
	return NewHttpError(constants.StatusBadRequest, error, data)
}

func NewInternalServerError(error string, data interface{}) *HttpError {
	return NewHttpError(constants.StatusInternalServerError, error, data)
}
