package log

import (
	"log"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func InitLogger(level string, fmt string) {
	fmtLower := strings.ToLower(fmt)
	var formatter logrus.Formatter

	switch fmtLower {
	case "text":
		formatter = &logrus.TextFormatter{}
	default:
		formatter = &logrus.JSONFormatter{}
	}

	logrus.SetFormatter(formatter)
	logrus.SetOutput(os.Stdout)

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatalf("Bad logging level: %s", err)
	}
	logrus.SetLevel(lvl)
}

func Debug(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warning(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}
