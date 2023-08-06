package postgresql

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/storage/entities"
)

func (c *DatabaseClient) CreateObject(ctx context.Context, object entities.Object) error {
	const Op errors.Op = "postgresql.CreateObject"

	query := fmt.Sprintf(`INSERT INTO %s (id, bucket, name, mime_type, size, upload_status, metadata, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, entities.ObjectCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return errors.NewError(Op, "failed to prepare statement", err)
	}

	metadataJson, err := json.Marshal(object.Metadata)
	if err != nil {
		return errors.NewError(Op, "failed to marshal metadata", err)
	}

	_, err = stmt.ExecContext(ctx, object.Id, object.Bucket, object.Name, object.MimeType, object.Size, object.UploadStatus, metadataJson, object.CreatedAt)

	return nil
}

func (c *DatabaseClient) UpdateObject(ctx context.Context, object entities.Object) error {
	const Op errors.Op = "postgresql.UpdateObject"

	query := fmt.Sprintf(`UPDATE %s SET bucket = $1, name = $2, mime_type = $3, size = $4, upload_status = $5, metadata = $6, updated_at = $7 WHERE id = $8`, entities.ObjectCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return errors.NewError(Op, "failed to prepare statement", err)
	}

	metadataJson, err := json.Marshal(object.Metadata)
	if err != nil {
		return errors.NewError(Op, "failed to marshal metadata", err)
	}

	_, err = stmt.ExecContext(ctx, object.Bucket, object.Name, object.MimeType, object.Size, object.UploadStatus, metadataJson, object.UpdatedAt, object.Id)

	return nil
}

func (c *DatabaseClient) DeleteObject(ctx context.Context, id string) error {
	const Op errors.Op = "postgresql.DeleteObject"

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, entities.ObjectCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return errors.NewError(Op, "failed to prepare statement", err)
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return errors.NewError(Op, "failed to delete object", err)
	}

	return nil
}

func (c *DatabaseClient) CheckIfObjectExistsById(ctx context.Context, id string) (bool, error) {
	const Op errors.Op = "postgresql.CheckIfObjectExistsById"
	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM %s WHERE id = $1)`, entities.ObjectCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return false, errors.NewError(Op, "failed to prepare statement", err)
	}

	result := stmt.QueryRowContext(ctx, id)
	err = result.Scan(&exists)
	if err != nil {
		return false, errors.NewError(Op, "failed to check if object exists by id", err)
	}
	if exists {
		return true, nil
	}

	return false, nil
}

func (c *DatabaseClient) GetObjectById(ctx context.Context, id string) (*entities.Object, error) {
	const Op errors.Op = "postgresql.GetObjectById"
	var object entities.Object
	var metadataJson json.RawMessage

	query := fmt.Sprintf(`SELECT id, bucket, name, mime_type, size, upload_status, metadata, created_at, updated_at FROM %s WHERE id = $1`, entities.ObjectCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.NewError(Op, "failed to prepare statement", err)
	}

	result := stmt.QueryRowContext(ctx, id)
	err = result.Scan(&object.Id, &object.Bucket, &object.Name, &object.MimeType, &object.Size, &object.UploadStatus, &metadataJson, &object.CreatedAt, &object.UpdatedAt)
	if err != nil {
		return nil, errors.NewError(Op, "failed to get object by id", err)
	}

	err = json.Unmarshal(metadataJson, &object.Metadata)
	if err != nil {
		return nil, errors.NewError(Op, "failed to unmarshal metadata", err)
	}

	return &object, nil
}

