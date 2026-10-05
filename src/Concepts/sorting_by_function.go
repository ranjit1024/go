package main

import (
	"cmp"
	"fmt"
	"slices"
)

func sorting_by_funcion() {
	fruits := []string{"Peaach", "banaana", "apple", "Watermelon"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

}
