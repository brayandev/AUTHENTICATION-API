package main

import (
	"net/http"
	"time"

	"github.com/authentication-api/authentication"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config := authentication.NewConfig()

	logger, zErr := configLog(zap.NewAtomicLevelAt(config.LogLevel.Value)).Build()
	if zErr != nil {
		panic(zErr)
	}
	defer logger.Sync()

	db, dbErr := authentication.NewMongoDB(config.MongoDBEndpoint)
	if dbErr != nil {
		logger.Error("failed on database start", zap.NamedError("error", dbErr))
	}
	repository := authentication.NewRepository(db)
	service := authentication.NewService(repository)

	handler, hErr := createServerHandler(service, logger, db)
	if hErr != nil {
		logger.Error("failed on server start", zap.NamedError("error", hErr))
	}
	http.ListenAndServe(":8080", handler)
}

func configLog(level zap.AtomicLevel) zap.Config {
	return zap.Config{
		Level:         level,
		Development:   false,
		DisableCaller: true,
		Sampling:      nil,
		Encoding:      "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: millisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func millisDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt(int(float64(d) / float64(time.Millisecond)))
}
