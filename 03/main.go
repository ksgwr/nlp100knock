package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

const str = "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."

var stemReg = regexp.MustCompile(`[,.]$`)

func main() {
	fmt.Println(WordLengthList(str))
}

func WordLengthList(str string) []int {
	words := strings.Split(str, " ")
	res := make([]int, len(words))

	for i, w := range words {
		w = stem(w)
		res[i] = utf8.RuneCountInString(w)
	}

	return res
}

func stem(s string) string {
	return stemReg.ReplaceAllString(s, "")
}