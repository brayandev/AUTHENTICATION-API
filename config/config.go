package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config env
type Config struct {
	MongoDBEndpoint string `envconfig:"HOST_MONGODB"`
}

// NewConfig ...
func NewConfig() *Config {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
