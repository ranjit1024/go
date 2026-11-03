package main

import (
	"fmt"
	"time"
)

func epoch_demo() {

	fmt.Println("Hello World....")
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
}
