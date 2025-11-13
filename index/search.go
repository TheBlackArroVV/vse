package index

import (
	"elastic_go/utils"
	"strings"
)

func (index *Index) Search(searchableString string) []IndexDocument {
	foundDocuments := []IndexDocument{}
	foundDocumentIds := utils.Set[int64]{}

	for _, foundDocumentId := range index.mappedIndexData.mappedData[searchableString] {
		foundDocumentIds.Add(index.documents[foundDocumentId].id)
	}

	for key := range index.mappedIndexData.mappedData {
		if levenshteinDistance(searchableString, key) == 1 {
			for _, foundDocumentId := range index.mappedIndexData.mappedData[key] {
				foundDocumentIds.Add(index.documents[foundDocumentId].id)
			}
		}
	}

	for _, foundDocumentId := range foundDocumentIds.Values {
		foundDocuments = append(foundDocuments, index.documents[foundDocumentId])
	}

	return foundDocuments
}

func levenshteinDistance(word string, comparable string) int {
	distance := 0

	splitedWord := strings.Split(word, "")
	splitedComparable := strings.Split(comparable, "")

	for idx, letterInWord := range splitedWord {
		if idx >= len(splitedComparable) {
			break
		}

		if letterInWord != splitedComparable[idx] {
			distance = distance + 1
		}
	}

	distance = distance + absInt(len(splitedWord)-len(splitedComparable))

	return distance
}

func absInt(number int) int {
	if number > 0 {
		return number
	}

	return number * -1
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
	foundDocumentIds := utils.Set[int64]{}

	for _, searchableString := range searchableWords {
		for _, foundDocumentId := range index.mappedIndexData.mappedData[searchableString] {
			foundDocumentIds.Add(index.documents[foundDocumentId].id)
		}
	}

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
