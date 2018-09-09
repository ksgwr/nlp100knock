package main

import "testing"

const cstr = "I xlfowm'g yvorvev gsag I xlfow axgfaoob fmwvihgamw dsag I dah ivawrmt : gsv ksvmlnvmao kldvi lu gsv sfnam nrmw ."

func TestCipher(t *testing.T) {
	input := str

	actual := Cipher(input)

	expected := cstr

	if actual != expected {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}

func TestDecrypytion(t *testing.T) {
	input := cstr

	actual := Decrypytion(input)

	expected := str

	if actual != expected {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}