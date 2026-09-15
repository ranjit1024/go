package main

import "fmt"

func loop() {
	for i := range 3 {
		fmt.Println(i)
		i++
	}
}
