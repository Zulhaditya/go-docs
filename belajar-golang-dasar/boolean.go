package main

import "fmt"

func main() {
	// operasi && => hasil operasi bernilai true jika kedua nilai memiliki nilai true
	// jika salah satu value bernilai false maka hasil operasi bernilai false

	// operasi || => hasil operasi bernilai false jika kedua nilai memiliki nilai false
	// jika salah satu value bernilai true maka hasil operasi bernilai true

	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)

}
