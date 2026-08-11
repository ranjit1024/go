package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums)
}
func main() {
	fmt.Print("Data is kind ")
	sum(112, 12, 23, 3, 3, 3, 3)
}
