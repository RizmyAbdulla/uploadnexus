package minio

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/errors"
	"time"
)

func (c *Client) CreatePresignedPutUrl(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error) {
	const Op errors.Op = "minio.CreatePresignedPutUrl"

	url, err := c.objectStore.PresignedPutObject(ctx, bucketName, objectName, time.Duration(expiry))
	if err != nil {
		return "", errors.NewError(Op, errors.Msg("failed to create presigned put url"), err)
	}

	return url.String(), nil
}

func (c *Client) CratedPresignedGetUrl(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error) {
	const Op errors.Op = "minio.CreatePresignedPutUrl"

	url, err := c.objectStore.PresignedPutObject(ctx, bucketName, objectName, time.Duration(expiry))
	if err != nil {
		return "", errors.NewError(Op, errors.Msg("failed to create presigned get url"), err)
	}

	return url.String(), nil
}
