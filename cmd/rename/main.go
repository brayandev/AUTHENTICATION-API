package main

import (
	"net/http"
	"time"

	"github.com/student-api/student"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config := student.NewConfig()

	logger, zErr := configLog(zap.NewAtomicLevelAt(config.LogLevel.Value)).Build()
	if zErr != nil {
		panic(zErr)
	}
	defer logger.Sync()

	db, dbErr := student.NewMongoDB(config.MongoDBEndpoint)
	if dbErr != nil {
		logger.Error("failed on database start", zap.NamedError("error", dbErr))
	}

	repository := student.NewRepository(db, config.MongoDBName, config.MongoDBCollection)
	service := student.NewService(repository)

	handler, hErr := createServerHandler(service, logger)
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
