package main

import "testing"

func TestReverse(t *testing.T) {
	input := str

	actual := Reverse(input)

	expected := "desserts"

	if actual != expected {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}