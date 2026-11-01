package main

import "fmt"

func mayPanic() {
	panic("a Problem")
}
func demoRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd Erroor", r)
		}
	}()

	mayPanic()
	fmt.Println("After maypanic")
}
