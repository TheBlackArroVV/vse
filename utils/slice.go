package utils

import "slices"

func SortArray(arr []int64) {
	slices.SortFunc(arr, func(a, b int64) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

}
