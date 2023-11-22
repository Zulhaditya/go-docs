package main

import "fmt"

func main() {
	// konversi number
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// konversi string
	var name = "zulhaditya"
	var z uint8 = name[0]
	var stringZ = string(z)

	fmt.Println(name)
	fmt.Println(z)
	fmt.Println(stringZ)
}
