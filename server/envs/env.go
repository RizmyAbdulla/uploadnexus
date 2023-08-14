package envs

import (
	"github.com/ArkamFahry/uploadnexus/server/errors"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func InitEnv() error {
	const Op errors.Op = "envs.InitEnv"
	var err error
	var env Env

	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return errors.NewError(Op, "error reading config file", err)
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&env); err != nil {
		return errors.NewError(Op, "error parsing config file", err)
	}

	if env.AppPort == "" {
		env.AppPort = "8080"
	}

	if env.AppSecret == "" {
		env.AppSecret = uuid.NewString() + uuid.NewString()
	}

	if env.AppIdType == "" {
		env.AppIdType = "uuid"
	}

	if env.DatabaseType == "" {
		env.DatabaseType = "postgresql"
	}

	if env.ObjectStoreType == "" {
		env.ObjectStoreType = "minio"
	}

	if env.ObjectStoreRegion == "" {
		env.ObjectStoreRegion = "us-east-1"
	}

	if env.PresignedPutObjectExpiration == 0 {
		env.PresignedPutObjectExpiration = 3600
	}

	if env.PresignedGetObjectExpiration == 0 {
		env.PresignedGetObjectExpiration = 3600
	}

	if env.CacheType == "" {
		env.CacheType = "redis"
	}

	if env.EventStoreType == "" {
		env.EventStoreType = "redis"
	}

	if env.DatabaseHost == "" {
		return errors.NewError(Op, "error database host is empty", err)
	}

	if env.DatabasePort == "" {
		return errors.NewError(Op, "error database port is empty", err)
	}

	if env.DatabaseUser == "" {
		return errors.NewError(Op, "error database user is empty", err)
	}

	if env.DatabasePassword == "" {
		return errors.NewError(Op, "error database password is empty", err)
	}

	if env.DatabaseName == "" {
		return errors.NewError(Op, "error database name is empty", err)
	}

	if env.ObjectStoreEndpoint == "" {
		return errors.NewError(Op, "error object store endpoint is empty", err)
	}

	if env.BucketName == "" {
		return errors.NewError(Op, "error buckets is empty", err)
	}

	IntiEnvStore(env)

	return nil
}
