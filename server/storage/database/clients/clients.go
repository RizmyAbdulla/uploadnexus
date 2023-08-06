package clients

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/models"
)

type DatabaseClient interface {
	CreateBucket(ctx context.Context, bucket models.Bucket) error
	UpdateBucket(ctx context.Context, bucket models.Bucket) error
	DeleteBucket(ctx context.Context, id string) error
	CheckIfBucketExistsById(ctx context.Context, id string) (bool, error)
	GetBucketById(ctx context.Context, id string) (*models.Bucket, error)
	CheckIfBucketExistsByName(ctx context.Context, name string) (bool, error)
	GetBucketByName(ctx context.Context, name string) (*models.Bucket, error)
	GetBuckets(ctx context.Context) (*[]models.Bucket, error)

	CreateObject(ctx context.Context, object models.Object) error
	UpdateObject(ctx context.Context, object models.Object) error
	DeleteObject(ctx context.Context, id string) error
	CheckIfObjectExistsById(ctx context.Context, id string) (bool, error)
	GetObjectById(ctx context.Context, id string) (*models.Object, error)
	CheckIfObjectExistsByName(ctx context.Context, name string) (bool, error)
	GetObjectByName(ctx context.Context, name string) (*models.Object, error)
	GetObjectsByBucketId(ctx context.Context, bucketId string) (*[]models.Object, error)
	GetObjectsByBucketName(ctx context.Context, bucketName string) (*[]models.Object, error)
	GetObjects(ctx context.Context) (*[]models.Object, error)
}
