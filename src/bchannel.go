package main

import "fmt"

func main() {
	fmt.Println("Data is the king")
	stage1 := make(chan int, 3)
	stage1 <- 12
	for i := 1; i <= 10; i++ {
		fmt.Println(<-stage1)
	}
}
