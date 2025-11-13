package utils_test

import (
	"elastic_go/utils"
	"testing"
)

func TestArraySort(t *testing.T) {
	arr := []int64{5, 3, 4, 1, 2}
	expected := []int64{1, 2, 3, 4, 5}

	utils.SortArray(arr)

	for idx, element := range arr {
		if element != expected[idx] {
			t.Error("Sorting went wrong")
		}
	}
}
