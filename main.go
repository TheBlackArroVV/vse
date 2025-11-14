package main

import (
	"elastic_go/index"
	"elastic_go/models"
	"fmt"
)

type Query = models.Query

func main() {
	index := index.New("test_index")

	index.Write("first document", "What a weather we have today")
	index.Write("second document", "what beatiful we have today")
	index.Write("third document", "test")

	query := Query{
		Must: []string{"what", "weather", "we"},
	}

	fmt.Println(index.SearchByQuery(query))
}
