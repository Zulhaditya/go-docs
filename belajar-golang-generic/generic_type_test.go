package belajar_golang_generic

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	// buat slice dengan tipe data string
	names := Bag[string]{"Ackxle", "Inayah", "Zulhaditya"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	// buat slice dengan tipe data integer
	number := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(number)
}
