package main

import (
	"fmt"
)

func main() {
	fmt.Println("data")
	a := [4]int{1, 2, 3, 4}
	len(a)
	b := a[:]
	b = append(b, 12)
	fmt.Println(b)
}
