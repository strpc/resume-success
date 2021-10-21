package logging

import (
	"log"

	"github.com/sirupsen/logrus"
)

type LogType string

var jsonType LogType = "json"
var plainType LogType = "plain"

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (l *Logger) IsDebug() bool {
	return l.Logger.Level == logrus.DebugLevel
}

func NewLogger(logLevel string, logType LogType) *Logger {
	TimestampFormat := "2006-01-02 15:04:05"
	var formatter logrus.Formatter
	switch logType {
	case jsonType:
		formatter = &logrus.JSONFormatter{TimestampFormat: TimestampFormat}
	case plainType:
		formatter = &logrus.TextFormatter{
			DisableColors:   true,
			FullTimestamp:   true,
			TimestampFormat: TimestampFormat,
		}
	default:
		log.Fatalf("logType %s is not ident. Exit...", logType)
	}
	l, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Log level %s is not parsed. Exit...", logLevel)
	}

	logger := logrus.New()
	logger.SetFormatter(formatter)
	logger.SetLevel(l)
	if l == logrus.DebugLevel {
		logger.SetReportCaller(true)
	}
	e = logrus.NewEntry(logger)
	return &Logger{e}
}
