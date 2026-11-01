package main

import (
	"fmt"
	"regexp"
)

func demo_regex() {
	text := "My phone Number is 7571712612"
	re := regexp.MustCompile(`[0-9]{10}`)
	result := re.FindString(text)
	fmt.Println(result)
}
