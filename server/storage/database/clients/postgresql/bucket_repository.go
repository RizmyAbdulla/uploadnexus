package postgresql

import (
	"context"
	"fmt"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/storage/entities"
	"github.com/lib/pq"
)

func (c *DatabaseClient) CreateBucket(ctx context.Context, bucket entities.Bucket) error {
	const Op errors.Op = "postgresql.CreateBucket"

	query := fmt.Sprintf(`INSERT INTO %s (id, name, description, allowed_mime_types, allowed_object_size, is_public, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`, entities.BucketCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return errors.NewError(Op, "failed to prepare statement", err)
	}

	_, err = stmt.ExecContext(ctx, bucket.Id, bucket.Name, bucket.Description, pq.Array(bucket.AllowedMimeTypes), bucket.AllowedObjectSize, bucket.IsPublic, bucket.CreatedAt)
	if err != nil {
		return errors.NewError(Op, "failed to create bucket", err)
	}

	return nil
}

func (c *DatabaseClient) UpdateBucket(ctx context.Context, bucket entities.Bucket) error {
	const Op errors.Op = "postgresql.UpdateBucket"

	query := fmt.Sprintf(`UPDATE %s SET name = $1, description = $2, allowed_mime_types = $3, allowed_object_size = $4, is_public = $5, updated_at = $6 WHERE id = $7`, entities.BucketCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return errors.NewError(Op, "failed to prepare statement", err)
	}

	_, err = stmt.ExecContext(ctx, bucket.Name, bucket.Description, pq.Array(bucket.AllowedMimeTypes), bucket.AllowedObjectSize, bucket.IsPublic, bucket.UpdatedAt, bucket.Id)
	if err != nil {
		return errors.NewError(Op, "failed to update bucket", err)
	}

	return nil
}

func (c *DatabaseClient) DeleteBucketByID(ctx context.Context, id string) error {
	const Op errors.Op = "postgresql.DeleteBucketByID"

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, entities.BucketCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return errors.NewError(Op, "failed to prepare statement", err)
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return errors.NewError(Op, "failed to delete bucket", err)
	}

	return nil
}

func (c *DatabaseClient) BucketExistsByID(ctx context.Context, id string) (bool, error) {
	const Op errors.Op = "postgresql.BucketExistsByID"

	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM %s WHERE id = $1)`, entities.BucketCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return false, errors.NewError(Op, "failed to prepare statement", err)
	}

	result := stmt.QueryRowContext(ctx, id)
	err = result.Scan(&exists)
	if err != nil {
		return false, errors.NewError(Op, "failed to check if bucket exists by id", err)
	}
	if exists {
		return true, nil
	}

	return false, nil
}

func (c *DatabaseClient) GetBucketByID(ctx context.Context, id string) (*entities.Bucket, error) {
	const Op errors.Op = "postgresql.GetBucketByID"

	var bucket entities.Bucket
	var allowedMimeTypes pq.StringArray

	query := fmt.Sprintf(`SELECT id, name, description, allowed_mime_types, allowed_object_size, is_public, created_at, updated_at FROM %s WHERE id = $1`, entities.BucketCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.NewError(Op, "failed to prepare statement", err)

	}

	result := stmt.QueryRowContext(ctx, id)
	err = result.Scan(&bucket.Id, &bucket.Name, &bucket.Description, &allowedMimeTypes, &bucket.AllowedObjectSize, &bucket.IsPublic, &bucket.CreatedAt, &bucket.UpdatedAt)
	if err != nil {
		return nil, errors.NewError(Op, "failed to get bucket by id", err)
	}

	bucket.AllowedMimeTypes = (*[]string)(&allowedMimeTypes)

	return &bucket, nil
}

func (c *DatabaseClient) BucketExistsByName(ctx context.Context, name string) (bool, error) {
	const Op errors.Op = "postgresql.BucketExistsByName"

	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM %s WHERE name = $1)`, entities.BucketCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return false, errors.NewError(Op, "failed to prepare statement", err)
	}

	result := stmt.QueryRowContext(ctx, name)
	err = result.Scan(&exists)
	if err != nil {
		return false, errors.NewError(Op, "failed to check if bucket exists by name", err)
	}
	if exists {
		return true, nil
	}

	return false, nil
}

func (c *DatabaseClient) GetBucketByName(ctx context.Context, name string) (*entities.Bucket, error) {
	const Op errors.Op = "postgresql.GetBucketByName"

	var bucket entities.Bucket
	var allowedMimeTypes pq.StringArray

	query := fmt.Sprintf(`SELECT id, name, description, allowed_mime_types, allowed_object_size, is_public, created_at, updated_at FROM %s WHERE name = $1`, entities.BucketCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.NewError(Op, "failed to prepare statement", err)
	}

	result := stmt.QueryRowContext(ctx, name)
	err = result.Scan(&bucket.Id, &bucket.Name, &bucket.Description, &allowedMimeTypes, &bucket.AllowedObjectSize, &bucket.IsPublic, &bucket.CreatedAt, &bucket.UpdatedAt)
	if err != nil {
		return nil, errors.NewError(Op, "failed to get bucket by name", err)
	}

	bucket.AllowedMimeTypes = (*[]string)(&allowedMimeTypes)

	return &bucket, nil
}

func (c *DatabaseClient) ListBuckets(ctx context.Context, limit int, offset int) (*[]entities.Bucket, error) {
	const Op errors.Op = "postgresql.ListBuckets"

	var buckets []entities.Bucket

	query := fmt.Sprintf(`SELECT id, name, description, allowed_mime_types, allowed_object_size, is_public, created_at, updated_at FROM %s LIMIT $1 OFFSET $2`, entities.BucketCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.NewError(Op, "failed to prepare statement", err)
	}

	rows, err := stmt.QueryContext(ctx, limit, offset)
	if err != nil {
		return nil, errors.NewError(Op, "failed to get buckets", err)
	}

	for rows.Next() {
		var bucket entities.Bucket
		var allowedMimeTypes pq.StringArray
		err := rows.Scan(&bucket.Id, &bucket.Name, &bucket.Description, &allowedMimeTypes, &bucket.AllowedObjectSize, &bucket.IsPublic, &bucket.CreatedAt, &bucket.UpdatedAt)
		if err != nil {
			return nil, errors.NewError(Op, "failed to get bucket", err)
		}
		bucket.AllowedMimeTypes = (*[]string)(&allowedMimeTypes)
		buckets = append(buckets, bucket)
	}

	return &buckets, nil
}
