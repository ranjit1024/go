package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	res := make(map[string]int)
	words := strings.Split(text, "")
	fmt.Println(words)
	return res
}

func Second() {
	text := "go is simple and go is powerful and go is fast"
	fmt.Println(text)

	fmt.Println("Data is the oil")

}
