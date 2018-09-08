package main

import (
	"reflect"
	"testing"
)

func TestNgramStr(t *testing.T) {
	input := str

	actual, _ := Ngram(input, 2)

	expected := []string{"I ", " a", "am", "m ", " a", "an", "n ", " N", "NL", "LP", "Pe", "er"}

	if (!reflect.DeepEqual(actual, expected)) {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}

func TestNgramSlice(t *testing.T) {
	input := []string{"I", "am", "an", "NLPer"}

	actual, _ := Ngram(input, 2)

	expected := []string{"I/am", "am/an", "an/NLPer"}

	if (!reflect.DeepEqual(actual, expected)) {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}