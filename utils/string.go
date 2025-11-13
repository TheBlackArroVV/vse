package utils

import "strings"

func LevenshteinDistance(word string, comparable string) int {
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

	distance = distance + AbsInt(len(splitedWord)-len(splitedComparable))

	return distance
}
