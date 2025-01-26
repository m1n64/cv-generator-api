package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var logger *zap.Logger

func InitLogs() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		"app.log",
		"stdout",
	}

	config.Encoding = "console"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, _ = config.Build()
	defer func() {
		if syncErr := logger.Sync(); syncErr != nil {
			log.Printf("Failed to sync logger: %v", syncErr)
		}
	}()
}

// GetLogger возвращает логгер
func GetLogger() *zap.Logger {
	return logger
}