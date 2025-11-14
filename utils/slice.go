package utils

import (
	"elastic_go/models"
	"slices"
)

func SortArray(arr []int64, order string) {
	lower := 0
	high := 0

	if order == "ASC" {
		lower = -1
		high = 1
	}

	if order == "DESC" {
		lower = 1
		high = -1
	}

	slices.SortFunc(arr, func(a, b int64) int {
		if a < b {
			return lower
		}
		if a > b {
			return high
		}
		return 0
	})
}

func SortIndexDocument(arr []models.IndexDocument, order string) {
	lower := 0
	high := 0

	if order == "ASC" {
		lower = -1
		high = 1
	}

	if order == "DESC" {
		lower = 1
		high = -1
	}

	slices.SortFunc(arr, func(a, b models.IndexDocument) int {
		if a.Id < b.Id {
			return lower
		}
		if a.Id > b.Id {
			return high
		}
		return 0
	})

}
