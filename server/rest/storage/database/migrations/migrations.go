package migrations

import (
	"embed"
	"github.com/ArkamFahry/uploadnexus/server/rest/envs"
	"github.com/ArkamFahry/uploadnexus/server/rest/exceptions"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/rs/zerolog/log"
)

//go:embed postgresql/*.sql
var PostgresqlMigrations embed.FS

const PostgresqlMigrationsFolder = "postgresql"

func MigrateDatabase(migrationFiles embed.FS, migrationFolder string, databaseUrl string) error {
	const Op exceptions.Op = "migrations.MigrateDatabase"

	if envs.EnvStoreInstance.GetEnv().DatabaseAutoMigrate {
		d, err := iofs.New(migrationFiles, migrationFolder)
		if err != nil {
			return exceptions.NewError(Op, "failed to open migration files", err)
		}

		m, err := migrate.NewWithSourceInstance("iofs", d, databaseUrl)
		if err != nil {
			return err
		}

		if err := m.Up(); err != nil {
			if err.Error() == "no change" {
				log.Info().Msg(exceptions.NewError(Op, "no database schema changes", err).Error())
				return nil
			}
			return err
		}

		log.Info().Msg(exceptions.NewError(Op, "database migration completed successfully", nil).Error())
	}

	return nil
}
