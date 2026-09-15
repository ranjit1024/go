package main

import "fmt"

func demoRangeOverChannel() {
	queue := make(chan string, 2)
	queue <- "Ranjit"
	queue <- "Sammmu"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
	fmt.Println("Data is the oil")
}
