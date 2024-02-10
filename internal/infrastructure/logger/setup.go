package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

var log *logrus.Logger = nil

func NewLogger() *logrus.Logger {
	return log
}

func SetupLogger(env string) *logrus.Logger {
	if log != nil {
		return log
	}

	switch env {
	case EnvLocal:
		log = setupLocalLog(os.Stdout)
	case EnvDev:
		log = logrus.New()
	case EnvProd:
		log = logrus.New()
	}

	return log
}

func setupLocalLog(out io.Writer) *logrus.Logger {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}
