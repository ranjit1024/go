package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(text) {
		counts[word]++
	}
	return counts
}

func demoWordCount() {
	text := "go is simple and go is powerful and go is fast"
	fmt.Println(text)
	fmt.Println(countWords(text))
}
