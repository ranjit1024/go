package main

import "fmt"

func count(yield func(int) bool) {
	for i := 1; i <= 5; i++ {
		if !yield(i) {
			return
		}
	}
}

func main() {
	for x := range count {
		fmt.Println(x)
	}
	count(func(i int) bool {
		return true
	})
}
