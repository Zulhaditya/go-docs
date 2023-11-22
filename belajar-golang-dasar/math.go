package main

import "fmt"

func main() {
	// normal assignments
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d*e

	// augmented assignments
	var i = 10
	i += 10
	i += 5

	// unary operator
	var x = 1
	x++ // x = x + 1
	x++ // x = x + 1
	x-- // x = x - 1

	fmt.Println(c)
	fmt.Println(i)
	fmt.Println(x)
}
