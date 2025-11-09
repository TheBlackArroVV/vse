package main

import (
	"elastic_go/index"
	"fmt"
)

func main() {
	index := index.New("test_index")

	index.Write("name", "What a weather we have today")
	index.Write("name1", "what beatiful we have today")
	index.Write("name2", "test")

	// fmt.Println(index)
	// fmt.Println(index.Search("What"))

	params := make(map[string][]string)
	params["should"] = []string{"test", "weather", "beatiful", "what"}

	fmt.Println(index.SearchByQuery(params))
}
