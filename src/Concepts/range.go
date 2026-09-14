package main

import "fmt"

func main() {
	nums := []int{12, 1, 212, 1}
	for index, num := range nums {
		fmt.Println(index)
		fmt.Println(num)
	}
}
