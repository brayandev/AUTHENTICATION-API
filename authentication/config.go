package authentication

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config env
type Config struct {
	MongoDBEndpoint   string `envconfig:"HOST_MONGODB"`
	MongoDBCollection string `envconfig:"MONGO_DB_COLLECTION"`
	MongoDBName       string `envconfig:"MONGO_DB_NAME"`
}

// NewConfig ...
func NewConfig() *Config {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
