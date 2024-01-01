package belajar_golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.Equal(t, true, IsSame[string]("ackxle", "ackxle"))
	assert.Equal(t, true, IsSame[int](7, 7))
}
