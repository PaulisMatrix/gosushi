package main

import (
	"strings"
	//	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	myWords := strings.Fields(s)

	result := make(map[string]int)

	for _, word := range myWords {
		_, ok := result[word]
		if ok {
			result[word]++
		} else {
			result[word] = 1
		}
	}
	return result
}

//func main() {
//	wc.Test(WordCount)
//}
