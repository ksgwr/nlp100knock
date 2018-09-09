package main

import (
	"fmt"
	set "github.com/deckarep/golang-set"
)

const str1 = "paraparaparadise"
const str2 = "paragraph"

func main() {
	setX := set.NewSetFromSlice(Ngram(str1, 2))
	setY := set.NewSetFromSlice(Ngram(str2, 2))

	fmt.Println(setX.Union(setY))
	fmt.Println(setX.Intersect(setY))
	fmt.Println(setX.Difference(setY))

	fmt.Println(setX.Contains("se"))
	fmt.Println(setY.Contains("se"))
}

func Ngram(str string, n int) []interface{} {
	var res []interface{}
	for i := 0; i < len(str) - n + 1; i++ {
		res = append(res, str[i:i+n])
	}
	return res
}
