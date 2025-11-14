package index

import (
	"elastic_go/models"
	"testing"
)

func TestSearch(t *testing.T) {
	index := New("index")

	index.Write("1", "test")
	index.Write("2", "other")
	index.Write("3", "tesc")
	index.Write("4", "tes")
	index.Write("5", "test1")

	results := index.Search("test")

	if len(results) != 4 {
		t.Error("too many results")
	}

	if results[0].Words[0] != "test" || results[0].Name != "1" {
		t.Error("You got wrong result")
	}

	if results[1].Words[0] != "tesc" || results[1].Name != "3" {
		t.Error("You got wrong result")
	}

	if results[2].Words[0] != "tes" || results[2].Name != "4" {
		t.Error("You got wrong result")
	}

	if results[3].Words[0] != "test1" || results[3].Name != "5" {
		t.Error("You got wrong result")
	}
}

func TestSearchByShouldQuery(t *testing.T) {
	index := New("index")

	index.Write("1", "test")
	index.Write("2", "other")
	index.Write("3", "tesc")
	index.Write("4", "tes")
	index.Write("5", "test1")

	query := models.Query{
		Should: []string{"test", "other"},
	}
	results := index.SearchByQuery(query)

	if len(results) != 5 {
		t.Error("too little results")
	}

	if results[0].Words[0] != "test" || results[0].Name != "1" {
		t.Error("You got wrong result")
	}

	if results[1].Words[0] != "other" || results[1].Name != "2" {
		t.Error("You got wrong result")
	}

	if results[2].Words[0] != "tesc" || results[2].Name != "3" {
		t.Error("You got wrong result")
	}

	if results[3].Words[0] != "tes" || results[3].Name != "4" {
		t.Error("You got wrong result")
	}

	if results[4].Words[0] != "test1" || results[4].Name != "5" {
		t.Error("You got wrong result")
	}
}

func TestSearchByMustQuery(t *testing.T) {
	index := New("index")

	index.Write("1", "test other")
	index.Write("2", "other")

	query := models.Query{
		Must: []string{"test", "other"},
	}
	results := index.SearchByQuery(query)

	if len(results) != 1 {
		t.Error("too little results")
	}

	if results[0].Words[0] != "test" || results[0].Name != "1" {
		t.Error("You got wrong result")
	}
}

func TestSearchByBothMustAndShouldQuery(t *testing.T) {
	index := New("index")

	index.Write("1", "test other")
	index.Write("2", "other")
	index.Write("3", "second")
	index.Write("4", "secon")
	index.Write("5", "seconc")

	query := models.Query{
		Must:   []string{"test", "other"},
		Should: []string{"second", "test"},
	}
	results := index.SearchByQuery(query)

	if len(results) != 4 {
		t.Error("too little results")
	}

	if results[0].Words[0] != "test" {
		t.Error("You got wrong result")
	}

	if results[0].Words[1] != "other" {
		t.Error("You got wrong result")
	}

	if results[0].Name != "1" {
		t.Error("You got wrong result")
	}

	if results[1].Words[0] != "second" || results[1].Name != "3" {
		t.Error("You got wrong result")
	}

	if results[2].Words[0] != "secon" || results[2].Name != "4" {
		t.Error("You got wrong result")
	}

	if results[3].Words[0] != "seconc" || results[3].Name != "5" {
		t.Error("You got wrong result")
	}
}
