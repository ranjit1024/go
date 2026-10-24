package main

import (
	"fmt"
	"time"
)

func rate_limter() {
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("Req", req, time.Now())
	}

	brustyLimiter := make(chan time.Time, 3)

	for range 3 {
		brustyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			brustyLimiter <- t
		}
	}()

	brustyRequest := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		brustyRequest <- i
	}

	fmt.Println("Hello World")
}
