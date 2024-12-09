package utils

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

// InitLogger initializes a production logger
func InitLogger() {
	var err error
	Logger, err = zap.NewProduction() // Use NewDevelopment() for local development
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
}

// Info logs an informational message
func Info(message string, fields ...zap.Field) {
	Logger.Info(message, fields...)
}

// Error logs an error message
func Error(message string, fields ...zap.Field) {
	Logger.Error(message, fields...)
}

// Debug logs a debug message
func Debug(message string, fields ...zap.Field) {
	Logger.Debug(message, fields...)
}

// Warn logs a warning message
func Warn(message string, fields ...zap.Field) {
	Logger.Warn(message, fields...)
}
