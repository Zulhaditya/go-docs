package main

import "fmt"

func main() {
	name := "ackxle"

	if name == "ackxle" {
		fmt.Println("hello ackxle")
	} else if name == "adit" {
		fmt.Println("hello adit")
	} else if name == "inayah" {
		fmt.Println("hello inayah")
	} else {
		fmt.Println("hello unknown")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
