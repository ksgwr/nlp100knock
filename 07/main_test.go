package main

import "testing"

func TestTemplate(t *testing.T) {
	actual := Template(12, "気温", 22.4)

	expected := "12時の気温は22.4"

	if actual != expected {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}
