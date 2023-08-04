package postgresql

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/databaseentities"
)

func (c *DatabaseClient) CreateBucket(ctx context.Context, bucketReference databaseentities.Bucket) error {
	return nil
}

func (c *DatabaseClient) UpdateBucket(ctx context.Context, bucketReference databaseentities.Bucket) error {
	return nil
}

func (c *DatabaseClient) DeleteBucket(ctx context.Context, id string) error {
	return nil
}

func (c *DatabaseClient) CheckIfBucketExistsById(ctx context.Context, id string) (bool, error) {
	return false, nil
}

func (c *DatabaseClient) GetBucketById(ctx context.Context, id string) (*databaseentities.Bucket, error) {
	return nil, nil
}

func (c *DatabaseClient) CheckIfBucketExistsByName(ctx context.Context, name string) (bool, error) {
	return false, nil
}

func (c *DatabaseClient) GetBucketByName(ctx context.Context, name string) (*databaseentities.Bucket, error) {
	return nil, nil
}

func (c *DatabaseClient) GetBuckets(ctx context.Context) (*[]databaseentities.Bucket, error) {
	return nil, nil
}
