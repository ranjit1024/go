package main

import "fmt"

type Names []string

func SliceIndex[E comparable](s []E, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func main() {
	var names Names
	fmt.Println(names)
	SliceIndex([]string{"ranjit", "rahul"}, "ranjit")
}
