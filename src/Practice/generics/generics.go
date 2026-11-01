package main

import (
	"cmp"
	"fmt"
)

func demoGenerics() {
	fmt.Println(maxOf(3, 7))
	fmt.Println(maxOf("apple", "pear"))
}

func maxOf[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
