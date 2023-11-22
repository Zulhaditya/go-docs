package main

import "fmt"

func main() {

	// perbandingan string
	var name1 = "ackxle"
	var name2 = "ackxle"

	var resultString bool = name1 == name2
	fmt.Println(resultString)

	// perbandingan int
	var angka1 = 11
	var angka2 = 10

	var resultInt1 bool = angka1 > angka2
	var resultInt2 bool = angka1 >= angka2
	var resultInt3 bool = angka1 < angka2
	var resultInt4 bool = angka1 <= angka2
	var resultInt5 bool = angka1 != angka2

	fmt.Println(resultInt1) // true
	fmt.Println(resultInt2) // true
	fmt.Println(resultInt3) // false
	fmt.Println(resultInt4) // false
	fmt.Println(resultInt5) // true

}
