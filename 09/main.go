package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

const str = "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(Typoglycemia(str))
}

func Typoglycemia(str string) string {
	buf := bytes.NewBuffer(make([]byte, 0, len(str)))
	words := strings.Split(str, " ")

	for i, last := 0, len(words) - 1; i <= last; i++ {
		w := words[i]
		if utf8.RuneCountInString(w) > 4 {
			w = shuffle(w)
		}
		buf.WriteString(w)
		if (i != last) {
			buf.WriteByte(' ')
		}
	}
	return buf.String()
}

func shuffle(word string) string {
	r := []rune(word)
	for i, l := 1, len(r) - 2; i < l; i++ {
		j := rand.Intn(i + 1) + 1
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}