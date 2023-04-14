package logging

import (
	over "github.com/Trendyol/overlog"
	"github.com/rs/zerolog"
	"os"
)

func init() {
	//Setup logger
	logFile, _ := os.Create(os.Getenv("APP_HOME") + "/logs/spoc-service.log")
	fileLogger := zerolog.New(logFile).With().Logger()
	over.New(fileLogger)
	over.SetGlobalFields([]string{"customerId", "domainId", "requestId"})
}

func Info(format string, args ...interface{}) {
	over.Log().Infof(format, args...)
}

func Debug(format string, args ...interface{}) {
	over.Log().Debugf(format, args...)
}

func Warn(format string, args ...interface{}) {
	over.Log().Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	over.Log().Errorf(format, args...)
}
