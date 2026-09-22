package main

import (
	"fmt"
)

// / stuende grade analuzer
func first() {
	scores := []int{56, 99, 34, 23, 90, 34}
	fmt.Println(scores)

	for i := range len(scores) {
		for j := i; j < len(scores); j++ {
			if scores[i] > scores[j] {
				fmt.Println(scores[i])
			}
		}
	}
	for i := range len(scores) {

		for j := i; j < len(scores); j++ {
			if scores[i] > scores[j] {
				fmt.Println(scores[i])
			}
		}
		fmt.Println()

	}

	fmt.Println("Data is the key")
}
