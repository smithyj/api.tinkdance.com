package config

import (
	"github.com/spf13/viper"
	"tinkdance/pkg/database"
	"tinkdance/pkg/env"
	"tinkdance/pkg/redis"
)

type Config struct {
	Debug    bool            `yaml:"debug"`
	Redis    redis.Config    `yaml:"redis"`
	Database database.Config `yaml:"database"`
}

func Setup() *Config {
	var config = new(Config)
	viper.SetConfigName(env.Active().Value())
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	return config
}
