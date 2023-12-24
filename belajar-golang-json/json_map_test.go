package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMapDecode(t *testing.T) {
	jsonRequest := `{"id":"12345", "name":"Apple iPhone", "price":2000000}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestJSONMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Asus VivoBook",
		"price": 200000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
