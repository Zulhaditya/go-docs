package main

import "fmt"

/*
	- panic function digunakan untuk menghentikan program
	- panic function biasanya dipanggil ketika terjadi panic pada saat program berjalan
	- saat panic function dipanggil, program akan terhenti, namun defer akan tetap dieksekusi
*/

/*
	- recover adalah function yang bisa digunakan untuk menangkap data panic
	- dengan recover proses panic akan terhenti, sehingga program bisa terus berjalan
*/

func endApp() {
	fmt.Println("end app")
	message := recover()
	fmt.Println("terjadi panic:", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("error")
	}
}

func main() {
	runApp(true)
}
