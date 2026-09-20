package main

import (
	"fmt"
	"sync"
	"time"
)

func worker_1(id int) {
	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done \n ", id)
}

func demo_waitgroups() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Go(func() {
			worker_1(i)
		})
	}
	time.Sleep(time.Second)

}
