package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJSONStreamEncoder(t *testing.T) {
	writer, _ := os.Create("sample_encoder.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Muhammad",
		MiddleName: "Zulhaditya",
		LastName:   "Hapiz",
	}

	_ = encoder.Encode(customer)
	fmt.Println(customer)
}
