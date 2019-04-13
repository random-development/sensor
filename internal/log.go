package internal

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log logrus.Logger

func init() {
	log = logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{},
		Level:     logrus.DebugLevel,
	}
}
