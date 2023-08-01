package utils

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/exceptions"
	"github.com/gofiber/fiber/v2"
)

func ParseRequestBody(ctx *fiber.Ctx, data interface{}) *exceptions.HttpError {
	err := ctx.BodyParser(data)
	if err != nil {
		return exceptions.NewBadRequestError("invalid request body", err)
	}

	return nil
}
