package logger

import (
	"os"
	"sport_helper/internal/config"
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

var instance Logger
var once sync.Once

func GetLogger() Logger {
	once.Do(func() {
		l := logrus.New()
		l.SetOutput(os.Stdout)

		instance = Logger{logrus.NewEntry(l)}
	})

	return instance
}

func Setup(c config.Config) {
	if c.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
