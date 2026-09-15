package main

import "fmt"

type status int

const (
	Pending status = iota
	Running
	Completed
	Failed
)

func demoEnums() {
	fmt.Println("fadsfa")
	fmt.Println(Pending, Running, Completed, Failed)
}
