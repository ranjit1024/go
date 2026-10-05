package main

import (
	"cmp"
	"fmt"
	"slices"
)

type student struct {
	name string
	age  int
}

func sorting_by_funcion() {
	fruits := []string{"Peaach", "banaana", "apple", "Watermelon"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	people := []student{
		student{name: "Samruddhi", age: 45},
		student{name: "Ranjit", age: 12},
		student{name: "Sammmmu", age: 90},
	}

	fmt.Println(people)

	slices.SortFunc(people, func(a, b student) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
}
