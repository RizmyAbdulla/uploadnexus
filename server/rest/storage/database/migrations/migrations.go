package migrations

import (
	"embed"
	"github.com/ArkamFahry/uploadnexus/server/rest/envs"
	"github.com/ArkamFahry/uploadnexus/server/rest/errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/rs/zerolog/log"
)

//go:embed postgresql/*.sql
var PostgresqlMigrations embed.FS

const PostgresqlMigrationsFolder = "postgresql"

func MigrateDatabase(migrationFiles embed.FS, migrationFolder string, databaseUrl string) error {
	const Op errors.Op = "migrations.MigrateDatabase"

	if envs.EnvStoreInstance.GetEnv().DatabaseAutoMigrate {
		d, err := iofs.New(migrationFiles, migrationFolder)
		if err != nil {
			return errors.E(Op, errors.Msg("failed to open migration files"), err)
		}

		m, err := migrate.NewWithSourceInstance("iofs", d, databaseUrl)
		if err != nil {
			return err
		}

		if err := m.Up(); err != nil {
			if err.Error() == "no change" {
				log.Info().Msg(errors.E(Op, errors.Msg("no database schema changes"), err).Error())
				return nil
			}
			return err
		}
	}

	log.Info().Msg(errors.E(Op, errors.Msg("database migration completed successfully"), nil).Error())

	return nil
}
