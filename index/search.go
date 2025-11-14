package index

import (
	"elastic_go/models"
	. "elastic_go/utils"
)

const MAXIMUM_ALLOWED_DISTANCE = 1

func (index *Index) Search(searchableString string) []models.IndexDocument {
	foundDocumentIds := Set[int64]{}

	for _, foundDocumentId := range index.mappedIndexData.mappedData[searchableString] {
		foundDocumentIds.Add(index.documents[foundDocumentId].Id)
	}

	for key := range index.mappedIndexData.mappedData {
		if LevenshteinDistance(searchableString, key) == MAXIMUM_ALLOWED_DISTANCE {
			for _, foundDocumentId := range index.mappedIndexData.mappedData[key] {
				foundDocumentIds.Add(index.documents[foundDocumentId].Id)
			}
		}
	}

	SortArray(foundDocumentIds.Values)

	return index.FindDocumentsByIds(foundDocumentIds.Values)
}

func (index *Index) SearchByQuery(query Query) []models.IndexDocument {
	foundDocuments := IndexDocumentSet{}

	for _, foundDocument := range index.searchByShould(query.Should) {
		foundDocuments.Add(foundDocument)
	}
	for _, foundDocument := range index.searchByMust(query.Must) {
		foundDocuments.Add(foundDocument)
	}

	return foundDocuments.Values
}

func (index *Index) searchByShould(searchableWords []string) []models.IndexDocument {
	if len(searchableWords) == 0 {
		return []models.IndexDocument{}
	}

	foundDocumentIds := Set[int64]{}

	for _, searchableString := range searchableWords {
		for _, foundDocumentId := range index.mappedIndexData.mappedData[searchableString] {
			foundDocumentIds.Add(index.documents[foundDocumentId].Id)
		}

		for key := range index.mappedIndexData.mappedData {
			if LevenshteinDistance(searchableString, key) == MAXIMUM_ALLOWED_DISTANCE {
				for _, foundDocumentId := range index.mappedIndexData.mappedData[key] {
					foundDocumentIds.Add(index.documents[foundDocumentId].Id)
				}
			}
		}
	}

	SortArray(foundDocumentIds.Values)

	return index.FindDocumentsByIds(foundDocumentIds.Values)
}

func (index *Index) searchByMust(searchableWords []string) []models.IndexDocument {
	if len(searchableWords) == 0 {
		return []models.IndexDocument{}
	}

	foundDocuments := []models.IndexDocument{}
	includedWords := make(map[int64][]int)

	for idx, searchableString := range searchableWords {
		for _, foundDocumentId := range index.mappedIndexData.mappedData[searchableString] {
			includedWords[foundDocumentId] = append(includedWords[foundDocumentId], idx)
		}
	}

	for foundDocumentId, includedWord := range includedWords {
		if len(includedWord) == len(searchableWords) {
			foundDocuments = append(foundDocuments, index.documents[foundDocumentId])
		}
	}

	return foundDocuments
}
