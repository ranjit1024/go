package main

import "fmt"

func count(yield func(int) bool) {
	yield(1)
	yield(2)
	yield(3)
}

func main() {

	x := count(func(i int) bool {
		return true
	})
	fmt.Println(x)
	fmt.Println("Data is the oil")
}
