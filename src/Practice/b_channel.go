package main

import "fmt"

func demo_b_channel() {
	message := make(chan string, 2)
	message <- "Hellow"
	message <- "Workld"
	message <- "Ranjit"
	fmt.Println("Hello world..")

}
