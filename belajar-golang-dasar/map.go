package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "ackxle"
	// person["address"] = "lingga"

	person := map[string]string{
		"name":    "ackxle",
		"address": "lingga",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "buku golang"
	book["author"] = "ackxle"
	book["ups"] = "salah"

	fmt.Println(book)
	delete(book, "ups")

	fmt.Println(book)
}
