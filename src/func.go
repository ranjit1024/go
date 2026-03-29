package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}
func constant() (float64, float64) {
	return 32.14, 1.18
}
func main() {
	fmt.Println("data this is data we are talking about this is main datae ")
	a := add(12, 12)
	_, dat := constant()
	fmt.Println(dat)
	fmt.Println(a)
}
