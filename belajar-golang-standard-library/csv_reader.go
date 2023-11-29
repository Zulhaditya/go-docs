package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "muhammad,zulhaditya,hapiz\n" +
		"muhammad,syapiq,alfarazi\n" +
		"muhammad,iqbal,ramadhan\n" +
		"inayah,fitri,wulandari\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
