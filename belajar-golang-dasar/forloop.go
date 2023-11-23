package main

import "fmt"

func main() {
	// counter := 1
	//
	// for counter <= 10 {
	// 	fmt.Println("perulangan ke -", counter)
	// 	counter++
	// }

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("perulangan ke -", counter)
	// }

	// for range
	/*
		- for bisa digunakan untuk melakukan iterasi terhadap semua data collection
		- data collection contohnya array, slice, dan map
	*/

	// secara manual
	names := []string{"ackxle", "adit", "inayah"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// penggunaan for range
	for index, name := range names {
		fmt.Println("index", index, ":", name)
	}

	// jika tidak butuh index ganti dengan _
	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println("selesai!")

}
