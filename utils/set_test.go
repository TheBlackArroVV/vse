package utils_test

import (
	"elastic_go/models"
	"elastic_go/utils"
	"testing"
)

func TestSetInt_Add(t *testing.T) {
	set := utils.Set[int]{}

	if set.Length != 0 {
		t.Error("It's started not empty")
	}

	set.Add(1)

	if set.Length != 1 {
		t.Error("It's not writing")
	}

	if set.Values[0] != 1 {
		t.Error("Wrong value")
	}

	set.Add(1)

	if set.Length != 1 {
		t.Error("Duplicate!")
	}
}

func TestSetIndexDocument_Add(t *testing.T) {
	set := utils.IndexDocumentSet{}

	if set.Length != 0 {
		t.Error("It's started not empty")
	}

	set.Add(models.IndexDocument{
		Id: 1, Words: []string{},
	})

	if set.Length != 1 {
		t.Error("It's not writing")
	}

	if set.Values[0].Id != 1 {
		t.Error("Wrong value")
	}

	set.Add(models.IndexDocument{
		Id: 1, Words: []string{},
	})

	if set.Length != 1 {
		t.Error("Duplicate!")
	}
}
