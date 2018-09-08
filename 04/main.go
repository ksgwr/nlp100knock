package main

import (
	"fmt"
	"strings"
)

const str = "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."

func main() {
	fmt.Println(CreateElemntMap(str))
}

func CreateElemntMap(s string) map[string]int {
	words := strings.Split(s, " ")
	res := make(map[string]int, len(words))

	for i, w := range words {
		switch i {
		case 0, 4, 5, 6, 7, 8, 14, 15, 18:
			res[w[0:1]] = i
		default:
			res[w[0:2]] = i
		}
	}
	return res
}
