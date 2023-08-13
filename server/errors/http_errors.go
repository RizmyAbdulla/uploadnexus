package errors

import "github.com/ArkamFahry/uploadnexus/server/constants"

type HttpError struct {
	Code    int         `json:"code"`
	Error   string      `json:"error"`
	Message interface{} `json:"message,omitempty"`
}

func NewHttpError(code int, error string, message interface{}) *HttpError {
	return &HttpError{
		Code:    code,
		Error:   error,
		Message: message,
	}
}

func NewBadRequestError(message interface{}) *HttpError {
	return NewHttpError(constants.StatusBadRequest, "bad request", message)
}

func NewNotFoundError(message interface{}) *HttpError {
	return NewHttpError(constants.StatusNotFound, "resource not found", message)
}

func NewInternalServerError(message interface{}) *HttpError {
	return NewHttpError(constants.StatusInternalServerError, "internal server error", message)
}

func NewInvalidMediaTypeError(message interface{}) *HttpError {
	return NewHttpError(constants.StatusInvalidMediaType, "invalid media type", message)
}
