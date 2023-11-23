package main

import "fmt"

func main() {
	name := "aditya"

	switch name {
	case "ackxle":
		fmt.Println("hello ackxle")
	case "aditya":
		fmt.Println("hello adit")
	case "inayah":
		fmt.Println("hello inayah")
	default:
		fmt.Println("hello unknown")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	panjangNama := len(name)
	switch {
	case panjangNama > 10:
		fmt.Println("nama terlalu panjang")
	case panjangNama > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
