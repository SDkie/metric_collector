package logger

import (
	"os"

	"github.com/Sirupsen/logrus"
)

var log *logrus.Logger

func Init() {
	log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)

	switch os.Getenv("MODE") {
	case "debug":
		log.Level = logrus.DebugLevel
	default:
		log.Level = logrus.DebugLevel
	}

	Info("Logger Successfully Initialize")
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Err(args ...interface{}) {
	log.Error(args...)
}

func Errf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func PanicfIfError(err error, format string, args ...interface{}) {
	if err != nil {
		Panicf(format, args...)
	}
}
