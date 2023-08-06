package utils

import (
	"encoding/json"
	"github.com/ArkamFahry/uploadnexus/server/errors"
)

func ParseRequestBody(body []byte, data interface{}) *errors.HttpError {
	err := json.Unmarshal(body, data)
	if err != nil {
		return errors.NewBadRequestError("invalid request body", err)
	}

	return nil
}
