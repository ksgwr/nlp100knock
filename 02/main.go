package main

import (
	"bytes"
	"fmt"
	"log"
)

const str1 = "パトカー"
const str2 = "タクシー"

func main() {
	res, err := MixString(str1, str2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func MixString(str1 string, str2 string) (string, error) {
	l1 := len(str1)
	l2 := len(str2)

	if l1 != l2 {
		return "", fmt.Errorf("word length is diffrent, s1 = %s, s2 = %s", str1, str2)
	}

	buf := bytes.NewBuffer(make([]byte, 0, l1 + l2))
	r1 := []rune(str1)
	r2 := []rune(str2)

	for i := 0; i < len(r1); i++ {
		buf.WriteRune(r1[i])
		buf.WriteRune(r2[i])
	}

	return buf.String(), nil
}