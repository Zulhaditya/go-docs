package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "put your database host")
	username := flag.String("username", "root", "put your database username")
	password := flag.String("password", "admin123", "put your database password")
	var port *int = flag.Int("port", 3000, "database port")

	flag.Parse()
	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
	fmt.Println("port:", *port)
	fmt.Println("host:", *host)

	// cara menjalankan:
	// go run flag.go -username=ackxle -password=secret -host=fedora -port=5432
}
