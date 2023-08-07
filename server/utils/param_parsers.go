package utils

import (
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func GetParamId(ctx *fiber.Ctx) (string, *errors.HttpError) {
	id := ctx.Params("id")
	unescapedId, err := url.QueryUnescape(id)
	if err != nil {
		return "", errors.NewBadRequestError("invalid parameter")
	}

	if unescapedId == "" {
		return "", errors.NewBadRequestError("invalid parameter")
	}

	return unescapedId, nil
}

func GetParamName(ctx *fiber.Ctx) (string, *errors.HttpError) {
	name := ctx.Params("name")
	unescapedName, err := url.QueryUnescape(name)
	if err != nil {
		return "", errors.NewBadRequestError("invalid parameter")
	}

	if unescapedName == "" {
		return "", errors.NewBadRequestError("invalid parameter")
	}
	return unescapedName, nil
}

func GetParamBucketName(ctx *fiber.Ctx) (string, *errors.HttpError) {
	bucketName := ctx.Params("bucket_name")
	unescapedBucketName, err := url.QueryUnescape(bucketName)
	if err != nil {
		return "", errors.NewBadRequestError("invalid parameter")
	}

	if unescapedBucketName == "" {
		return "", errors.NewBadRequestError("invalid parameter")
	}
	return unescapedBucketName, nil
}

func GetParamWildcard(ctx *fiber.Ctx) (string, *errors.HttpError) {
	wildCard := ctx.Params("*")
	unescapedWildcard, err := url.QueryUnescape(wildCard)
	if err != nil {
		return "", errors.NewBadRequestError("invalid parameter")
	}

	if unescapedWildcard == "" {
		return "", errors.NewBadRequestError("invalid parameter")
	}

	return unescapedWildcard, nil
}
