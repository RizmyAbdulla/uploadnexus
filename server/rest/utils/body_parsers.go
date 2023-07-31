package utils

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/constants"
	"github.com/ArkamFahry/uploadnexus/server/rest/errors"
	"github.com/gofiber/fiber/v2"
)

func ParseRequestBody(ctx *fiber.Ctx, data interface{}) *errors.HttpError {
	err := ctx.BodyParser(data)
	if err != nil {
		return errors.NewHttpError(constants.StatusBadRequest, "invalid request body", err)
	}

	return nil
}
