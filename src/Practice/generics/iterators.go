package main

import (
	"fmt"
	"iter"
)

func countToFive() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= 5; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func demoIterators() {
	for n := range countToFive() {
		fmt.Println(n)
	}
}
