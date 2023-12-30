package belajar_golang_logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()
	logger.Println("Hello logger!")
	fmt.Println("Hello logger with fmt.")
}

func TestLevel(t *testing.T) {
	// logging level default adalah info ke atas
	// untuk mengubah logging level gunakan logger.SetLevel()
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) // logging dimulai dari trace level

	logger.Trace("Logging in level Trace")
	logger.Debug("Logging in level Debug")
	logger.Info("Logging in level Info")
	logger.Warn("Logging in level Warn")
	logger.Error("Logging in level Error")
}
