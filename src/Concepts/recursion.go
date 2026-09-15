package main

import "fmt"

func fact(n int) int{
	if n == 0 {
		return  1
	}
	return n * fact(n-1)
}
func demoRecursion() {
	fmt.Println("Data is the king ")
	res := fact(5)
	fmt.Println(res)
}
