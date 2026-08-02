package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}

func newSugaredLogger() *zap.SugaredLogger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, _ := cfg.Build()
	defer logger.Sync()
	sugar := logger.Sugar()
	return sugar
}

var Instance = newSugaredLogger()
