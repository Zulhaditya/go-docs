package belajarg_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// jalankan: go test -v -run=NamaFunctionTesting
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

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Inayah Wulandari"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(3 * time.Second)
}

// channel hanya untuk mengirim data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Inayah Wulandari"
}

// channel hanya untuk menerima data
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Ackxle"
		channel <- "Zulhaditya"
		channel <- "Inayah"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}
