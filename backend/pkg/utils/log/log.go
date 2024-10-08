package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

// ErrorP errors with prefix
func ErrorP(prefix string, args ...interface{}) {
	logrus.Error(fmt.Sprintf("[%s]:", prefix), args)
}

func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}
