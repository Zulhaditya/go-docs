package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("hello", value.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	/*
				- interface adalah tipe data abstract, dan tidak memiliki implementasi secara langsung
				- sebuah interface berisikan definisi-definisi method
				- biasanya interface digunakan sebagai kontrak
		    - gunakan struct untuk menginisialisasi interface
	*/
	person := Person{Name: "ackxle"}
	SayHello(person)

	animal := Animal{Name: "kucing"}
	SayHello(animal)

}
