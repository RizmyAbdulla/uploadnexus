package postgresql

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/models"
)

func (c *DatabaseClient) CreateObject(ctx context.Context, objectReference models.Object) error {
	return nil
}

func (c *DatabaseClient) UpdateObject(ctx context.Context, objectReference models.Object) error {
	return nil
}

func (c *DatabaseClient) DeleteObject(ctx context.Context, id string) error {
	return nil
}

func (c *DatabaseClient) CheckIfObjectExistsById(ctx context.Context, id string) (bool, error) {
	return false, nil
}

func (c *DatabaseClient) GetObjectById(ctx context.Context, id string) (*models.Object, error) {
	return nil, nil
}

func (c *DatabaseClient) CheckIfObjectExistsByFileKey(ctx context.Context, fileKey string) (bool, error) {
	return false, nil
}

func (c *DatabaseClient) GetObjectByFileKey(ctx context.Context, fileKey string) (*models.Object, error) {
	return nil, nil
}

func (c *DatabaseClient) GetObjectsByBucketId(ctx context.Context, bucketId string) (*[]models.Bucket, error) {
	return nil, nil
}

func (c *DatabaseClient) GetObjects(ctx context.Context) (*[]models.Object, error) {
	return nil, nil
}
