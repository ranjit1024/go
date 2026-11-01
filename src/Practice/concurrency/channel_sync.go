package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func demo_sync_channel() {
	status := make(chan bool, 1)
	go worker(status)

	stat := <-status
	fmt.Println(stat)
	fmt.Println("Hello World...")
}
