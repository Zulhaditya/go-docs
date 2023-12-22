package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJSON(data interface{}) {
	// lakukan konversi menggunakan function marshal
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	logJSON("Ackxle")                        // string
	logJSON(7)                               // number
	logJSON(true)                            // boolean
	logJSON([]string{"Inayah", "Wulandari"}) // slice
}
