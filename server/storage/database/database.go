package database

import (
	"github.com/ArkamFahry/uploadnexus/server/constants"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/storage/database/clients"
	"github.com/ArkamFahry/uploadnexus/server/storage/database/clients/postgresql"
)

var Client clients.DatabaseClient

func InitDatabase() error {
	const Op errors.Op = "database.InitDatabase"
	var err error

	isPostgreSQL := envs.EnvStoreInstance.GetEnv().DatabaseType == constants.DatabaseTypePostgreSQL

	if isPostgreSQL {
		Client, err = postgresql.NewClient()
		if err != nil {
			return errors.NewError(Op, "failed to create database provider", err)
		}
	}

	return nil
}

func GetDatabaseClient() clients.DatabaseClient {
	return Client
}
