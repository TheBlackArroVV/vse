package index

import (
	"testing"
)

func TestWrite(t *testing.T) {
	index := New("index")

	if len(index.documents) != 0 {
		t.Error("Documents are not empty")
	}

	index.Write("test", "test")

	if len(index.documents) == 0 {
		t.Error("Writing is broken")
	}

	if len(index.documents) != 1 {
		t.Error("Something is wrong with writing")
	}

	if index.documents[0].name != "test" {
		t.Error("Something wrong with index naming")
	}

	if len(index.documents[0].words) != 1 {
		t.Error("something wrong with index value length")
	}

	if index.documents[0].words[0] != "test" {
		t.Error("something wrong with index value")
	}
}
