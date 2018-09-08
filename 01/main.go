package main

import (
	"bytes"
	"fmt"
)

const str = "パタトクカシーー"

func main() {
	fmt.Println(ExstractOddCharacter(str))
}

func ExstractOddCharacter(str string) string {
	buf := bytes.NewBuffer(make([]byte, 0, len(str)/2))
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		if i % 2 == 0 {
			buf.WriteRune(r[i])
		}
	}
	return buf.String()
}