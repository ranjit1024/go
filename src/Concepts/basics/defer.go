package main

import (
	"fmt"
)

func demo_defer() {
	defer fmt.Println("This is last run ")
	fmt.Println("This is first")
	fmt.Println("Data is the key")
}
