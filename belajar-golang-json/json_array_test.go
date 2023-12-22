package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestArrayJSON(t *testing.T) {
	customer := Customer{
		FirstName:  "Inayah",
		MiddleName: "Fitri",
		LastName:   "Wulandari",
		Age:        23,
		Married:    true,
		Hobbies: []string{
			"Gaming",
			"Reading",
			"Coding",
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestArrayJSONDecode(t *testing.T) {
	jsonString := `{"FirstName":"Inayah","MiddleName":"Fitri","LastName":"Wulandari","Age":23,"Married":true,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestArrayJSONComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Ackxle",
		Addresses: []Address{
			{
				Street:     "Jalan kenangan",
				Country:    "Indonesia",
				PostalCode: "6772",
			},
			{
				Street:     "Jalan indah",
				Country:    "Singapore",
				PostalCode: "7777",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestArrayJSONComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Ackxle","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan kenangan","Country":"Indonesia","PostalCode":"6772"},{"Street":"Jalan indah","Country":"Singapore","PostalCode":"7777"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyArrayJSONComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan kenangan","Country":"Indonesia","PostalCode":"6772"},{"Street":"Jalan indah","Country":"Singapore","PostalCode":"7777"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
