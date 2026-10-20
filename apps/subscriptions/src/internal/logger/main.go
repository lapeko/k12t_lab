package logger

import (
	"os"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/utils"
	"github.com/sirupsen/logrus"
)

func New(env utils.APP_ENV) *logrus.Logger {
	log := logrus.New()

	log.SetOutput(os.Stdout)

	if env != "prod" {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
		})
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetFormatter(&logrus.JSONFormatter{})
		log.SetLevel(logrus.InfoLevel)
	}

	return log
}
