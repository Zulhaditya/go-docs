package main

import "fmt"

// basic function
func sayHello() {
	fmt.Println("hello")
}

/*
  - saat membuat function terkadang kita membutuhkan data dari luar,
    atau kita sebut parameter
  - kita bisa menambahkan parameter di function lebih dari satu
  - parameter tidaklah wajib, kita bisa membuat function tanpa parameter
  - jika kita menambahkan parameter, maka ketika memanggil function tersebut,
    kita wajib memasukkan data parameternya
*/
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("hello", firstName, lastName)
}

/*
  - function bisa mengembalikan data
  - untuk memberitahu bahwa function akan mengembalikan data, kita harus,
    menuliskan tipe data kembalian dari function tersebut
  - jika function tersebut dideklarasikan dengan tipe data pengembalian/return,
    maka wajib functionnya harus mengembalikan data
  - untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return,
    diikuti dengan datanya
*/
func getHello(name string) string {
	hello := "hello " + name
	return hello
}

/*
  - function tidak hanya dapat mengembalikan satu value, tapi bisa juga
    multiple value
  - untuk memberitahu jika function mengembalikan multiple value, kita harus
    menulis semua tipe data return valuenya di function
  - multiple return value wajib ditangkap semua valuenya
  - untuk melakukan ignore return value gunakan tanda _ (garis bawah)
*/
func getFullName() (string, string) {
	return "inayah", "wulandari"
}

/*
  - biasanya saat kita memberitahu bahwa sebuah function mengembalikan value,
    maka kita hanya mendeklarasikan tipe data return value di function
  - namun kita juga bisa membuat variable secara langsung di tipe data return functionnya
*/
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "muhammad"
	middleName = "zulhaditya"
	lastName = "hapiz"

	return firstName, middleName, lastName
}

/*
- parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
- varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja seperti array
- apa perbedaannya dengan array?
  - jika parameter tipe array, kita wajib membuat array terlebih dahulu sebelum
    mengirim ke function
  - jika parameter menggunakan varargs, kita bisa langsung mengirimkan datanya,
    jika lebih dari satu, cukup gunakan tanda koma
  - kita bisa menjadikan slice sebagai vararg parameter
*/
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

/*
- function juga merupakan tipe data dan bisa disimpan didalam variable
*/
func getGoodBye(name string) string {
	return "good bye " + name
}

// function sebagai parameter
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	filteredName := filter(name)
// 	fmt.Println("hello", filteredName)
// }
//
// func spamFilter(name string) string {
// 	if name == "anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

/*
  - jika function terlalu panjang, agak ribet menuliskannya didalam parameter
  - type declaration bisa digunakan untuk membuat alias function, sehingga akan
    mempermudah kita menggunakan function sebagai parameter
*/
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	// panggil semua function
	sayHello()

	sayHelloTo("ackxle", "adit")
	sayHelloTo("inayah", "wulandari")

	result := getHello("syapiq")
	fmt.Println(result)
	fmt.Println(getHello("iqbal"))
	fmt.Println(getHello("aditya"))

	firstName, lastName := getFullName()
	namaAwal, _ := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println(namaAwal)

	namaPertama, namaTengah, namaAkhir := getCompleteName()
	fmt.Println(namaPertama, namaTengah, namaAkhir)

	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))

	goodBye := getGoodBye
	goodBye1 := getGoodBye
	fmt.Println(goodBye("ackxle"))
	fmt.Println(goodBye1("ackxle"))

	sayHelloWithFilter("ackxle", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}
