package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	b = [...]int{12, 12, 12, 34, 34}
	fmt.Println(b)

}
