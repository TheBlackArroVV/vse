package index

import (
	"elastic_go/models"
	"elastic_go/utils"
	"strings"
)

type Index struct {
	name            string
	documents       map[int64]models.IndexDocument
	mappedIndexData MappedIndexData
	currentIdx      int64
}

type MappedIndexData struct {
	values     utils.Set[int64]
	mappedData map[string][]int64
}

type Query struct {
	Should []string
	Must   []string
}

func New(name string) Index {
	return Index{
		name:      name,
		documents: make(map[int64]models.IndexDocument),
	}
}

func (index *Index) Write(value string) *Index {
	documentId := index.currentIdx + 1
	transformedString := utils.TransformStrings(value)
	indexedData := models.IndexDocument{
		Id:    documentId,
		Words: strings.Split(transformedString, " "),
	}

	if index.mappedIndexData.mappedData == nil {
		index.mappedIndexData.mappedData = make(map[string][]int64)
	}

	for word := range strings.SplitSeq(transformedString, " ") {
		index.mappedIndexData.values.Add(documentId)
		index.mappedIndexData.mappedData[word] = append(index.mappedIndexData.mappedData[word], documentId)
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
