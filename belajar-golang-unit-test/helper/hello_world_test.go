package helper

import "testing"

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
		panic("Result is not 'Hello Ackxle'")
	}
}

func TestHelloWorldInayah(t *testing.T) {
	result := HelloWorld("Inayah")
	if result != "Hello Inayah" {
		// unit test gagal
		panic("Result is not 'Hello Inayah'")
	}
}
