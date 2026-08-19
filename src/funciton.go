package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}
func main() {
	a := [4]int{1, 2, 3, 4}
	s := make([]int, 2, 4)
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println("Data is the king")
	res := sum(3, 6)
	fmt.Println(res)
}
