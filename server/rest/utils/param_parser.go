package utils

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/exceptions"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func GetParamId(ctx *fiber.Ctx) (string, *exceptions.HttpError) {
	id := ctx.Params("id")
	unescapedId, err := url.QueryUnescape(id)
	if err != nil {
		return "", exceptions.NewBadRequestError("invalid parameter", err)
	}

	if unescapedId == "" {
		return "", exceptions.NewBadRequestError("invalid parameter", nil)
	}

	return unescapedId, nil
}

func GetParamName(ctx *fiber.Ctx) (string, *exceptions.HttpError) {
	name := ctx.Params("name")
	unescapedName, err := url.QueryUnescape(name)
	if err != nil {
		return "", exceptions.NewBadRequestError("invalid parameter", nil)
	}

	if unescapedName == "" {
		return "", exceptions.NewBadRequestError("invalid parameter", nil)
	}
	return unescapedName, nil
}
