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

	err := envs.InitEnv()
	if err != nil {
		err := errors.NewError(Op, "error initializing env", err)
		if err != nil {
			return
		}
	}

	err = database.InitDatabase()
	if err != nil {
		err := errors.NewError(Op, "error initializing database", err)
		if err != nil {
			return
		}
	}

	err = objectstore.InitObjectStore()
	if err != nil {
		err := errors.NewError(Op, "error initializing object store", err)
		if err != nil {
			return
		}
	}
}
