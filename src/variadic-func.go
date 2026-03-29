package main

import "fmt"

func sum(num ...int) {
	total := 0
	for _, num := range num {
		total += num
	}
	fmt.Println(total)
}
func main() {
	sum(1, 2, 3, 4, 5, 5)
}
