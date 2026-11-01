package main

import (
	"fmt"
	"sync"
	"time"
)

func waitGroupWorker(id int) {
	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Second)
}

func demoWaitGroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			waitGroupWorker(i)
		})
	}
	wg.Wait()
	fmt.Println("Hello World...")
}
