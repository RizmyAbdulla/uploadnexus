package postgresql

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/databaseentities"
)

func (c *DatabaseClient) CreateObjectReference(ctx context.Context, objectReference databaseentities.ObjectReference) error {
	return nil
}

func (c *DatabaseClient) UpdateObjectReference(ctx context.Context, objectReference databaseentities.ObjectReference) error {
	return nil
}

func (c *DatabaseClient) DeleteObjectReference(ctx context.Context, id string) error {
	return nil
}

func (c *DatabaseClient) CheckIfObjectReferenceById(ctx context.Context, id string) (bool, error) {
	return false, nil
}

func (c *DatabaseClient) GetObjectReferenceById(ctx context.Context, id string) (*databaseentities.ObjectReference, error) {
	return nil, nil
}

func (c *DatabaseClient) CheckIfObjectReferenceByFileKey(ctx context.Context, fileKey string) (bool, error) {
	return false, nil
}

func (c *DatabaseClient) GetObjectReferenceByFileKey(ctx context.Context, fileKey string) (*databaseentities.ObjectReference, error) {
	return nil, nil
}

func (c *DatabaseClient) GetObjectReferencesByBucketId(ctx context.Context, bucketId string) (*[]databaseentities.BucketReference, error) {
	return nil, nil
}

func (c *DatabaseClient) GetObjectReferences(ctx context.Context) (*[]databaseentities.ObjectReference, error) {
	return nil, nil
}
