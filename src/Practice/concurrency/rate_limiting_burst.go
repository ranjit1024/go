package main

import (
	"fmt"
	"time"
)

func demoRateLimitingBurst() {
	jobs := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	burstLimiter := make(chan time.Time, 3)
	for range 3 {
		burstLimiter <- time.Now()
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for tick := range ticker.C {
			burstLimiter <- tick
		}
	}()

	for job := range jobs {
		<-burstLimiter
		fmt.Println("processing job", job, "at", time.Now())
	}
}
