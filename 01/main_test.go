package main

import (
	"testing"
)

func TestConcatOdd(t *testing.T) {
	input := str

	actual := ExstractOddCharacter(input)

	expected := "パトカー"

	if actual != expected {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}