package dependency

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
  Log *logrus.Logger
}

func NewLogger() *Logger {
  logger := logrus.New()
  logger.SetOutput(os.Stdout)
  return &Logger{
    Log : logger,
  }
}
