package main

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/api/routes"
	"github.com/ArkamFahry/uploadnexus/server/rest/config"
	"github.com/ArkamFahry/uploadnexus/server/rest/envs"
	"github.com/ArkamFahry/uploadnexus/server/rest/exceptions"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/objectstore"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

func main() {
	const Op exceptions.Op = "app.Serve"
	var err error

	config.InitZeroLogger()

	err = envs.InitEnv()
	if err != nil {
		log.Fatal().Msg(exceptions.NewError(Op, "error initializing env", err).Error())
	}

	err = database.InitDatabase()
	if err != nil {
		log.Fatal().Msg(exceptions.NewError(Op, "error initializing database", err).Error())
	}

	err = objectstore.InitObjectStore()
	if err != nil {
		log.Fatal().Msg(exceptions.NewError(Op, "error initializing object store", err).Error())
	}

	app := fiber.New()

	recoverMiddleware := recover.New()

	app.Use(recoverMiddleware)

	api := app.Group("/api")
	routes.RegisterHealthRoutes(api)

	appPort := envs.EnvStoreInstance.GetEnv().AppPort

	err = app.Listen(":" + appPort)
	if err != nil {
		log.Fatal().Msg(exceptions.NewError(Op, "error while starting server", err).Error())
	}
}
