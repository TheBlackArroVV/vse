package utils

import "strings"

func TransformStrings(string string) string {
	return normalize(strings.ToLower(string))
}

func normalize(string string) string {
	var sb strings.Builder

	for letter := range strings.SplitSeq(string, "") {
		sb.WriteString(normalizeLetter(letter))
	}

	return sb.String()
}

func normalizeLetter(letter string) string {
	return letter
}
