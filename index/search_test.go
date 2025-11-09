package index

import (
	"testing"
)

func TestSearch(t *testing.T) {
	index := New("index")

	index.Write("test", "test")
	index.Write("other", "other")

	results := index.Search("test")

	if len(results) != 1 {
		t.Error("too many results")
	}

	if results[0].name != "test" {
		t.Error("You got wrong result")
	}
}
