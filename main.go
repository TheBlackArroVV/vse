package main

import (
	"elastic_go/index"
	"fmt"
)

type Query = index.Query

func main() {
	index := index.New("test_index")

	index.Write("What a weather we have today")
	index.Write("what beatiful we have today")
	index.Write("test")

	query := Query{
		Must: []string{"what", "weather", "we"},
	}

	fmt.Println(index.SearchByQuery(query))
}
