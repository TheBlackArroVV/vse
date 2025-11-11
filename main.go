package main

import (
	"elastic_go/index"
	"fmt"
)

func main() {
	index := index.New("test_index")

	index.Write("What a weather we have today")
	index.Write("what beatiful we have today")
	index.Write("test")

	params := make(map[string][]string)
	params["must"] = []string{"what", "weather", "we"}

	fmt.Println(index.SearchByQuery(params))
}
