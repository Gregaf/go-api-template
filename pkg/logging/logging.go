package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// Ideally this would have each function defined for logging to hide this global
	// but I will not do that for simplicity.
	Std = NewLogger(logrus.DebugLevel, os.Stdout)
)

// func getLogDir() (string, error) {
// 	ex, err := os.Executable()
// 	if err != nil {
// 		return "", errors.Wrap(err, "failed to get executable path")
// 	}

// 	exPath := filepath.Dir(ex)

// 	return (exPath + "/logs"), nil
// }

type Logger struct {
	*logrus.Logger
}

// TODO: Separate log level from logrus
func NewLogger(level logrus.Level, writers ...io.Writer) *Logger {
	logger := logrus.New()

	multiWriter := io.MultiWriter(writers...)

	logger.SetLevel(level)
	// TODO: Use JSON formatter for production
	// logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(multiWriter)

	if level > logrus.DebugLevel {
		logger.SetReportCaller(true)
	}

	return &Logger{logger}
}
