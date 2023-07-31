package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func Validate(modelValidate interface{}) ([]map[string]interface{}, error) {
	var messages []map[string]interface{}

	validate := validator.New()

	err := validate.Struct(modelValidate)

	if err != nil {
		modelType := reflect.TypeOf(modelValidate)

		for _, err := range err.(validator.ValidationErrors) {

			field, _ := modelType.FieldByName(err.StructField())

			fieldName := getJSONKeyFromField(field)

			message := map[string]interface{}{
				"field": fieldName,
			}

			param := err.Param()
			if param != "" {
				message["message"] = "this field should be " + err.Tag() + " " + param
			} else {
				message["message"] = "this field is " + err.Tag()
			}

			messages = append(messages, message)
		}

		return messages, fmt.Errorf("%v", messages)
	}

	return messages, nil
}

func getJSONKeyFromField(field reflect.StructField) string {
	jsonTag := field.Tag.Get("json")
	if jsonTag != "" && jsonTag != "-" {
		parts := strings.Split(jsonTag, ",")
		return parts[0]
	}
	return field.Name
}
