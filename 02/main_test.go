package main

import "testing"

func TestMixString(t *testing.T) {
	input1 := str1
	input2 := str2

	actual, _ := MixString(input1, input2)

	expected := "パタトクカシーー"

	if actual != expected {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}

func TestMixStringError(t *testing.T) {
	input1 := str1
	input2 := ""

	_, err := MixString(input1, input2)

	if err == nil {
		t.Errorf("Required error")
	}
}