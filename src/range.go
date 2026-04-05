package main

import "fmt"

func main() {
	fmt.Println("Data")
	nums := []int{10, 20, 30, 40}

	for _, val := range nums {
		fmt.Println(val)
	}
}
