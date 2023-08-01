package minio

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/envs"
	"github.com/ArkamFahry/uploadnexus/server/rest/exceptions"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	objectStore *minio.Client
}

func NewClient() (*Client, error) {
	const Op exceptions.Op = "minio.NewClient"
	var err error
	var objectStore *minio.Client

	endPoint := envs.EnvStoreInstance.GetEnv().ObjectStoreEndpoint
	accessKey := envs.EnvStoreInstance.GetEnv().ObjectStoreAccessKey
	secretKey := envs.EnvStoreInstance.GetEnv().ObjectStoreSecretKey
	ssl := envs.EnvStoreInstance.GetEnv().ObjectStoreSsl

	objectStore, err = minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: ssl,
	})

	if err != nil {
		return nil, exceptions.NewError(Op, exceptions.Msg("error creating minio client"), err)
	}

	return &Client{
		objectStore: objectStore,
	}, nil
}
