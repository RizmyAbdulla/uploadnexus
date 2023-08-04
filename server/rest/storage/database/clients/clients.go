package databaseclients

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/databaseentities"
)

type DatabaseClient interface {
	CreateBucket(ctx context.Context, bucketReference databaseentities.Bucket) error
	UpdateBucket(ctx context.Context, bucketReference databaseentities.Bucket) error
	DeleteBucket(ctx context.Context, id string) error
	CheckIfBucketExistsById(ctx context.Context, id string) (bool, error)
	GetBucketById(ctx context.Context, id string) (*databaseentities.Bucket, error)
	CheckIfBucketExistsByName(ctx context.Context, name string) (bool, error)
	GetBucketByName(ctx context.Context, name string) (*databaseentities.Bucket, error)
	GetBuckets(ctx context.Context) (*[]databaseentities.Bucket, error)

	CreateObject(ctx context.Context, objectReference databaseentities.Object) error
	UpdateObject(ctx context.Context, objectReference databaseentities.Object) error
	DeleteObject(ctx context.Context, id string) error
	CheckIfObjectExistsById(ctx context.Context, id string) (bool, error)
	GetObjectById(ctx context.Context, id string) (*databaseentities.Object, error)
	CheckIfObjectExistsByFileKey(ctx context.Context, fileKey string) (bool, error)
	GetObjectByFileKey(ctx context.Context, fileKey string) (*databaseentities.Object, error)
	GetObjectsByBucketId(ctx context.Context, bucketId string) (*[]databaseentities.Bucket, error)
	GetObjects(ctx context.Context) (*[]databaseentities.Object, error)
}
