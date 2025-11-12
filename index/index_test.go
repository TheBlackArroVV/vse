package index

import (
	"testing"
)

func TestWrite(t *testing.T) {
	index := New("index")

	if index.currentIdx != 0 {
		t.Error("Documents are not empty")
	}

	index.Write("test")

	if index.currentIdx == 0 {
		t.Error("Writing is broken")
	}

	if index.currentIdx != 1 {
		t.Error("Something is wrong with writing")
	}

	for idx, words := range index.documents {
		if words.words[0] != "test" {
			t.Error("Something wrong with index naming")
		}

		if idx != 1 {
			t.Error("Something wrong with index naming")
		}
	}
}
