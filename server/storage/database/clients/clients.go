package clients

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/storage/entities"
)

type DatabaseClient interface {
	// CreateBucket to save bucket information into database
	CreateBucket(ctx context.Context, bucket entities.Bucket) error
	// UpdateBucket to update bucket information in database
	UpdateBucket(ctx context.Context, bucket entities.Bucket) error
	// DeleteBucketById to delete bucket from database
	DeleteBucketById(ctx context.Context, id string) error
	// CheckIfBucketExistsById to check if bucket exists in database using the bucket id
	CheckIfBucketExistsById(ctx context.Context, id string) (bool, error)
	// GetBucketById to get bucket information from database using the bucket id
	GetBucketById(ctx context.Context, id string) (*entities.Bucket, error)
	// CheckIfBucketExistsByName to check if bucket exists in database using the bucket name
	CheckIfBucketExistsByName(ctx context.Context, name string) (bool, error)
	// GetBucketByName to get bucket information from database using the bucket name
	GetBucketByName(ctx context.Context, name string) (*entities.Bucket, error)
	// ListBuckets to get list of buckets from database
	ListBuckets(ctx context.Context) (*[]entities.Bucket, error)

	// CreateObject to save object information into database
	CreateObject(ctx context.Context, object entities.Object) error
	// UpdateObject to update object information in database
	UpdateObject(ctx context.Context, object entities.Object) error
	// UpdateObjectBucketName to update object bucket information in database
	UpdateObjectBucketName(ctx context.Context, oldBucketName string, newBucketName string) error
	// DeleteObjectById to delete object from database
	DeleteObjectById(ctx context.Context, id string) error
	// DeleteObjectsByBucketId to delete object from database using the object id
	DeleteObjectsByBucketId(ctx context.Context, bucketId string) error
	// CheckIfObjectExistsById to check if object exists in database using the object id
	CheckIfObjectExistsById(ctx context.Context, id string) (bool, error)
	// GetObjectById to get object information from database using the object id
	GetObjectById(ctx context.Context, id string) (*entities.Object, error)
	// CheckIfObjectExistsByBucketNameAndObjectName to check if object exists in database using the bucket name and object name
	CheckIfObjectExistsByBucketNameAndObjectName(ctx context.Context, bucketName string, objectName string) (bool, error)
	// GetObjectByBucketNameAndObjectName to get object information from database using the bucket name and object name
	GetObjectByBucketNameAndObjectName(ctx context.Context, bucketName string, objectName string) (*entities.Object, error)
	// ListObjectsByBucketId to get list of objects from database using the bucket id
	ListObjectsByBucketId(ctx context.Context, bucketId string) (*[]entities.Object, error)
	// ListObjectsByBucketName to get list of objects from database using the bucket name
	ListObjectsByBucketName(ctx context.Context, bucketName string) (*[]entities.Object, error)
	// ListObjects to get list of objects from database
	ListObjects(ctx context.Context) (*[]entities.Object, error)
}
