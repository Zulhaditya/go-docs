package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	/*
		- saat membuat parameter di function, secara default adalah pass by value,
		- data akan di copy lalu dikirimkan ke function tersebut
		- jika kita ingin mengubah data didalam function, data aslinya tidak akan pernah berubah
		- hal ini membuat variable menjadi aman, karena tidak bisa diubah
		- jika kita ingin mengubah data asli di parameter
		- gunakan operator * sebagai parameternya
	*/

	var address Address = Address{}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address)

}
