package main

import (
	"cmp"
	"fmt"
	"slices"
)

func sorting_by_funcion() {
	fmt.Println("Data is the key")
	fruits := []string{"peach", "banana", "kiwi"}
	lencp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lencp)
	fmt.Println(fruits)
}
