package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func demoChannelDirections() {
	pings := make(chan string)
	pongs := make(chan string)

	go ping(pings, "message passed")
	go pong(pings, pongs)

	fmt.Println(<-pongs)
}
