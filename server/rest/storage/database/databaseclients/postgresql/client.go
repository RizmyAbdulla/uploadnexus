package postgresql

import (
	"fmt"
	"github.com/ArkamFahry/uploadnexus/server/rest/envs"
	"github.com/ArkamFahry/uploadnexus/server/rest/exceptions"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/migrations"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseClient struct {
	db *sqlx.DB
}

func NewClient() (*DatabaseClient, error) {
	const Op exceptions.Op = "postgresql.NewClient"
	var err error
	var db *sqlx.DB

	postgresqlHost := envs.EnvStoreInstance.GetEnv().DatabaseHost
	postgresqlPort := envs.EnvStoreInstance.GetEnv().DatabasePort
	postgresqlUser := envs.EnvStoreInstance.GetEnv().DatabaseUser
	postgresqlPassword := envs.EnvStoreInstance.GetEnv().DatabasePassword
	postgresqlDatabaseName := envs.EnvStoreInstance.GetEnv().DatabaseName
	postgresqlSsl := envs.EnvStoreInstance.GetEnv().DatabaseSsl

	postgresqlUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", postgresqlUser, postgresqlPassword, postgresqlHost, postgresqlPort, postgresqlDatabaseName, postgresqlSsl)

	err = migrations.MigrateDatabase(migrations.PostgresqlMigrations, migrations.PostgresqlMigrationsFolder, postgresqlUrl)
	if err != nil {
		return nil, exceptions.NewError(Op, exceptions.Msg("failed to migrate database"), err)
	}

	db, err = sqlx.Open("postgres", postgresqlUrl)
	if err != nil {
		return nil, exceptions.NewError(Op, exceptions.Msg("failed to open database connection"), err)
	}

	return &DatabaseClient{
		db: db,
	}, nil
}
