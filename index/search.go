package index

import (
	"vse/models"
	. "vse/utils"
)

const MAXIMUM_ALLOWED_DISTANCE = 1

func (index *Index) Search(searchableString string) []models.IndexDocument {
	foundDocumentIds := Set[int64]{}

	for _, foundDocumentId := range index.mappedData[searchableString] {
		foundDocumentIds.Add(index.documents[foundDocumentId].Id)
	}

	for key := range index.mappedData {
		if LevenshteinDistance(searchableString, key) == MAXIMUM_ALLOWED_DISTANCE {
			for _, foundDocumentId := range index.mappedData[key] {
				foundDocumentIds.Add(index.documents[foundDocumentId].Id)
			}
		}
	}

	SortArray(foundDocumentIds.Values, string(models.ASC))

	return index.FindDocumentsByIds(foundDocumentIds.Values)
}

func (index *Index) SearchByQuery(query models.Query) []models.IndexDocument {
	foundDocuments := IndexDocumentSet{}

	foundDocuments.AddMany(index.searchByShould(query.Should))
	foundDocuments.AddMany(index.searchByMust(query.Must))

	if string(query.Order) == "" {
		query.Order = models.ASC
	}
	SortIndexDocument(foundDocuments.Values, string(query.Order))

	return foundDocuments.Values
}

func (index *Index) searchByShould(searchableWords []string) []models.IndexDocument {
	if len(searchableWords) == 0 {
		return []models.IndexDocument{}
	}

	foundDocumentIds := Set[int64]{}

	for _, searchableString := range searchableWords {
		searchableString = TransformStrings(searchableString)
		for _, foundDocumentId := range index.mappedData[searchableString] {
			foundDocumentIds.Add(index.documents[foundDocumentId].Id)
		}

		for key := range index.mappedData {
			if LevenshteinDistance(searchableString, key) == MAXIMUM_ALLOWED_DISTANCE {
				for _, foundDocumentId := range index.mappedData[key] {
					foundDocumentIds.Add(index.documents[foundDocumentId].Id)
				}
			}
		}
	}

	return index.FindDocumentsByIds(foundDocumentIds.Values)
}

func (index *Index) searchByMust(searchableWords []string) []models.IndexDocument {
	if len(searchableWords) == 0 {
		return []models.IndexDocument{}
	}

	foundDocuments := []models.IndexDocument{}
	includedWords := make(map[int64][]int)

	for idx, searchableString := range searchableWords {
		searchableString = TransformStrings(searchableString)
		for _, foundDocumentId := range index.mappedData[searchableString] {
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
