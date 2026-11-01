package main

import "fmt"

func demoChannels() {
	message := make(chan string)
	go func() {
		message <- "Ping"
	}()
	msg := <-message
	fmt.Println(msg)
	fmt.Println("Hello World...")
}
