package main

import "fmt"

func main() {
	// fmt.Println("")
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i * 2
	}
	b := a[:]
	b = append(b, 222, 3434)
	fmt.Println(b)
	fmt.Println(a)
}
