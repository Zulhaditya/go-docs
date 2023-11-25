package main

import "fmt"

/*
  - golang bukan bahasa pemrograman OOP
  - biasanya pada OOP, ada satu data parent di puncak yang bisa dianggap
    sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
  - contoh seperti di java ada java.lang.Object
  - untuk menangani kasus ini, pada golang kita bisa membuat interface kosong
  - interface kosong adalah interface yang tidak memiliki deklarasi method apapun,
    hal ini membuat semua tipe data akan menjadi implementasinya
  - interface kosong juga memiliki alias yaitu any
*/
func Ups() any {
	// return 1
	// return true
	return "ups"
}

// func Ups() interface{} {
// 	// return 1
// 	// return true
// 	return "ups"
// }

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}
