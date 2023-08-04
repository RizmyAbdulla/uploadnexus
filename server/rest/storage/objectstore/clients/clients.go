package clients

import (
	"context"
)

type ObjectStoreClient interface {
	CreatePresignedPutUrl(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error)
	CratedPresignedGetUrl(ctx context.Context, bucketName string, objectName string, expiry int64) (string, error)

	DeleteObject(ctx context.Context, bucketName string, objectName string) error
}
