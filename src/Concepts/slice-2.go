package main

import (
	"fmt"
)

func demoSlice2() {
	s := make([]int, 0, 3)
	fmt.Println(cap(s))
}
