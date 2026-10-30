package main

import "fmt"

func intSeq() func() int {
	i := 12

	return func() int {
		i++
		return i
	}

}

func demoClosure() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
