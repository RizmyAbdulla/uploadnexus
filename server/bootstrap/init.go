package bootstrap

import (
	"github.com/ArkamFahry/uploadnexus/server/config"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/storage/database"
	"github.com/ArkamFahry/uploadnexus/server/storage/objectstore"
	"github.com/rs/zerolog/log"
)

func Init() {
	const Op errors.Op = "bootstrap.Init"

	config.InitZeroLogger()

	err := envs.InitEnv()
	if err != nil {
		log.Fatal().Msg(errors.NewError(Op, "error initializing env", err).Error())
	}

	err = database.InitDatabase()
	if err != nil {
		log.Fatal().Msg(errors.NewError(Op, "error initializing database", err).Error())
	}

	err = objectstore.InitObjectStore()
	if err != nil {
		log.Fatal().Msg(errors.NewError(Op, "error initializing object store", err).Error())
	}
}
