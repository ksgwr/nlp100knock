package main

import (
	"reflect"
	"testing"
)

func TestCreateElemntMap(t *testing.T) {
	input := str

	actual := CreateElemntMap(input)

	expected := map[string]int{"K": 18, "Li": 2, "Ne": 9, "B": 4, "C": 5, "N": 6, "Si": 13, "Cl": 16, "H": 0, "He": 1, "Mi": 11, "P": 14, "S": 15, "Ar": 17, "Ca": 19, "O": 7, "Na": 10, "Al": 12, "Be": 3, "F": 8}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}