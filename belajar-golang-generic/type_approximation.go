package belajar_golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type NumberApproximation interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func MinApproximation[T NumberApproximation](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

// testing
func TestMinApproximation(t *testing.T) {
	assert.Equal(t, int(100), MinApproximation[int](100, 200))
	assert.Equal(t, int64(100), MinApproximation[int64](100, 200))
	assert.Equal(t, float64(100.0), MinApproximation[float64](100.0, 200.0))
}
