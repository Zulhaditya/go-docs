package main

import "fmt"

/*
- type assertions dapat merubah tipe data menjadi tipe data yang diinginkan
- fitur ini digunakan ketika kita bertemu dengan data interface kosong
*/

func random() interface{} {
	return "ok"
}

func main() {
	// result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	// type assertions dengan switch
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown")
	}
}
