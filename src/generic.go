package main

import "fmt"

func Add[T int | float64](a T, b T) T {
	return a + b
}
func main() {
	a := Add(12, 12.4)
	fmt.Println(a)
}
