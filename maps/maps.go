package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	stringList := strings.Fields(s)
	wordMaps := make(map[string]int)
	for _, word := range stringList {
		wordMaps[word] += 1
	}
	return wordMaps
}

func main() {
	wc.Test(WordCount)
}
