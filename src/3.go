package main

import "fmt"

func main() {
	// var a [5]int
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	/// two d array

	twoD := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	a := "12";
	fmt.Println(a)
	fmt.Println(twoD)
	fmt.Println("Working")
}
