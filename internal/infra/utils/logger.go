package utils

import (
	"os"

	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

var logger *zerolog.Logger

func GetLogger() *zerolog.Logger {
	if logger == nil {
		service := GetStringValueOrDefault(os.Getenv("SV_SERVICE_NAME"), "api")
		newLogger := httplog.NewLogger(service, httplog.Options{
			JSON:            StringToBool(os.Getenv("LOG_JSON"), false),
			LogLevel:        GetStringValueOrDefault(os.Getenv("LOG_LEVEL"), "info"),
			TimeFieldFormat: "2006-01-02T15:04:05Z07:00",
		})
		logger = &newLogger
	}
	return logger
}