func (c *DatabaseClient) CheckIfObjectExistsByName(ctx context.Context, name string) (bool, error) {
	const Op errors.Op = "postgresql.CheckIfObjectExistsByName"
	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM %s WHERE name = $1)`, entities.ObjectCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return false, errors.NewError(Op, "failed to prepare statement", err)
	}

	result := stmt.QueryRowContext(ctx, name)
	err = result.Scan(&exists)
	if err != nil {
		return false, errors.NewError(Op, "failed to check if object exists by name", err)
	}
	if exists {
		return true, nil
	}

	return false, nil
}

func (c *DatabaseClient) GetObjectByName(ctx context.Context, name string) (*entities.Object, error) {
	const Op errors.Op = "postgresql.GetObjectByName"
	var object entities.Object
	var metadataJson json.RawMessage

	query := fmt.Sprintf(`SELECT id, bucket, name, mime_type, size, upload_status, metadata, created_at, updated_at FROM %s WHERE name = $1`, entities.ObjectCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.NewError(Op, "failed to prepare statement", err)
	}

	result := stmt.QueryRowContext(ctx, name)
	err = result.Scan(&object.Id, &object.Bucket, &object.Name, &object.MimeType, &object.Size, &object.UploadStatus, &metadataJson, &object.CreatedAt, &object.UpdatedAt)
	if err != nil {
		return nil, errors.NewError(Op, "failed to get object by name", err)
	}

	err = json.Unmarshal(metadataJson, &object.Metadata)
	if err != nil {
		return nil, errors.NewError(Op, "failed to unmarshal metadata", err)
	}

	return nil, nil
}

func (c *DatabaseClient) GetObjectsByBucketId(ctx context.Context, bucketId string) (*[]entities.Object, error) {
	const Op errors.Op = "postgresql.GetObjectsByBucketId"
	var objects []entities.Object

	query := fmt.Sprintf(`SELECT id, bucket, name, mime_type, size, upload_status, metadata, created_at, updated_at FROM %s WHERE bucket = $1`, entities.ObjectCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.NewError(Op, "failed to prepare statement", err)
	}

	result, err := stmt.QueryContext(ctx, bucketId)
	if err != nil {
		return nil, errors.NewError(Op, "failed to get objects by bucket id", err)
	}

	for result.Next() {
		var object entities.Object
		var metadataJson json.RawMessage
		err = result.Scan(&object.Id, &object.Bucket, &object.Name, &object.MimeType, &object.Size, &object.UploadStatus, &metadataJson, &object.CreatedAt, &object.UpdatedAt)
		if err != nil {
			return nil, errors.NewError(Op, "failed to get object by id", err)
		}
		err = json.Unmarshal(metadataJson, &object.Metadata)
		if err != nil {
			return nil, errors.NewError(Op, "failed to unmarshal metadata", err)
		}
		objects = append(objects, object)
	}

	return &objects, nil
}

func (c *DatabaseClient) GetObjectsByBucketName(ctx context.Context, bucketName string) (*[]entities.Object, error) {
	const Op errors.Op = "postgresql.GetObjectsByBucketName"
	var objects []entities.Object

	query := fmt.Sprintf(`SELECT id, bucket, name, mime_type, size, upload_status, metadata, created_at, updated_at FROM %s WHERE bucket = (SELECT id FROM %s WHERE name = $1)`, entities.ObjectCollection, entities.BucketCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.NewError(Op, "failed to prepare statement", err)
	}

	result, err := stmt.QueryContext(ctx, bucketName)
	if err != nil {
		return nil, errors.NewError(Op, "failed to get objects by bucket name", err)
	}

	for result.Next() {
		var object entities.Object
		var metadataJson json.RawMessage
		err = result.Scan(&object.Id, &object.Bucket, &object.Name, &object.MimeType, &object.Size, &object.UploadStatus, &metadataJson, &object.CreatedAt, &object.UpdatedAt)
		if err != nil {
			return nil, errors.NewError(Op, "failed to get object by id", err)
		}
		err = json.Unmarshal(metadataJson, &object.Metadata)
		if err != nil {
			return nil, errors.NewError(Op, "failed to unmarshal metadata", err)
		}
		objects = append(objects, object)
	}

	return nil, nil
}

func (c *DatabaseClient) GetObjects(ctx context.Context) (*[]entities.Object, error) {
	const Op errors.Op = "postgresql.GetObjects"
	var objects []entities.Object

	query := fmt.Sprintf(`SELECT id, bucket, name, mime_type, size, upload_status, metadata, created_at, updated_at FROM %s`, entities.ObjectCollection)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.NewError(Op, "failed to prepare statement", err)
	}

	result, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, errors.NewError(Op, "failed to get objects", err)
	}

	for result.Next() {
		var object entities.Object
		var metadataJson json.RawMessage
		err = result.Scan(&object.Id, &object.Bucket, &object.Name, &object.MimeType, &object.Size, &object.UploadStatus, &metadataJson, &object.CreatedAt, &object.UpdatedAt)
		if err != nil {
			return nil, errors.NewError(Op, "failed to get object by id", err)
		}
		err = json.Unmarshal(metadataJson, &object.Metadata)
		if err != nil {
			return nil, errors.NewError(Op, "failed to unmarshal metadata", err)
		}
		objects = append(objects, object)
	}

	return &objects, nil
}
