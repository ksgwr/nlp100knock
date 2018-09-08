package main

import (
	"reflect"
	"testing"
)

func TestWordLengthList(t *testing.T) {
	input := str

	actual := WordLengthList(input)

	expected := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}


