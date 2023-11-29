package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"hello", "world", "heyyyoooo!"})
	_ = writer.Write([]string{"can", "i", "help u?"})

	writer.Flush()
}
