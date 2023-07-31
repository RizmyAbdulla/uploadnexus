package main

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/api/handlers"
	"github.com/ArkamFahry/uploadnexus/server/rest/config"
	"github.com/ArkamFahry/uploadnexus/server/rest/envs"
	"github.com/ArkamFahry/uploadnexus/server/rest/errors"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/objectstore"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	const Op errors.Op = "app.Serve"
	var err error

	config.InitZeroLogger()

	err = envs.InitEnv()
	if err != nil {
		log.Fatal().Msg(errors.NewError(Op, errors.Msg("error initializing env"), err).Error())
	}

	err = database.InitDatabase()
	if err != nil {
		log.Fatal().Msg(errors.NewError(Op, errors.Msg("error initializing database"), err).Error())
	}

	err = objectstore.InitObjectStore()
	if err != nil {
		log.Fatal().Msg(errors.NewError(Op, errors.Msg("error initializing object store"), err).Error())
	}

	app := fiber.New()

	api := app.Group("/api")
	handlers.RegisterHealthRoutes(api)

	appPort := envs.EnvStoreInstance.GetEnv().AppPort

	err = app.Listen(":" + appPort)
	if err != nil {
		log.Fatal().Msg(errors.NewError(Op, errors.Msg("error while starting server"), err).Error())
	}
}
