package main

import "fmt"

func zeroveal(n int) {
	n = 12
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	zeroveal(i)
	zeroPtr(&i)
	fmt.Println(i)
	fmt.Println("Data is the king")

}
