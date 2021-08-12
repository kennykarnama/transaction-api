package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	RestPort        string `envconfig:"REST_PORT" default:"8080"`
	ServiceName     string `envconfig:"SERVICE_NAME" default:"user_api"`
}

func Get() Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return cfg
}
