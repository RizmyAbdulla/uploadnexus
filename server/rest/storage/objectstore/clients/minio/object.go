package minio

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/errors"
	"github.com/minio/minio-go/v7"
)

func (c *Client) DeleteObject(ctx context.Context, bucketName string, objectName string) error {
	const Op errors.Op = "minio.DeleteObject"

	err := c.objectStore.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return errors.NewError(Op, errors.Msg("failed to delete object"), err)
	}

	return nil
}
