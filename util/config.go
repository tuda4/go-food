package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbDriver string `mapstructure:"DB_DRIVER"`
	DbSource string `mapstructure:"DB_SOURCE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
