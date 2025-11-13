package index

import (
	. "elastic_go/utils"
)

var MAXIMUM_ALLOWED_DISTANCE = 1

func (index *Index) Search(searchableString string) []IndexDocument {
	foundDocuments := []IndexDocument{}
	foundDocumentIds := Set[int64]{}

	for _, foundDocumentId := range index.mappedIndexData.mappedData[searchableString] {
		foundDocumentIds.Add(index.documents[foundDocumentId].id)
	}

	for key := range index.mappedIndexData.mappedData {
		if LevenshteinDistance(searchableString, key) == MAXIMUM_ALLOWED_DISTANCE {
			for _, foundDocumentId := range index.mappedIndexData.mappedData[key] {
				foundDocumentIds.Add(index.documents[foundDocumentId].id)
			}
		}
	}

	SortArray(foundDocumentIds.Values)

	for _, foundDocumentId := range foundDocumentIds.Values {
		foundDocuments = append(foundDocuments, index.documents[foundDocumentId])
	}

	return foundDocuments
}

func (index *Index) SearchByQuery(query Query) []IndexDocument {
	foundDocuments := []IndexDocument{}

	foundDocuments = append(foundDocuments, index.searchByShould(query.Should)...)
	foundDocuments = append(foundDocuments, index.searchByMust(query.Must)...)

	return foundDocuments
}

func (index *Index) searchByShould(searchableWords []string) []IndexDocument {
	if len(searchableWords) == 0 {
		return []IndexDocument{}
	}

	foundDocuments := []IndexDocument{}
	foundDocumentIds := Set[int64]{}

	for _, searchableString := range searchableWords {
		for _, foundDocumentId := range index.mappedIndexData.mappedData[searchableString] {
			foundDocumentIds.Add(index.documents[foundDocumentId].id)
		}

		for key := range index.mappedIndexData.mappedData {
			if LevenshteinDistance(searchableString, key) == MAXIMUM_ALLOWED_DISTANCE {
				for _, foundDocumentId := range index.mappedIndexData.mappedData[key] {
					foundDocumentIds.Add(index.documents[foundDocumentId].id)
				}
			}
		}
	}

	SortArray(foundDocumentIds.Values)

	for _, foundDocumentId := range foundDocumentIds.Values {
		foundDocuments = append(foundDocuments, index.documents[foundDocumentId])
	}

	return foundDocuments
}

func (index *Index) searchByMust(searchableWords []string) []IndexDocument {
	if len(searchableWords) == 0 {
		return []IndexDocument{}
	}

	foundDocuments := []IndexDocument{}
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
