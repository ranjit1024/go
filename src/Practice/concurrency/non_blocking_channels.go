package main

import "fmt"

func demo_non_blocking_channel() {
	message := make(chan string)
	select {
	case msg := <-message:
		fmt.Println("Received messge", msg)
	default:
		fmt.Println("No message received")
	}
	fmt.Println("Hello World..")
}
