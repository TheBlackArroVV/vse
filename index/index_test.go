package index

import (
	"testing"
)

func TestWrite(t *testing.T) {
	index := New("index")

	if len(index.documents) != 0 {
		t.Error("Documents are not empty")
	}

	index.Write("test")

	if len(index.documents) == 0 {
		t.Error("Writing is broken")
	}

	if len(index.documents) != 1 {
		t.Error("Something is wrong with writing")
	}

	for _, words := range index.documents {
		if words.words[0] != "test" {
			t.Error("Something wrong with index naming")
		}
	}
}
