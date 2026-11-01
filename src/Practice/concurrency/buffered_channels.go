package main

import "fmt"

func demoBufferedChannels() {
	messages := make(chan string, 3)
	messages <- "Hello"
	messages <- "World"
	messages <- "Ranjit"
	close(messages)

	for message := range messages {
		fmt.Println(message)
	}
}
