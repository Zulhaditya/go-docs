package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/*
		- sebelumnya untuk membuat pointer kita menggunakan operator &
		- function new juga bisa digunakan untuk membuat pointer
		- function new hanya mengembalikan pointer ke data kosong, tidak ada data di awal
	*/

	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.Country = "Indonesia"
	fmt.Println(alamat1)
	fmt.Println(alamat2)

}
