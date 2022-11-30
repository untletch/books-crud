package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBURI string `mapstructure:"DB_URI"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigName("debug")
	viper.SetConfigType("env")
	viper.AddConfigPath("./pkg/common/config/envs")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, errors.Unwrap(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, errors.Unwrap(err)
	}

	return config, nil
}
