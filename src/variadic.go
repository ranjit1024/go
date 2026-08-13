package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums)
}
func main() {
	fmt.Println("Data")
	sum(1, 2, 2, 3, 4, 5)
}
