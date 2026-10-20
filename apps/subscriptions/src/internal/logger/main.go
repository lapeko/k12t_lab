package logger

import (
	"os"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/utils"
	"github.com/sirupsen/logrus"
)

func New(env utils.APP_ENV) *logrus.Logger {
	log := logrus.New()

	log.SetOutput(os.Stdout)

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
		},
	})

	if env != utils.APP_ENV_DEV {
		log.SetFormatter(&logrus.JSONFormatter{})
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetFormatter(&logrus.TextFormatter{ForceColors: true})
		log.SetLevel(logrus.DebugLevel)
	}

	return log
}
