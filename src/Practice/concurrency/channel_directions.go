package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func demo_channel_direction() {
	pings := make(chan string)
	pongs := make(chan string)
	ping(pings, "Message passed")
	pong(pings, pongs)
	fmt.Println("Hello World..")
}
