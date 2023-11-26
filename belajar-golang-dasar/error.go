package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian dengan nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 0)
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println("hasil:", hasil)
	}
}
