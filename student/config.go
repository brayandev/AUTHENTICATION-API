package student

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap/zapcore"
)

// Config env
type Config struct {
	MongoDBEndpoint   string         `envconfig:"HOST_MONGODB"`
	MongoDBCollection string         `envconfig:"MONGO_DB_COLLECTION"`
	MongoDBName       string         `envconfig:"MONGO_DB_NAME"`
	LogLevel          LogLevelConfig `envconfig:"SM_LOG_LEVEL" default:"info"`
}

// NewConfig ...
func NewConfig() *Config {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}

// LogLevelConfig log level config.
type LogLevelConfig struct {
	Value zapcore.Level
}
