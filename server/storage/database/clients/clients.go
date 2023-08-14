package clients

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/storage/entities"
)

type DatabaseClient interface {
	CreateBucket(ctx context.Context, bucket entities.Bucket) error
	UpdateBucket(ctx context.Context, bucket entities.Bucket) error
	DeleteBucketByID(ctx context.Context, id string) error
	BucketExistsByID(ctx context.Context, id string) (bool, error)
	GetBucketByID(ctx context.Context, id string) (*entities.Bucket, error)
	BucketExistsByName(ctx context.Context, name string) (bool, error)
	GetBucketByName(ctx context.Context, name string) (*entities.Bucket, error)
	ListBuckets(ctx context.Context) (*[]entities.Bucket, error)

	CreateObject(ctx context.Context, object entities.Object) error
	UpdateObject(ctx context.Context, object entities.Object) error
	UpdateObjectBucketName(ctx context.Context, oldBucketName string, newBucketName string) error
	DeleteObjectByID(ctx context.Context, id string) error
	DeleteObjectsByBucketID(ctx context.Context, bucketId string) error
	CheckIfObjectExistsByID(ctx context.Context, id string) (bool, error)
	GetObjectByID(ctx context.Context, id string) (*entities.Object, error)
	ObjectExistsByBucketNameAndObjectName(ctx context.Context, bucketName string, objectName string) (bool, error)
	GetObjectByBucketNameAndObjectName(ctx context.Context, bucketName string, objectName string) (*entities.Object, error)
	ListObjectsByBucketID(ctx context.Context, bucketId string) (*[]entities.Object, error)
	ListObjectsByBucketName(ctx context.Context, bucketName string) (*[]entities.Object, error)
	ListObjects(ctx context.Context) (*[]entities.Object, error)
}
