package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAckxle NoKTP = "1239393"

	var contoh string = "2222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpAckxle)
	fmt.Println(contohKtp)
}
