package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"inayah", 23},
		{"fitri", 31},
		{"ackxle", 39},
		{"zulhaditya", 22},
		{"hapiz", 28},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
