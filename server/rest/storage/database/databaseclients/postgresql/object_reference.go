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

func (c *DatabaseClient) GetObjectReferenceById(ctx context.Context, id string) (*databaseentities.ObjectReference, error) {
	return nil, nil
}

func (c *DatabaseClient) GetObjectReferenceByName(ctx context.Context, name string) (*databaseentities.ObjectReference, error) {
	return nil, nil
}

func (c *DatabaseClient) GetObjectReferences(ctx context.Context) (*[]databaseentities.ObjectReference, error) {
	return nil, nil
}
