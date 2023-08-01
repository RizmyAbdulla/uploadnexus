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

	IntiEnvStore(env)

	return nil
}
