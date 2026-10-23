package main

import "fmt"

func rate_limter() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	fmt.Println("Hello World")
}
