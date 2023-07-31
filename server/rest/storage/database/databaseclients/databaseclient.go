package databaseclients

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/databaseentities"
)

type DatabaseClient interface {
	CreateApplication(ctx context.Context, application databaseentities.Application) error
	UpdateApplication(ctx context.Context, application databaseentities.Application) error
	DeleteApplication(ctx context.Context, id string) error
	GetApplicationById(ctx context.Context, id string) (*databaseentities.Application, error)
	GetApplicationByName(ctx context.Context, name string) (*databaseentities.Application, error)
	GetApplications(ctx context.Context) (*[]databaseentities.Application, error)

	CreateBucketReference(ctx context.Context, bucketReference databaseentities.BucketReference) error
	UpdateBucketReference(ctx context.Context, bucketReference databaseentities.BucketReference) error
	DeleteBucketReference(ctx context.Context, id string) error
	GetBucketReferenceById(ctx context.Context, id string) (*databaseentities.BucketReference, error)
	GetBucketReferenceByName(ctx context.Context, name string) (*databaseentities.BucketReference, error)
	GetBucketReferences(ctx context.Context) (*[]databaseentities.BucketReference, error)

	CreateObjectReference(ctx context.Context, objectReference databaseentities.ObjectReference) error
	UpdateObjectReference(ctx context.Context, objectReference databaseentities.ObjectReference) error
	DeleteObjectReference(ctx context.Context, id string) error
	GetObjectReferenceById(ctx context.Context, id string) (*databaseentities.ObjectReference, error)
	GetObjectReferenceByName(ctx context.Context, name string) (*databaseentities.ObjectReference, error)
	GetObjectReferences(ctx context.Context) (*[]databaseentities.ObjectReference, error)
}
