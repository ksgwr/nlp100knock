package main

import (
	"reflect"
	"testing"
	set "github.com/deckarep/golang-set"
)

var (
	X = set.NewSetFromSlice(Ngram(str1, 2))
	Y = set.NewSetFromSlice(Ngram(str2, 2))
)

func TestUnion(t *testing.T) {
	actual := X.Union(Y)

	expected := set.NewSetFromSlice([]interface{}{"ad", "ag", "ap", "ar", "di", "gr", "is", "pa", "ph", "ra", "se"})

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}

func TestIntersect(t *testing.T) {
	actual := X.Intersect(Y)

	expected := set.NewSetFromSlice([]interface{}{"ap", "ar", "pa", "ra"})

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}

func TestDifference(t *testing.T) {
	actual := X.Difference(Y)

	expected := set.NewSetFromSlice([]interface{}{"is", "se", "ad", "di"})

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}

func TestContainSe(t *testing.T) {
	actual1 := X.Contains("se")
	actual2 := Y.Contains("se")

	if !actual1 {
		t.Error("actual1 want true")
	}
	if actual2 {
		t.Error("actual2 want false")
	}
}