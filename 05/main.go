package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

const str = "I am an NLPer"

func main() {
	res, err := Ngram(str, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(res, "/"))

	res2, err2 := Ngram(strings.Split(str, " "), 2)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(res2)
}

func Ngram(input interface{}, n int) ([]string, error) {
	var res []string

	switch reflect.ValueOf(input).Kind() {
	case reflect.String:
		str := input.(string)
		for i := 0; i < len(str) - n + 1; i++ {
			res = append(res, str[i:i+n])
		}
	case reflect.Slice:
		arr := input.([]string)
		for i := 0; i < len(arr) - n + 1; i++ {
			res = append(res, strings.Join(arr[i:i+n], "/"))
		}
	default:
		return res, fmt.Errorf("unsupported type")
	}
	return res, nil
}
