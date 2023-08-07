package bootstrap

import (
	"github.com/ArkamFahry/uploadnexus/server/api/routes"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

func Serve() {
	const Op errors.Op = "bootstrap.Serve"

	app := fiber.New()

	recoverMiddleware := recover.New()

	app.Use(recoverMiddleware)

	api := app.Group("/api")
	routes.RegisterHealthRoutes(api)
	routes.RegisterBucketRoutes(api)
	routes.RegisterObjectRoutes(api)
	routes.RegisterUploadRoutes(api)

	appPort := envs.EnvStoreInstance.GetEnv().AppPort

	err := app.Listen(":" + appPort)
	if err != nil {
		log.Fatal().Msg(errors.NewError(Op, "error while starting server", err).Error())
	}
}
