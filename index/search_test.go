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

func TestSearchByShouldQuery(t *testing.T) {
	index := New("index")

	index.Write("test", "test")
	index.Write("other", "other")

	params := make(map[string][]string)
	params["should"] = []string{"test", "other"}
	results := index.SearchByQuery(params)

	if len(results) != 2 {
		t.Error("too little results")
	}

	if results[0].name != "test" {
		t.Error("You got wrong result")
	}

	if results[1].name != "other" {
		t.Error("You got wrong result")
	}
}

func TestSearchByMustQuery(t *testing.T) {
	index := New("index")

	index.Write("test", "test other")
	index.Write("other", "other")

	params := make(map[string][]string)
	params["must"] = []string{"test", "other"}
	results := index.SearchByQuery(params)

	if len(results) != 1 {
		t.Error("too little results")
	}

	if results[0].name != "test" {
		t.Error("You got wrong result")
	}
}
