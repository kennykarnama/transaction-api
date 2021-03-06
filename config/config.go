package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	RestPort           string `envconfig:"REST_PORT" default:"8080"`
	ServiceName        string `envconfig:"SERVICE_NAME" default:"transaction_api"`
	UserApiHttpBaseURL string `envconfig:"USER_API_HTTP_BASE_URL" default:"http://localhost:8080"`
}

func Get() Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return cfg
}
