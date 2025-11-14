package index

import (
	"strings"
	"vse/models"
	"vse/utils"
)

type Index struct {
	name       string
	documents  map[int64]models.IndexDocument
	currentIdx int64
	mappedData map[string][]int64
}

func New(name string) Index {
	return Index{
		name:       name,
		documents:  make(map[int64]models.IndexDocument),
		mappedData: make(map[string][]int64),
		currentIdx: 0,
	}
}

func (index *Index) Write(name, value string) *Index {
	documentId := index.currentIdx + 1
	transformedString := utils.TransformStrings(value)
	indexedData := models.IndexDocument{
		Id:    documentId,
		Name:  name,
		Words: strings.Split(transformedString, " "),
	}

	for word := range strings.SplitSeq(transformedString, " ") {
		index.mappedData[word] = append(index.mappedData[word], documentId)
	}
	index.documents[documentId] = indexedData
	index.currentIdx = index.currentIdx + 1

	return index
}

func (index *Index) FindDocumentsByIds(ids []int64) []models.IndexDocument {
	foundDocuments := []models.IndexDocument{}

	for _, foundDocumentId := range ids {
		foundDocuments = append(foundDocuments, index.documents[foundDocumentId])
	}

	return foundDocuments
}
