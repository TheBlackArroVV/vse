package index

import "strings"

type IndexDocument struct {
	name  string
	words []string
}

type Index struct {
	name      string
	documents []IndexDocument
}

func New(name string) Index {
	return Index{
		name:      name,
		documents: []IndexDocument{},
	}
}

func (index *Index) Write(name, value string) *Index {
	indexedData := IndexDocument{
		name:  name,
		words: strings.Split(strings.ToLower(value), " "),
	}

	index.documents = append(index.documents, indexedData)

	return index
}
