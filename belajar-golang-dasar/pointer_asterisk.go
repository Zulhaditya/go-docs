package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/*
		- untuk mengubah semua variable yang mengacu pada sebuah pointer, gunakan operator *
	*/

	var address1 Address = Address{"Batam", "Kepulauan Riau", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Bandung"
	fmt.Println(address1) // ikut berubah valuenya
	fmt.Println(address2) // value berubah

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
