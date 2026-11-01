package main

import (
	"fmt"
	"time"
)

func printSequence(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}
func demoGoroutines() {
	go func(msg string) {
		fmt.Println(msg)
	}("Hellow World...")
	go printSequence("Goroutine")
	time.Sleep(time.Second)
	fmt.Println("Hellow World")

}
