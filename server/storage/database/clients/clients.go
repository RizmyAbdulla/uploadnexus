package clients

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/storage/entities"
)

type DatabaseClient interface {
	CreateBucket(ctx context.Context, bucket entities.Bucket) error
	UpdateBucket(ctx context.Context, bucket entities.Bucket) error
	DeleteBucket(ctx context.Context, id string) error
	CheckIfBucketExistsById(ctx context.Context, id string) (bool, error)
	GetBucketById(ctx context.Context, id string) (*entities.Bucket, error)
	CheckIfBucketExistsByName(ctx context.Context, name string) (bool, error)
	GetBucketByName(ctx context.Context, name string) (*entities.Bucket, error)
	GetBuckets(ctx context.Context) (*[]entities.Bucket, error)

	CreateObject(ctx context.Context, object entities.Object) error
	UpdateObject(ctx context.Context, object entities.Object) error
	DeleteObject(ctx context.Context, id string) error
	CheckIfObjectExistsById(ctx context.Context, id string) (bool, error)
	GetObjectById(ctx context.Context, id string) (*entities.Object, error)
	CheckIfObjectExistsByName(ctx context.Context, name string) (bool, error)
	GetObjectByName(ctx context.Context, name string) (*entities.Object, error)
	GetObjectsByBucketId(ctx context.Context, bucketId string) (*[]entities.Object, error)
	GetObjectsByBucketName(ctx context.Context, bucketName string) (*[]entities.Object, error)
	GetObjects(ctx context.Context) (*[]entities.Object, error)
}
