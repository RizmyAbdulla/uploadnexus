package migrations

import "embed"

//go:embed postgresql/*.sql
var PostgresqlMigrations embed.FS

const PostgresqlMigrationsFolder = "postgresql"
