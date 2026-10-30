package main

import "fmt"

func demoBasicChannel() {
	fmt.Println("Data is the king")
	stage1 := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		stage1 <- i
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(<-stage1)
	}
}
