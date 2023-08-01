package minio

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/exceptions"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/objectstore/objectstoreentities"
	"github.com/minio/minio-go/v7"
)

func (c *Client) CreateBucket(ctx context.Context, bucket objectstoreentities.Bucket) error {
	const Op exceptions.Op = "minio.CreateBucket"

	err := c.objectStore.MakeBucket(ctx, bucket.Name, minio.MakeBucketOptions{})
	if err != nil {
		return exceptions.NewError(Op, exceptions.Msg("failed to create bucket"), err)
	}

	return nil
}

func (c *Client) DeleteBucket(ctx context.Context, name string) error {
	const Op exceptions.Op = "minio.DeleteBucket"

	err := c.objectStore.RemoveBucket(ctx, name)
	if err != nil {
		return exceptions.NewError(Op, exceptions.Msg("failed to delete bucket"), err)
	}

	return nil
}

func (c *Client) CheckIfBucketExists(ctx context.Context, name string) (bool, error) {
	const Op exceptions.Op = "minio.CheckIfBucketExists"

	exists, err := c.objectStore.BucketExists(ctx, name)
	if err != nil {
		return false, exceptions.NewError(Op, exceptions.Msg("failed to check if bucket exists"), err)
	}

	return exists, nil
}
