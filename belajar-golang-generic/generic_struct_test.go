package belajar_golang_generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[T]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Ackxle",
		Second: "Inayah",
	}

	fmt.Println(data)
	assert.Equal(t, "Zulhaditya", data.ChangeFirst("Zulhaditya"))
	assert.Equal(t, "Hello Syapiq", data.SayHello("Syapiq"))
}
