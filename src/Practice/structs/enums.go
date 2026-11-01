package main

import "fmt"

type ServerState int

const (
	Pending ServerState = iota
	Running
	Stoped
	Retry
)

func enum_demo() {

	fmt.Println("Data is the keys", Pending)
}
