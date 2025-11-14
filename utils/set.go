package utils

import (
	"elastic_go/models"
	"slices"
)

type Set[T int | int64] struct {
	Values []T
	Length uint64
}

func (set *Set[T]) Add(newValue T) *Set[T] {
	if slices.Contains(set.Values, newValue) {
		return set
	}

	set.Values = append(set.Values, newValue)
	set.Length = set.Length + 1

	return set
}

type IndexDocumentSet struct {
	Values []models.IndexDocument
	Length uint64
}

func (set *IndexDocumentSet) Add(newValue models.IndexDocument) *IndexDocumentSet {
	alreadyPresent := false

	for _, indexDocument := range set.Values {
		if newValue.Id == indexDocument.Id {
			alreadyPresent = true
			break
		}
	}

	if alreadyPresent {
		return set
	}

	set.Values = append(set.Values, newValue)
	set.Length = set.Length + 1

	return set
}
