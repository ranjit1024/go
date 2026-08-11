package main

import "fmt"

func test() (int, int) {
	return 12, 24
}

func main() {
	fmt.Println("Date is the king")
	a, b := test()
	fmt.Println(a)
	fmt.Println(b)

	_, c := test()
	fmt.Println(c)
}
