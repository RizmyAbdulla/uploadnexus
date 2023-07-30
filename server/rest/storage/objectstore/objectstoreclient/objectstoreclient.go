package objectstoreclient

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/entities/objectstoreentities"
)

type ObjectStoreClient interface {
	CreateBucket(ctx context.Context, bucket objectstoreentities.Bucket) error
	DeleteBucket(ctx context.Context, name string) error
	CheckIfBucketExists(ctx context.Context, bucket objectstoreentities.Bucket) (bool, error)

	CreatePresignedPutUrl(ctx context.Context, createdPresignedUrl objectstoreentities.CreatedPresignedUrl) (*objectstoreentities.PresignedUrl, error)
	CratedPresignedGetUrl(ctx context.Context, createdPresignedUrl objectstoreentities.CreatedPresignedUrl) (*objectstoreentities.PresignedUrl, error)

	DeleteObject(ctx context.Context, object objectstoreentities.Object) error
}
