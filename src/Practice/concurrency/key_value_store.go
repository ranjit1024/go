package main

import (
	"fmt"
	"sync"
)

type read_op struct {
	key  string
	resp chan int
}
type write_op struct {
	key  string
	val  int
	resp chan bool
}

func demoKeyValueStore() {
	reads := make(chan read_op)
	writes := make(chan write_op)

	go func() {
		state := map[string]int{"apple": 10}
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	var wg sync.WaitGroup
	for range 5 {
		wg.Go(func() {
			for range 3 {
				read := read_op{key: "apple", resp: make(chan int)}
				reads <- read
				fmt.Println("apple =", <-read.resp)
			}
		})
	}

	for i := range 5 {
		wg.Go(func() {
			write := write_op{key: "apple", val: i + 1, resp: make(chan bool)}
			writes <- write
			<-write.resp
		})
	}
	wg.Wait()

	final := read_op{key: "apple", resp: make(chan int)}
	reads <- final
	fmt.Println("final apple =", <-final.resp)
}
