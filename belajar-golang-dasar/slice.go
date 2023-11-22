package main

import "fmt"

func main() {
	/*
		- tipe data slice adalah potongan dari data array
		- slice mirip dengan array, tetapi ukuran slice bisa berubah
		- slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian
		- atau seluruh data didalam array
	*/

	names := [...]string{"zulhaditya", "hapiz", "ackxle", "inayah", "fitri", "wulandari"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	var slice3 []string = names[3:]
	fmt.Println(slice3)

	var slice4 []string = names[:]
	fmt.Println(slice4)

	// function di slice
	/*
		- len(slice) => mendapatkan panjang slice bukan panjang array-nya
		- cap(slice) => mendapatkan kapasitas
		- append(slice, data) => membuat slice baru dengan menambahkan data ke posisi terakhir slice,
		  jika kapasitas sudah penuh maka akan membuat array baru
		- make([]TypeData, length, capacity) => membuat slice baru
		- copy(destination, source) => menyalin slice dari source ke destination
	*/

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:] // sabtu, minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru")
	daySlice2[0] = "sabtu lama"
	// 	daysBaru := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu", "libur baru"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// membuat slice dari awal
	newSlice := make([]string, 2, 5)
	newSlice[0] = "adit"
	newSlice[1] = "ackxle"
	// newSlice[2] = "inayah" // error harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "inayah")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "wulandari"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// cara copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// perbedaan slice dan array
	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
