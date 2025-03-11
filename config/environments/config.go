package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	DatabaseUrl  string `json:"databaseUrl"`
	DatabaseName string `json:"databaseName"`
	Environment  string `json:"environment"`
}

type Iconfig interface {
	GetConfig() config
}

func NewConfig() Iconfig {
	environment := os.Getenv("GO_ENV")

	if environment == "" {
		environment = "development"
	}

	configPath := filepath.Join("config", "environments", environment)

	viper.SetConfigName("backend-config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)
	err := viper.MergeInConfig()

	if err != nil {
		panic(err)
	}

	viper.SetConfigName("service-endpoints")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)

	if err := viper.MergeInConfig(); err != nil {
		panic(err)
	}

	c := config{}

	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}

	c.Environment = environment
	return c
}

func (c config) GetConfig() config {
	return c
}