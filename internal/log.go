package internal

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log is used as global logger
var Log logrus.Logger

func init() {
	Log = logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{},
		Level:     logrus.DebugLevel,
	}
}
