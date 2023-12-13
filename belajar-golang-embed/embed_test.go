package belajargolangembed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed cat.png
var catImage []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("cat_new.png", catImage, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
