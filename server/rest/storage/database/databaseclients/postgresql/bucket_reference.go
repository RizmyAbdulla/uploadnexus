package postgresql

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/databaseentities"
)

func (c *DatabaseClient) CreateBucketReference(ctx context.Context, bucketReference databaseentities.BucketReference) error {
	return nil
}

func (c *DatabaseClient) UpdateBucketReference(ctx context.Context, bucketReference databaseentities.BucketReference) error {
	return nil
}

func (c *DatabaseClient) DeleteBucketReference(ctx context.Context, id string) error {
	return nil
}

func (c *DatabaseClient) GetBucketReferenceById(ctx context.Context, id string) (*databaseentities.BucketReference, error) {
	return nil, nil
}

func (c *DatabaseClient) GetBucketReferenceByName(ctx context.Context, name string) (*databaseentities.BucketReference, error) {
	return nil, nil
}

func (c *DatabaseClient) GetBucketReferencesByApplicationId(ctx context.Context, applicationId string) (*[]databaseentities.BucketReference, error) {
	return nil, nil
}

func (c *DatabaseClient) GetBucketReferences(ctx context.Context) (*[]databaseentities.BucketReference, error) {
	return nil, nil
}
