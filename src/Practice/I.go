package main

import (
	"fmt"
)

// / stuende grade analuzer
func first() {
	scores := []int{56, 99, 34, 23, 90, 34}
	fmt.Println(scores)
	h_n := scores[0]
	l_n := scores[0]

	for i := 1; i < len(scores); i++ {
		if h_n < scores[i] {
			h_n = scores[i]
			fmt.Println(h_n)
		}
		if l_n > scores[i] {
			l_n = scores[i]
			fmt.Println(l_n)
		}
	}

	fmt.Println("Data is the key")
}
