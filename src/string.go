package main

import "fmt"

func main() {
	s := "A世😀"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
}
