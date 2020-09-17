package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount converts a string into a map containing count of every unique word
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordMap := make(map[string]int)

	for _, word := range words {
		wordMap[word]++
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}
