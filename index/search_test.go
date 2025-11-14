package index

import (
	"elastic_go/models"
	"testing"
)

func TestSearch(t *testing.T) {
	index := New("index")

	index.Write("test")
	index.Write("other")
	index.Write("tesc")
	index.Write("tes")
	index.Write("test1")

	results := index.Search("test")

	if len(results) != 4 {
		t.Error("too many results")
	}

	if results[0].Words[0] != "test" {
		t.Error("You got wrong result")
	}

	if results[1].Words[0] != "tesc" {
		t.Error("You got wrong result")
	}

	if results[2].Words[0] != "tes" {
		t.Error("You got wrong result")
	}

	if results[3].Words[0] != "test1" {
		t.Error("You got wrong result")
	}
}

func TestSearchByShouldQuery(t *testing.T) {
	index := New("index")

	index.Write("test")
	index.Write("other")
	index.Write("tesc")
	index.Write("tes")
	index.Write("test1")

	query := models.Query{
		Should: []string{"test", "other"},
	}
	results := index.SearchByQuery(query)

	if len(results) != 5 {
		t.Error("too little results")
	}

	if results[0].Words[0] != "test" {
		t.Error("You got wrong result")
	}

	if results[1].Words[0] != "other" {
		t.Error("You got wrong result")
	}

	if results[2].Words[0] != "tesc" {
		t.Error("You got wrong result")
	}

	if results[3].Words[0] != "tes" {
		t.Error("You got wrong result")
	}

	if results[4].Words[0] != "test1" {
		t.Error("You got wrong result")
	}
}

func TestSearchByMustQuery(t *testing.T) {
	index := New("index")

	index.Write("test other")
	index.Write("other")

	query := models.Query{
		Must: []string{"test", "other"},
	}
	results := index.SearchByQuery(query)

	if len(results) != 1 {
		t.Error("too little results")
	}

	if results[0].Words[0] != "test" {
		t.Error("You got wrong result")
	}
}

func TestSearchByBothMustAndShouldQuery(t *testing.T) {
	index := New("index")

	index.Write("test other")
	index.Write("other")
	index.Write("second")
	index.Write("secon")
	index.Write("seconc")

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

	if results[1].Words[0] != "second" {
		t.Error("You got wrong result")
	}

	if results[2].Words[0] != "secon" {
		t.Error("You got wrong result")
	}

	if results[3].Words[0] != "seconc" {
		t.Error("You got wrong result")
	}
}
