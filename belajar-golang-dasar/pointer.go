package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Batam", "Kepulauan Riau", "Indonesia"}
	// address2 := address1 // copy value
	//
	// // perubahan ini tidak berdampak ke address1
	// address2.City = "Bandung"
	// fmt.Println(address1) // value tidak berubah
	// fmt.Println(address2) // value berubah

	/*
		- pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama,
		  tanpa menduplikasi data yang sudah ada
		- sederhananya, dengan pointer kita bisa melakukan pass by reference
	*/

	var address1 Address = Address{"Batam", "Kepulauan Riau", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"
	fmt.Println(address1) // ikut berubah valuenya
	fmt.Println(address2) // value berubah
}
