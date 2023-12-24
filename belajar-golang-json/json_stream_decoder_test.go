package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJSONStreamDecoder(t *testing.T) {
	reader, _ := os.Open("sample.json")
	decoder := json.NewDecoder(reader)

	customer := Customer{}
	_ = decoder.Decode(&customer)
	fmt.Println(customer)
}
