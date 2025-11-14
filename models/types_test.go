package models_test

import (
	"testing"
	"vse/models"
)

func TestIndexDocument_Equals(t *testing.T) {
	tests := []struct {
		name          string
		otherDocument models.IndexDocument
		want          bool
	}{
		{
			name:          "when ids are the same",
			otherDocument: models.IndexDocument{Id: 1, Words: []string{}},
			want:          true,
		},
		{
			name:          "when ids and words are different",
			otherDocument: models.IndexDocument{Id: 2, Words: []string{"other"}},
			want:          false,
		},
		{
			name:          "when ids are different and words are the same",
			otherDocument: models.IndexDocument{Id: 2, Words: []string{"test"}},
			want:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexdocument := models.IndexDocument{Id: 1, Words: []string{"test"}}
			got := indexdocument.Equals(tt.otherDocument)
			if got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
