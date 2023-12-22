package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	// buat data json
	jsonRequest :=
		`{
      "FirstName":"Inayah",
      "MiddleName":"Fitri",
      "LastName":"Wulandari"
     }`

	// konversi data json ke byte untuk dikirim ke parameter unmarshal
	jsonBytes := []byte(jsonRequest)

	// buat data customer berdasarkan struct
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
