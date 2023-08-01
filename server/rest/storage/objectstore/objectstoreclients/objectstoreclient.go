package objectstoreclients

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/objectstore/objectstoreentities"
)

type ObjectStoreClient interface {
	CreateBucket(ctx context.Context, bucket objectstoreentities.Bucket) error
	DeleteBucket(ctx context.Context, name string) error
	CheckIfBucketExists(ctx context.Context, name string) (bool, error)

	CreatePresignedPutUrl(ctx context.Context, presignedUrl objectstoreentities.PresignedUrl) (string, error)
	CratedPresignedGetUrl(ctx context.Context, presignedUrl objectstoreentities.PresignedUrl) (string, error)

	DeleteObject(ctx context.Context, object objectstoreentities.Object) error
}
