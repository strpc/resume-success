package logging

import (
	"log"

	"github.com/sirupsen/logrus"
)

type LogType string

var jsonType LogType = "json"
var plainType LogType = "plain"

type Logger struct {
	*logrus.Entry
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
	logger.SetReportCaller(true)
	logger.SetFormatter(formatter)
	logger.SetLevel(l)
	e := logrus.NewEntry(logger)
	return &Logger{
		e,
	}
}
