package main

import (
	"fmt"
)

func demoGrades() {
	scores := []int{56, 99, 34, 23, 90, 34}

	highest := scores[0]
	lowest := scores[0]
	sum := 0
	for _, score := range scores {
		if score > highest {
			highest = score
		}
		if score < lowest {
			lowest = score
		}
		sum += score
	}

	fmt.Println("scores: ", scores)
	fmt.Println("highest:", highest)
	fmt.Println("lowest: ", lowest)
	fmt.Println("sum:    ", sum)
	fmt.Println("average:", float64(sum)/float64(len(scores)))
}
