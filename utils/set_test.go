package utils_test

import (
	"elastic_go/utils"
	"testing"
)

func TestSet_Add(t *testing.T) {
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
