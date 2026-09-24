package main

import (
	"fmt"
)

// / stuende grade analuzer
func first() {
	scores := []int{56, 99, 34, 23, 90, 34}

	h_n := scores[0]
	l_n := scores[0]
	sum := 0
	for i := 1; i < len(scores); i++ {
		if h_n < scores[i] {
			h_n = scores[i]

		}
		if l_n > scores[i] {
			l_n = scores[i]

		}
	}
	for i := range len(scores) {
		sum = sum + scores[i]
	}
	fmt.Println(sum)

	fmt.Println("Data is the key", h_n, l_n)
}
