package models

type Application struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}
