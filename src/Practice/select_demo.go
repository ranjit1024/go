package main
import (
	"fmt"
	"time"
)

func select_demo(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func(){
		time.Sleep(2 * time.Second)
		c2 <- "Two"
	}()

	for range 2 {
		select {
		case msg1 := <- c1:
			fmt.Println("Received", msg1)
		case msg2 := <- c2:
			fmt.Println("Received",  msg2)
		}

	}

	fmt.Println("Hello World...")
}