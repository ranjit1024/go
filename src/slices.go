package main

import "fmt"

func main() {
	fmt.Println("Data")
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:3]

	fmt.Println(b)
	fmt.Println(cap(b))

	b = append(b, 12)
	b = append(b, 34)

	fmt.Println(b)

	b = append(b, 100)
	fmt.Println(b)
}
