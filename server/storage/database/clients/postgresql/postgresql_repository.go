package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/storage/database/migrations"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DatabaseClient struct {
	db *sql.DB
}

func NewClient() (*DatabaseClient, error) {
	const Op errors.Op = "postgresql.NewClient"
	var err error
	var db *sql.DB

	postgresqlHost := envs.EnvStoreInstance.GetEnv().DatabaseHost
	postgresqlPort := envs.EnvStoreInstance.GetEnv().DatabasePort
	postgresqlUser := envs.EnvStoreInstance.GetEnv().DatabaseUser
	postgresqlPassword := envs.EnvStoreInstance.GetEnv().DatabasePassword
	postgresqlDatabaseName := envs.EnvStoreInstance.GetEnv().DatabaseName
	postgresqlSsl := envs.EnvStoreInstance.GetEnv().DatabaseSsl

	postgresqlUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", postgresqlUser, postgresqlPassword, postgresqlHost, postgresqlPort, postgresqlDatabaseName, postgresqlSsl)

	err = migrations.MigrateDatabase(migrations.PostgresqlMigrations, migrations.PostgresqlMigrationsFolder, postgresqlUrl)
	if err != nil {
		return nil, errors.NewError(Op, "failed to migrate database", err)
	}

	db, err = sql.Open("postgres", postgresqlUrl)
	if err != nil {
		return nil, errors.NewError(Op, "failed to open database connection", err)
	}

	return &DatabaseClient{
		db: db,
	}, nil
}
