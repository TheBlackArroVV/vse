package utils_test

import (
	"elastic_go/utils"
	"testing"
)

func TestArraySort(t *testing.T) {
	arr := []int64{5, 3, 4, 1, 2}
	expectedASC := []int64{1, 2, 3, 4, 5}
	expectedDESC := []int64{5, 4, 3, 2, 1}

	utils.SortArray(arr, "ASC")

	for idx, element := range arr {
		if element != expectedASC[idx] {
			t.Error("Sorting went wrong")
		}
	}

	utils.SortArray(arr, "DESC")

	for idx, element := range arr {
		if element != expectedDESC[idx] {
			t.Error("Sorting went wrong")
		}
	}
}
