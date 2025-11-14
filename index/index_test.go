package index

import (
	"testing"
)

func TestWrite(t *testing.T) {
	index := New("index")

	if index.currentIdx != 0 {
		t.Error("Documents are not empty")
	}

	index.Write("name", "test")

	if index.currentIdx == 0 {
		t.Error("Writing is broken")
	}

	if index.currentIdx != 1 {
		t.Error("Something is wrong with writing")
	}

	for idx, words := range index.documents {
		if words.Words[0] != "test" {
			t.Error("Something wrong with index naming")
		}

		if idx != 1 {
			t.Error("Something wrong with index naming")
		}

		if words.Id != 1 {
			t.Error("Something wrong with index naming")
		}

		if words.Name != "name" {
			t.Error("Something wrong with index naming")
		}

	}
}

func TestFindDocumentsByIds(t *testing.T) {
	index := New("index")
	index.Write("first", "1")
	index.Write("second", "2")
	index.Write("third", "3")
	index.Write("fourth", "4")
	index.Write("fifth", "5")

	results := index.FindDocumentsByIds([]int64{2, 4})

	if len(results) != 2 {
		t.Error("Wrong number of results")
	}

	if results[0].Id != 2 || results[0].Name != "second" || results[0].Words[0] != "2" {
		t.Error("Wrong record id returned")
	}

	if results[1].Id != 4 || results[1].Name != "fourth" || results[1].Words[0] != "4" {
		t.Error("Wrong record id returned")
	}
}
