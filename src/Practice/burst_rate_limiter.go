package main

import (
	"fmt"
	"time"
)

func brust_limit() {
	jobs := make(chan int, 10)
	for i := range 10 {
		jobs <- i
	}
	close(jobs)

	b_rate_limimter := make(chan time.Time, 3)

	for range 3 {
		b_rate_limimter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Second) {
			b_rate_limimter <- t
		}
	}()

	for job := range jobs {
		<-b_rate_limimter
		fmt.Println("Processing job", job, "at", time.Now())
	}

}
