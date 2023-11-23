package main

import "fmt"

func main() {
	/*
		- break untuk menghentikan seluruh perulangan
		- continue untuk menghentikan perulangan yang sedang berjalan,
		  dan langsung melanjutkan ke perulangan selanjutnya
	*/

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("perulangan ke -", i)
	}

	fmt.Println("")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("perulangan ke -", i)
	}
}
