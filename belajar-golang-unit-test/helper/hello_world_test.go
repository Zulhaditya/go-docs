package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
menjalankan unit test:
- go test
- go test -v
- go test -v
- go test -v -run=TestHelloWorldInayah => untuk testing fungsi tertentu
- go test -v ./... => running unit test di semua package (root folder)
*/

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ackxle")
	if result != "Hello Ackxle" {
		// unit test gagal
		t.Error("Result must be 'Hello Ackxle'")
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldInayah(t *testing.T) {
	result := HelloWorld("Inayah")
	if result != "Hello Inayah" {
		// unit test gagal
		t.Fatal("Result must be 'Hello Inayah'")
	}

	fmt.Println("TestHelloWorldInayah done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ackxle")
	assert.Equal(t, "Hello Ackxle", result, "Result must be 'Hello Ackxle'")
	fmt.Println("TestHelloWorldAssert with assert done!")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ackxle")
	require.Equal(t, "Hello Ackxle", result, "Result must be 'Hello Ackxle'")
	fmt.Println("TestHelloWorldAssert with require done!")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("cannot run on mac os")
	}

	result := HelloWorld("Ackxle")
	require.Equal(t, "Hello Ackxle", result, "Result must be 'Hello Ackxle'")

}
