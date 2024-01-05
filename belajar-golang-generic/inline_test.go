package belajar_golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin[int](100, 200))
	assert.Equal(t, int64(200), FindMin[int64](int64(200), int64(300)))
	assert.Equal(t, 100.0, FindMin(100.0, 200.0))
}

func GetFirst[T []E, E any](data T) E {
	first := data[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Ackxle", "Inayah", "Wulandari",
	}

	first := GetFirst[[]string, string](names)
	assert.Equal(t, "Ackxle", first)
}
