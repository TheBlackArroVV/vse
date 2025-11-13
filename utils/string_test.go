package utils_test

import (
	"elastic_go/utils"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		name       string
		word       string
		comparable string
		want       int
	}{
		{
			name:       "when words are the same",
			word:       "test",
			comparable: "test",
			want:       0,
		},
		{
			name:       "when words are reversed",
			word:       "asdf",
			comparable: "fdsa",
			want:       4,
		},
		{
			name:       "when words different size",
			word:       "asdf",
			comparable: "asd",
			want:       1,
		},
		{
			name:       "when multiple words",
			word:       "test two",
			comparable: "test four",
			want:       4,
		},
		{
			name:       "when different amount of words",
			word:       "test two",
			comparable: "one",
			want:       8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.LevenshteinDistance(tt.word, tt.comparable)
			if got != tt.want {
				t.Errorf("LevenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
