package index

import (
	"elastic_go/utils"
	"math/rand/v2"
	"strings"
)

type IndexDocument struct {
	words []string
}

type Index struct {
	name            string
	documents       map[int64]IndexDocument
	mappedIndexData MappedIndexData
}

type MappedIndexData struct {
	mappedData map[string]utils.Set[int64]
}

func New(name string) Index {
	return Index{
		name:      name,
		documents: make(map[int64]IndexDocument),
	}
}

func (index *Index) Write(value string) *Index {
	documentId := rand.Int64N(100000)
	transformedString := utils.TransformStrings(value)
	indexedData := IndexDocument{
		words: strings.Split(transformedString, " "),
	}

	if index.mappedIndexData.mappedData == nil {
		index.mappedIndexData.mappedData = make(map[string]utils.Set[int64])
	}

	for word := range strings.SplitSeq(transformedString, " ") {
		index.mappedIndexData.mappedData[word] = index.mappedIndexData.mappedData[word].Add(documentId)

	}

	index.documents[documentId] = indexedData

	return index
}
