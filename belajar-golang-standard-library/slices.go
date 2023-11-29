package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"john", "paul", "george", "ackxle"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "adit"))
	fmt.Println(slices.Index(names, "george"))
}
