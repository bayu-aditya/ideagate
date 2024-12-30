package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Debug(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warn(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// ErrorP errors with prefix
func ErrorP(prefix string, args ...interface{}) {
	logrus.Error(fmt.Sprintf("[%s]:", prefix), args)
}

func Panic(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}
