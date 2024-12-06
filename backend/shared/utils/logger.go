package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitializeLogger() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

func LogRequest(method, url string, statusCode int) {
	log.WithFields(logrus.Fields{
		"method":      method,
		"url":         url,
		"status_code": statusCode,
	}).Info("HTTP Request")
}

func LogError(message string, err error) {
	log.WithFields(logrus.Fields{
		"error": err.Error(),
	}).Error(message)
}
