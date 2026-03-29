package main

import "fmt"

func plus(a int, b int) int {
	if a == 12 && b == 12 {
		return a * b
	}
	return a + b
}

func main() {
	res := plus(12, 12)

	fmt.Println(res)
	fmt.Println("fad")
}
