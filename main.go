package main

import (
	"elastic_go/index"
	"fmt"
)

func main() {
	index := index.New("test_index")

	index.Write("name", "What a beatiful weather we have today")
	index.Write("name1", "what beatiful we have today")
	index.Write("name", "test")

	fmt.Println(index)

	fmt.Println(index.Search("What"))
}
