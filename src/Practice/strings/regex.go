package main

import (
	"fmt"
	"regexp"
)

func demoRegex() {
	text := "My phone Number is 7571712612"
	re := regexp.MustCompile(`[0-9]{10}`)
	result := re.FindString(text)
	fmt.Println(result)
}
