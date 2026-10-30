package main

import "fmt"

func demoStrings() {
	s := "A世😀"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
}
