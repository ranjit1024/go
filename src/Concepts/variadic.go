package main

import "fmt"

func variadicSum(nums ...int) {
	fmt.Println(nums)
}
func demoVariadic() {
	fmt.Println("Data")
	variadicSum(1, 2, 2, 3, 4, 5)
}
