package models

type GeneralResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewGeneralResponse(code int, message string, data interface{}) *GeneralResponse {
	return &GeneralResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
