package envs

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/errors"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func InitEnv() error {
	const Op errors.Op = "envs.InitEnv"
	var err error
	var env Env

	viper.AutomaticEnv()
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return errors.NewError(Op, errors.Msg("error reading config file"), err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		return errors.NewError(Op, errors.Msg("error parsing config file"), err)
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
