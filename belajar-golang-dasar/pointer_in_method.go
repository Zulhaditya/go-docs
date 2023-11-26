package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	ackxle := Man{"Ackxle"}
	ackxle.Married()

	fmt.Println(ackxle.Name)
}
