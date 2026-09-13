package main

import "fmt"

func main() {
	message := make(chan string)
	// signals := make(chan string)

	select {
	case msg := <-message:
		fmt.Println("Message received", msg)
	default:
		fmt.Println("No message recieved")
	}

	msg := "Hi"

	select {
	case message <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("No messsage sent")
	}

}
