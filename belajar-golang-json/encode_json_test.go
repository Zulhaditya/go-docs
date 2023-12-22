package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJSON(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	logJSON("Ackxle")
	logJSON(7)
	logJSON(true)
	logJSON([]string{"Inayah", "Wulandari"})
}
