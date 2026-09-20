package main

import (
	"fmt"
	"time"
)

func demo_rateLimiter() {
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("Requests", req, time.Now())
	}

	brust_limiter := make(chan time.Time, 3)
	fmt.Println(brust_limiter)
}
