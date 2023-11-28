package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("inayah")
	data.PushBack("fitri")
	data.PushBack("wulandari")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // inayah

	next := head.Next()
	fmt.Println(next.Value) // fitri

	next = next.Next()
	fmt.Println(next.Value) // wulandari

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
