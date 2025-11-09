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

	foundDocuments = append(foundDocuments, index.searchByShould(query["should"])...)
	foundDocuments = append(foundDocuments, index.searchByMust(query["must"])...)

	return foundDocuments
}

func (index *Index) searchByShould(searchableWords []string) []IndexDocument {
	if len(searchableWords) == 0 {
		return []IndexDocument{}
	}

	foundDocuments := []IndexDocument{}

	for _, searchableString := range searchableWords {
		foundDocuments = append(foundDocuments, index.Search(searchableString)...)
	}

	return foundDocuments
}

func (index *Index) searchByMust(searchableWords []string) []IndexDocument {
	if len(searchableWords) == 0 {
		return []IndexDocument{}
	}

	foundDocuments := []IndexDocument{}
	includedWords := 0

	for _, document := range index.documents {
		includedWords = 0
		for _, searchableWord := range searchableWords {
			for _, word := range document.words {
				if word == searchableWord {
					includedWords = includedWords + 1
				}
			}
		}
		if includedWords == len(searchableWords) {
			foundDocuments = append(foundDocuments, document)
		}
	}

	return foundDocuments
}
