package main

import (
	"cmp"
	"fmt"
	"slices"
)

func demoSortingByFunction() {
	names := []string{"Ranjit", "Sammmmmmmu", "Vishal"}
	fmt.Println("before:", names)

	slices.SortFunc(names, func(a, b string) int {
		return cmp.Compare(len(b), len(a))
	})
	fmt.Println("after: ", names)
}
