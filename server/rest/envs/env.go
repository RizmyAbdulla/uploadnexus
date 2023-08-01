package envs

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/exceptions"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func InitEnv() error {
	const Op exceptions.Op = "envs.InitEnv"
	var err error
	var env Env

	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return exceptions.NewError(Op, "error reading config file", err)
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&env)
	if err != nil {
		return exceptions.NewError(Op, "error parsing config file", err)
	}

	if env.AppPort == "" {
		env.AppPort = "8080"
	}

	if env.AppSecret == "" {
		env.AppSecret = uuid.NewString() + uuid.NewString()
	}

	if env.AppAdminSecret != "" {
		env.AppAdminSecret = uuid.NewString() + uuid.NewString()
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

	if env.CacheType == "" {
		env.CacheType = "redis"
	}

	if env.EventStoreType == "" {
		env.EventStoreType = "redis"
	}

	if env.DatabaseHost == "" {
		return exceptions.NewError(Op, "error database host is empty", err)
	}

	if env.DatabasePort == "" {
		return exceptions.NewError(Op, "error database port is empty", err)
	}

	if env.DatabaseUser == "" {
		return exceptions.NewError(Op, "error database user is empty", err)
	}

	if env.DatabasePassword == "" {
		return exceptions.NewError(Op, "error database password is empty", err)
	}

	if env.DatabaseName == "" {
		return exceptions.NewError(Op, "error database name is empty", err)
	}

	if env.ObjectStoreEndpoint == "" {
		return exceptions.NewError(Op, "error object store endpoint is empty", err)
	}

	if len(env.Buckets) == 0 {
		return exceptions.NewError(Op, "error buckets is empty", err)
	}

	IntiEnvStore(env)

	return nil
}
