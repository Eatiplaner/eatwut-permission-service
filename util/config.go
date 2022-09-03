package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	GRPC_SERVER_PORT string `mapstructure:"GRPC_SERVER_PORT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	return config, err
}
