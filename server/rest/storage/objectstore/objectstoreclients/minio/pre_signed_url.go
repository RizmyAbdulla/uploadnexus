package minio

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/errors"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/objectstore/objectstoreentities"
)

func (c *Client) CreatePresignedPutUrl(ctx context.Context, createdPresignedUrl objectstoreentities.CreatedPresignedUrl) (string, error) {
	const Op errors.Op = "minio.CreatePresignedPutUrl"

	url, err := c.objectStore.PresignedPutObject(ctx, createdPresignedUrl.BucketName, createdPresignedUrl.ObjectName, createdPresignedUrl.GetExpiry())
	if err != nil {
		return "", errors.NewError(Op, errors.Msg("failed to create presigned put url"), err)
	}

	return url.String(), nil
}

func (c *Client) CratedPresignedGetUrl(ctx context.Context, createdPresignedUrl objectstoreentities.CreatedPresignedUrl) (string, error) {
	const Op errors.Op = "minio.CreatePresignedPutUrl"

	url, err := c.objectStore.PresignedPutObject(ctx, createdPresignedUrl.BucketName, createdPresignedUrl.ObjectName, createdPresignedUrl.GetExpiry())
	if err != nil {
		return "", errors.NewError(Op, errors.Msg("failed to create presigned get url"), err)
	}

	return url.String(), nil
}
