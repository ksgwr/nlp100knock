package main

import (
	"bytes"
	"fmt"
)

const key = 219

const str = "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

func main() {
	cstr := Cipher(str)
	fmt.Println(cstr)
	dstr := Decrypytion(cstr)
	fmt.Println(dstr)
}

func Cipher(str string) string {
	buf := bytes.NewBuffer(make([]byte, 0, len(str)))
	runes := []rune(str)
	for _, r := range runes {
		if 'a' < r && r < 'z' {
			r = key - r
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

func Decrypytion(str string) string {
	buf := bytes.NewBuffer(make([]byte, 0, len(str)))
	runes := []rune(str)
	const scode = key - 'z'
	const ecode = key - 'a'
	for _, r := range runes {
		if scode < r && r < ecode {
			r = key - r
		}
		buf.WriteRune(r)
	}
	return buf.String()
}