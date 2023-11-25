package main

import "fmt"

/*
- nil adalah data kosong
- nil hanya bisa digunakan pada interface, function, map, slice, pointer dan juga channel
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("data map masih kosong")
	} else {
		fmt.Println(data)
	}
}
