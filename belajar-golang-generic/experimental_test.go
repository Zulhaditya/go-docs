package belajar_golang_generic

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func ExperimentalMin[T constraints.Ordered](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 200.0))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Ackxle",
	}

	second := map[string]string{
		"Name": "Ackxle",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlices(t *testing.T) {
	first := []string{"Ackxle"}
	second := []string{"Ackxle"}
	assert.True(t, slices.Equal(first, second))
}
