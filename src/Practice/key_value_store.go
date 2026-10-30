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

func key_value_store() {
	reads := make(chan read_op)
	writes := make(chan write_op)
	//first one is done
	go func() {
		state := map[string]int{
			"apple": 10}

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
	wg.Add(2)

	for range 100 {

		go func() {
			defer wg.Done()
			for {
				read := read_op{
					key:  "apple",
					resp: make(chan int),
				}
				reads <- read
				value := <-read.resp
				fmt.Println(value)
			}
		}()
	}

	for range 10 {
		defer wg.Done()
		go func() {

		}()
	}
	fmt.Println("heelo World...")
}
