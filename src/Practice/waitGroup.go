package main

import (
	"fmt"
	"sync"
	"time"
)

func worker_wait_group(id int) {
	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Second)
}

func demo_wait_group() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker_wait_group(i)
		})
	}
	wg.Wait()
	fmt.Println("Hello World...")
}
