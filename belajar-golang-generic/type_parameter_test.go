package belajar_golang_generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// T adalah simbol generic
// gunakan interface{} untuk menerima semua tipe data
// return value T
func Length[T any](param T) T {
	// fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	assert.True(t, true)
}

func TestLength(t *testing.T) {
	// [string] tentukan tipenya string
	var result string = Length[string]("Ackxle")
	fmt.Println(result)

	var resultNumber int = Length[int](100)
	fmt.Println(resultNumber)

	var resultBoolean bool = Length[bool](false)
	fmt.Println(resultBoolean)
}
