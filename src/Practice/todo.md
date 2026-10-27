package main

import (
	"fmt"
	"sync"
)

type readOp struct {
	key  string
	resp chan int
}

type writeOp struct {
	key  string
	val  int
	resp chan bool
}

func main() {
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// The handler is the ONLY goroutine that accesses state.
	go func() {
		state := map[string]int{
			"apple": 10,
		}

		for {
			select {
			case read := <-reads:
				// TODO 1:
				// Read the requested value from state.
				// Send it through read.resp.

			case write := <-writes:
				// TODO 2:
				// Update state using write.key and write.val.
				// Send true through write.resp.
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	// Reader goroutine
	go func() {
		defer wg.Done()

		req := readOp{
			key:  "apple",
			resp: make(chan int),
		}

		// TODO 3:
		// Send req through reads.
		// Receive the result from req.resp.
		// Print the result.
	}()

	// Writer goroutine
	go func() {
		defer wg.Done()

		req := writeOp{
			key:  "apple",
			val:  20,
			resp: make(chan bool),
		}

		// TODO 4:
		// Send req through writes.
		// Wait for the acknowledgment.
		// Print whether the write completed.
	}()

	wg.Wait()
	fmt.Println("All workers finished")
}
