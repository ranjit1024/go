package main

import (
	"fmt"
	"time"
)

func tickerDemo() {

	ticker1 := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker1.C:
				fmt.Println("Ticker at", t)
			}
		}
	}()

	time.Sleep(2000 * time.Millisecond)
	ticker1.Stop()
	done <- true
	fmt.Println("Timr Stoped")

}
