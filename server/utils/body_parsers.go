package utils

import (
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/gofiber/fiber/v2"
)

func ParseRequestBody(ctx *fiber.Ctx, data interface{}) *errors.HttpError {
	err := ctx.BodyParser(data)
	if err != nil {
		return errors.NewBadRequestError("invalid request body", err)
	}

	return nil
}
