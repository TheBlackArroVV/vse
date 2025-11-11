package utils

import "slices"

type Set[T int | int64] struct {
	Values []T
	length uint64
}

func (set Set[T]) Add(newValue T) Set[T] {
	if slices.Contains(set.Values, newValue) {
		return set
	}

	set.Values = append(set.Values, newValue)
	set.length = set.length + 1

	return set
}
