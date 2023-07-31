package database

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/constants"
	"github.com/ArkamFahry/uploadnexus/server/rest/envs"
	"github.com/ArkamFahry/uploadnexus/server/rest/errors"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/databaseclients"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/databaseclients/postgresql"
)

var Client databaseclients.DatabaseClient

func InitDatabase() error {
	const Op errors.Op = "database.InitDatabase"
	var err error

	isPostgreSQL := envs.EnvStoreInstance.GetEnv().DatabaseType == constants.DatabaseTypePostgreSQL

	if isPostgreSQL {
		Client, err = postgresql.NewClient()
		if err != nil {
			return errors.NewError(Op, errors.Msg("failed to create database provider"), err)
		}
	}

	return nil
}

func GetDatabaseClient() databaseclients.DatabaseClient {
	return Client
}
