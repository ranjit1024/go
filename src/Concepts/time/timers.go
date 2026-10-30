package main

import (
	"fmt"
	"time"
)

func demoTimers() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg1 := <-c2:
			fmt.Println(msg1)
		}

	}
}
