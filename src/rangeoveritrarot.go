package main

import "fmt"

func doSomething(f func(int)) {
	f(10)
}

func main() {
	doSomething(func(i int) {
		fmt.Println(i)
	})
}
