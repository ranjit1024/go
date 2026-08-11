package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	marks := make([]int, 4, 10)
	fmt.Println(marks)
	m[1] = "ranjit"
	fmt.Println(m)
	clear(m)
	a := 12
	fmt.Println(a)
}
