package main

import (
	"fmt"
	"math"
)

func main() {
	// pembulatan keatas
	fmt.Println(math.Ceil(1.40))

	// pembulatan kebawah
	fmt.Println(math.Floor(2.90))

	// pembulatan terdekat
	fmt.Println(math.Round(9.80))

	// nilai terbesar
	fmt.Println(math.Max(10, 17))

	// nilai terkecil
	fmt.Println(math.Min(2, 5))
}
