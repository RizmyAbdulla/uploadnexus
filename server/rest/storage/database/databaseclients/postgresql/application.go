package postgresql

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/errors"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/databaseentities"
)

func (c *DatabaseClient) CreateApplication(ctx context.Context, application databaseentities.Application) error {
	const Op errors.Op = "postgresql.CreateApplication"

	query := "INSERT INTO applications (id, name, description, created_at, updated_at) VALUES (:id, :name, :description, :created_at)"

	_, err := c.db.NamedExecContext(ctx, query, application)
	if err != nil {
		return errors.NewError(Op, "failed to create application", err)
	}

	return nil
}

func (c *DatabaseClient) UpdateApplication(ctx context.Context, application databaseentities.Application) error {
	const Op errors.Op = "postgresql.UpdateApplication"

	query := "UPDATE applications SET name = :name, description = :description, updated_at = :updated_at WHERE id = :id"

	_, err := c.db.NamedExecContext(ctx, query, application)
	if err != nil {
		return errors.NewError(Op, "failed to update application", err)
	}

	return nil
}

func (c *DatabaseClient) DeleteApplication(ctx context.Context, id string) error {
	const Op errors.Op = "postgresql.DeleteApplication"

	query := "DELETE FROM applications WHERE id = $id"

	_, err := c.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.NewError(Op, "failed to delete application", err)
	}

	return nil
}

func (c *DatabaseClient) GetApplicationById(ctx context.Context, id string) (*databaseentities.Application, error) {
	const Op errors.Op = "postgresql.GetApplicationById"
	var result databaseentities.Application

	query := "SELECT id, name, description, created_at, updated_at FROM applications WHERE id = $id"

	err := c.db.GetContext(ctx, &result, query, id)

	if err != nil {
		return nil, errors.NewError(Op, "failed to get application by id", err)
	}
	return &result, nil
}

func (c *DatabaseClient) GetApplicationByName(ctx context.Context, name string) (*databaseentities.Application, error) {
	const Op errors.Op = "postgresql.GetApplicationByName"
	var result databaseentities.Application

	query := "SELECT id, name, description, created_at, updated_at FROM applications WHERE name = $name"

	err := c.db.GetContext(ctx, &result, query, name)
	if err != nil {
		return nil, errors.NewError(Op, "failed to get application by name", err)
	}

	return &result, nil
}

func (c *DatabaseClient) GetApplications(ctx context.Context) (*[]databaseentities.Application, error) {
	const Op errors.Op = "postgresql.GetApplications"
	var result []databaseentities.Application

	query := "SELECT id, name, description, created_at, updated_at FROM applications"

	err := c.db.SelectContext(ctx, &result, query)
	if err != nil {
		return nil, errors.NewError(Op, "failed to get applications", err)
	}

	return &result, nil
}
