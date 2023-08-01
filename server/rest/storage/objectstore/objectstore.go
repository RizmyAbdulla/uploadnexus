package objectstore

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/constants"
	"github.com/ArkamFahry/uploadnexus/server/rest/envs"
	"github.com/ArkamFahry/uploadnexus/server/rest/exceptions"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/objectstore/objectstoreclients"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/objectstore/objectstoreclients/minio"
)

var Client objectstoreclients.ObjectStoreClient

func InitObjectStore() error {
	const Op exceptions.Op = "objectstore.InitObjectStore"
	var err error
	isMinio := envs.EnvStoreInstance.GetEnv().ObjectStoreType == constants.ObjectStoreTypeMinio

	if isMinio {
		Client, err = minio.NewClient()
		if err != nil {
			return exceptions.NewError(Op, exceptions.Msg("failed to create object store provider"), err)
		}
	}

	return nil
}

func GetObjectStoreClient() objectstoreclients.ObjectStoreClient {
	return Client
}
