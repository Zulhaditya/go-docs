package main

import "fmt"

func main() {
	/*
		- array berisikan kumpulan data dengan tipe yang sama
		- saat membuat array kita harus menentukan jumlah data yang bisa ditampung array
		- daya tampung array tidak bisa bertambah
		- index di array dimulai dari 0
	*/

	var names [3]string
	names[0] = "Inayah"
	names[1] = "Fitri"
	names[2] = "Wulandari"

	fmt.Println(names[0], names[1], names[2])
	fmt.Println(names)

	// deklarasi array secara langsung
	var values = [3]int{70, 80, 90}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// function di array
	/*
		- len(array) => mendapatkan panjang array
		- array[index] => mendapatkan data di posisi index
		- array[index] = value => mengubah data di posisi index
	*/

	var nilai = [...]int{
		10,
		20,
		30,
	}

	fmt.Println(nilai)
	fmt.Println(len(nilai))
	nilai[0] = 100
	fmt.Println(nilai)
}
