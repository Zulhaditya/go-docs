package main

import "fmt"

/*
  - struct adalah sebuah template data yang digunakan untuk menggabungkan,
    nol atau lebih tipe data lainnya dalam satu kesatuan
  - struct biasanya representasi data dalam program aplikasi yang kita buat
  - data di struct disimpan didalam field
  - struct adalah kumpulan dari beberapa field
*/
type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("hello", name, "my name is", customer.Name)
}

func main() {
	var ackxle Customer
	ackxle.Name = "Ackxle"
	ackxle.Address = "Indonesia"
	ackxle.Age = 23
	fmt.Println(ackxle)
	fmt.Println(ackxle.Name)
	fmt.Println(ackxle.Address)
	fmt.Println(ackxle.Age)

	adit := Customer{
		Name:    "Aditya",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(adit)

	inayah := Customer{"Inayah", "Singapore", 27}
	fmt.Println(inayah)

	inayah.sayHello("syapiq")
}
