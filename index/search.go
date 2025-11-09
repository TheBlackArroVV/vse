package index

import (
	"strings"
)

func (index *Index) Search(searchableString string) []IndexDocument {
	foundDocuments := []IndexDocument{}

	for _, document := range index.documents {
		for _, word := range document.words {
			if word == strings.ToLower(searchableString) {
				foundDocuments = append(foundDocuments, document)
			}
		}
	}

	return foundDocuments
}

func (index *Index) SearchByQuery(query map[string][]string) []IndexDocument {
	foundDocuments := []IndexDocument{}

	searchableStrings := query["should"]
	for _, searchableString := range searchableStrings {
		foundDocuments = append(foundDocuments, index.Search(searchableString)...)
	}
	return foundDocuments
}
