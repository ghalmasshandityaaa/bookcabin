package logger

import (
	"bookcabin-backend/config"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(conf *config.Config) *logrus.Logger {
	log := logrus.New()

	log.SetOutput(os.Stdout)
	log.SetReportCaller(false)
	log.SetLevel(logrus.Level(conf.Logger.Level))

	formatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:        "timestamp",
			logrus.FieldKeyLevel:       "level",
			logrus.FieldKeyMsg:         "message",
			logrus.FieldKeyFunc:        "caller",
			logrus.FieldKeyLogrusError: "error",
		},
	}
	log.SetFormatter(formatter)

	return log
}
