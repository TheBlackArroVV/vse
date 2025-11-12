package index

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.SkipNow()
	index := New("index")

	index.Write("test")
	index.Write("other")

	results := index.Search("test")

	if len(results) != 1 {
		t.Error("too many results")
	}

	if results[0].words[0] != "test" {
		t.Error("You got wrong result")
	}
}

func TestSearchByShouldQuery(t *testing.T) {
	index := New("index")

	index.Write("test")
	index.Write("other")

	query := Query{
		Should: []string{"test", "other"},
	}
	results := index.SearchByQuery(query)

	if len(results) != 2 {
		t.Error("too little results")
	}

	if results[0].words[0] != "test" {
		t.Error("You got wrong result")
	}

	if results[1].words[0] != "other" {
		t.Error("You got wrong result")
	}
}

func TestSearchByMustQuery(t *testing.T) {
	t.SkipNow()
	index := New("index")

	index.Write("test other")
	index.Write("other")

	query := Query{
		Must: []string{"test", "other"},
	}
	results := index.SearchByQuery(query)

	if len(results) != 1 {
		t.Error("too little results")
	}

	if results[0].words[0] != "test" {
		t.Error("You got wrong result")
	}
}

func TestSearchByBothMustAndShouldQuery(t *testing.T) {
	t.SkipNow()
	index := New("index")

	index.Write("test other")
	index.Write("other")
	index.Write("second")

	query := Query{
		Must:   []string{"test", "other"},
		Should: []string{"second"},
	}
	results := index.SearchByQuery(query)

	if len(results) != 2 {
		t.Error("too little results")
	}

	if results[0].words[0] != "second" {
		t.Error("You got wrong result")
	}

	if results[1].words[0] != "test" {
		t.Error("You got wrong result")
	}

	if results[1].words[1] != "other" {
		t.Error("You got wrong result")
	}
}
