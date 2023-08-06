package main

import (
	"github.com/ArkamFahry/uploadnexus/server/api/routes"
	"github.com/ArkamFahry/uploadnexus/server/config"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/storage/database"
	"github.com/ArkamFahry/uploadnexus/server/storage/objectstore"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

func main() {
	const Op errors.Op = "app.Serve"
	var err error

	config.InitZeroLogger()

	err = envs.InitEnv()
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

	app := fiber.New()

	recoverMiddleware := recover.New()

	app.Use(recoverMiddleware)

	api := app.Group("/api")
	routes.RegisterHealthRoutes(api)
	routes.RegisterBucketRoutes(api)
	routes.RegisterObjectRoutes(api)
	routes.RegisterUploadRoutes(api)

	appPort := envs.EnvStoreInstance.GetEnv().AppPort

	err = app.Listen(":" + appPort)
	if err != nil {
		log.Fatal().Msg(errors.NewError(Op, "error while starting server", err).Error())
	}
}
