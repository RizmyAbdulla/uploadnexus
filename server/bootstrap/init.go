package bootstrap

import (
	"github.com/ArkamFahry/uploadnexus/server/config"
	"github.com/ArkamFahry/uploadnexus/server/envs"
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/ArkamFahry/uploadnexus/server/storage/database"
	"github.com/ArkamFahry/uploadnexus/server/storage/objectstore"
)

func Init() {
	const Op errors.Op = "bootstrap.Init"

	config.InitZeroLogger()

	if err := envs.InitEnv(); err != nil {
		if err := errors.NewError(Op, "error initializing env", err); err != nil {
			return
		}
	}

	if err := database.InitDatabase(); err != nil {
		if err := errors.NewError(Op, "error initializing database", err); err != nil {
			return
		}
	}

	if err := objectstore.InitObjectStore(); err != nil {
		if err := errors.NewError(Op, "error initializing object store", err); err != nil {
			return
		}
	}
}
