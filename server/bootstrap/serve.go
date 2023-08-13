package bootstrap

import (
	"github.com/ArkamFahry/uploadnexus/server/api/handlers"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Serve() {
	const Op errors.Op = "bootstrap.Serve"

	app := fiber.New()

	recoverMiddleware := recover.New()

	app.Use(recoverMiddleware)

	api := app.Group("/api")
	handlers.RegisterHealthRoutes(api)
	handlers.RegisterBucketRoutes(api)
	handlers.RegisterObjectRoutes(api)

	appPort := envs.EnvStoreInstance.GetEnv().AppPort

	err := app.Listen(":*" + appPort)
	if err != nil {
		err := errors.NewError(Op, "error while starting server", err)
		if err != nil {
			return
		}
	}
}
