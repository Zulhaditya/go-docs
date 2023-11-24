package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	// cara pertama membuat anonymous function
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("ackxle", blacklist)

	// cara kedua membuat anonymous function
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
