package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("hello info")
	logrus.Warn("hello warn")

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("hello error")
	logrus.Warn("hello debug")
}
