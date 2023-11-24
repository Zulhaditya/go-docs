package main

import "fmt"

func logging() {
	fmt.Println("dieksekusi paling akhir - 1")
}

func logging1() {
	fmt.Println("dieksekusi paling akhir - 2")
}

func runApplication() {
	defer logging()
	defer logging1()
	fmt.Println("start eksekusi")
}

func main() {
	runApplication()
}
