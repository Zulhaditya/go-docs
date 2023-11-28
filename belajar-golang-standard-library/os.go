package main

import (
	"fmt"
	"os"
)

func main() {
	// penggunaan argument
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	// dapatkan hostname
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
