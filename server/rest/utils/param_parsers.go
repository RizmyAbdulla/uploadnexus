package utils

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/errors"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func GetParamId(ctx *fiber.Ctx) (string, *errors.HttpError) {
	id := ctx.Params("id")
	unescapedId, err := url.QueryUnescape(id)
	if err != nil {
		return "", errors.NewBadRequestError("invalid parameter", err)
	}

	if unescapedId == "" {
		return "", errors.NewBadRequestError("invalid parameter", nil)
	}

	return unescapedId, nil
}

func GetParamName(ctx *fiber.Ctx) (string, *errors.HttpError) {
	name := ctx.Params("name")
	unescapedName, err := url.QueryUnescape(name)
	if err != nil {
		return "", errors.NewBadRequestError("invalid parameter", nil)
	}

	if unescapedName == "" {
		return "", errors.NewBadRequestError("invalid parameter", nil)
	}
	return unescapedName, nil
}
