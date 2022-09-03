package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	GRPC_SERVER_PORT   string `mapstructure:"GRPC_SERVER_PORT"`
	GRPC_CLIENT_DOMAIN string `mapstructure:"GRPC_CLIENT_DOMAIN"`
	GRPC_CLIENT_PORT   string `mapstructure:"GRPC_CLIENT_PORT"`
}

// loadConfig reads configuration from file or environment variables.
func loadConfig(path string) (config Config, err error) {
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

func Cfg() (config Config) {
	config, err := loadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	return
}
