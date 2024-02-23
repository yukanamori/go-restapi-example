package config

import (
	"fmt"
	"myapp/pkg/appenv"

	"github.com/kelseyhightower/envconfig"
)

// Config はアプリケーションの設定を表します。
type Config struct {
	// AppEnv はアプリケーションの環境を表します。
	AppEnv appenv.AppEnv `envconfig:"APP_ENV" required:"true"`

	// Server はアプリケーションのサーバー設定を表します。
	Server struct {
		Port string `envconfig:"APP_SERVER_PORT" required:"true"`
	}

	// DB はアプリケーションのDB設定を表します。
	DB struct {
		Host     string `envconfig:"APP_DB_HOST" required:"true"`
		Port     string `envconfig:"APP_DB_PORT" required:"true"`
		User     string `envconfig:"APP_DB_USER" required:"true"`
		Password string `envconfig:"APP_DB_PASSWORD" required:"true"`
		Name     string `envconfig:"APP_DB_NAME" required:"true"`
		Params   string `envconfig:"APP_DB_PARAMS" required:"true"`
	}

	// Log はアプリケーションのログ設定を表します。
	Log struct {
		Debug bool `envconfig:"APP_LOG_DEBUG" required:"true"`
	}
}

var config *Config

// GetConfig は設定を返します。
func GetConfig() *Config {
	return config
}

func init() {
	config = &Config{}
	if err := envconfig.Process("myapp", config); err != nil {
		panic(fmt.Errorf("error loading config from env: %s", err))
	}
}
