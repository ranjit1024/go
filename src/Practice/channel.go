package main

import "fmt"

func demo_channel() {
	message := make(chan string)
	go func() {
		message <- "Ping"
	}()
	msg := <-message
	fmt.Println(msg)
	fmt.Println("Hello World...")
}
