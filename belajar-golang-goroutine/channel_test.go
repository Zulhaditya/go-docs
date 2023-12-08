package belajarg_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// jalankan: go test -v -run=TestCreateChannel
func TestCreateChannel(t *testing.T) {
	/*
	   - channel bisa digunakan untuk mengirim dan menerima data
	   - mengirim data gunakan kode: channel <- data
	   - menerima data gunakan kode: data <- channel
	   - jika sudah selesai menggunakan channel, close menggunakan function close()
	*/

	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Zulhaditya Hapiz"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}
