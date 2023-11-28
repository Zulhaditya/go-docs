package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("zulhaditya", "adit"))
	fmt.Println(strings.Split("zulhaditya hapiz", " "))
	fmt.Println(strings.ToLower("zulhaditya hapiz"))
	fmt.Println(strings.ToUpper("zulhaditya hapiz"))
	fmt.Println(strings.Trim("             zulhaditya hapiz         ", " "))
	fmt.Println(strings.ReplaceAll("inayah fitri wulandari", "fitri", "ackxle"))
}
