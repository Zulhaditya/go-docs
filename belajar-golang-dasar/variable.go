package main

import "fmt"

func main() {
	// deklarasi variable menggunakan keyword var
	var name string

	name = "zulhaditya"
	fmt.Println(name)

	name = "ackxle"
	fmt.Println(name)

	var title = "belajar golang dasar"
	fmt.Println(title)

	title = "belajar basic golang nich!"
	fmt.Println(title)

	// deklarasi variable secara langsung
	data := "ini adalah contoh data"
	fmt.Println(data)

	data = "ini adalah contoh data 2"
	fmt.Println(data)

	// deklarasi multiple variable
	var (
		firstName = "Inayah"
		lastName  = "Wulandari"
		date      = 17
	)

	fmt.Println(firstName, lastName, date)

}
