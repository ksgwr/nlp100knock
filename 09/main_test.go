package main

import "testing"

func TestTypoglycemia(t *testing.T) {
	input := str

	actual := Typoglycemia(input)

	expected := "I coldnu't bieelve that I colud aculatly uarsntednd what I was raeding : the pmnohenael pweor of the hmuan mind ."

	if actual != expected {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}
