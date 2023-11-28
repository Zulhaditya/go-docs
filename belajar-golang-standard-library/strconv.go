package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println(boolean)
	}

	resultInt, err := strconv.Atoi("10000")
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	var stringInt string = strconv.Itoa(9999)
	fmt.Println(stringInt)
}
