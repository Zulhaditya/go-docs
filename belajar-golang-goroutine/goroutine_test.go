package belajarg_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World!")
}

// jalankan: go test -v -run=TestCreateGoroutine
func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups...")

	time.Sleep(1 * time.Second)
}
