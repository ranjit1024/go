package main

import (
	"fmt"
	"time"
)

func syncWorker(done chan bool) {
	fmt.Println("Working....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func demoChannelSync() {
	status := make(chan bool, 1)
	go syncWorker(status)

	stat := <-status
	fmt.Println(stat)
	fmt.Println("Hello World...")
}
