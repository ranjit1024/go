package main

import (
	"cmp"
	"fmt"
	"slices"
)

func sorting_by_funcion() {
	a := []string{"Ranjit", "Sammmmmmmu", "Vishal"}
	fmt.Println(a)
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(b), len(a))
	}
	slices.SortFunc(a, lenCmp)
	fmt.Println(a)
	fmt.Println("Hello World....")
}
