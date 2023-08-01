package databaseclients

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/databaseentities"
)

type DatabaseClient interface {
	CreateBucketReference(ctx context.Context, bucketReference databaseentities.BucketReference) error
	UpdateBucketReference(ctx context.Context, bucketReference databaseentities.BucketReference) error
	DeleteBucketReference(ctx context.Context, id string) error
	CheckIfBucketReferenceById(ctx context.Context, id string) (bool, error)
	GetBucketReferenceById(ctx context.Context, id string) (*databaseentities.BucketReference, error)
	CheckIfBucketReferenceByName(ctx context.Context, name string) (bool, error)
	GetBucketReferenceByName(ctx context.Context, name string) (*databaseentities.BucketReference, error)
	GetBucketReferencesByApplicationId(ctx context.Context, applicationId string) (*[]databaseentities.BucketReference, error)
	GetBucketReferences(ctx context.Context) (*[]databaseentities.BucketReference, error)

	CreateObjectReference(ctx context.Context, objectReference databaseentities.ObjectReference) error
	UpdateObjectReference(ctx context.Context, objectReference databaseentities.ObjectReference) error
	DeleteObjectReference(ctx context.Context, id string) error
	CheckIfObjectReferenceById(ctx context.Context, id string) (bool, error)
	GetObjectReferenceById(ctx context.Context, id string) (*databaseentities.ObjectReference, error)
	CheckIfObjectReferenceByFileKey(ctx context.Context, fileKey string) (bool, error)
	GetObjectReferenceByFileKey(ctx context.Context, fileKey string) (*databaseentities.ObjectReference, error)
	GetObjectReferencesByBucketId(ctx context.Context, bucketId string) (*[]databaseentities.BucketReference, error)
	GetObjectReferences(ctx context.Context) (*[]databaseentities.ObjectReference, error)
}
