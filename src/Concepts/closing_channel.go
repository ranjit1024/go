package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Receinved job", j)

			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Send job")
	}
	close(jobs)
	fmt.Println("Send all jobs")
	<-done

	_, ok := <-jobs
	fmt.Println("Received more Jobs", ok)
}
