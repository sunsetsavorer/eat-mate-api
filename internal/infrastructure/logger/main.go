package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)

	return &Logger{log}
}
