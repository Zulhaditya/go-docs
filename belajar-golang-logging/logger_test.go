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
