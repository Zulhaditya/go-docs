package helper

import "fmt"

var version = "1.0.0"
var Application = "Golang"

func sayGoodBye(name string) string {
	return "Good bye +" + name
}

func Contoh() {
	sayGoodBye("ackxle")
	fmt.Println(version)
}

// package harus sama dengan nama foldernya
func SayHello(name string) string {
	return "hello " + name
}
