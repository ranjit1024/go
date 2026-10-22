package main

import (
	"fmt"
	"time"
)

func worker_demo(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "Started Job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "Finished Job")
		results <- j * 2
	}
}

func demo_worker_pool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker_demo(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

	fmt.Println("Data")
}
