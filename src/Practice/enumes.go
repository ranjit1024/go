package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConneted
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:     "Idel",
	StateConneted: "Connected",
	StateError:    "Error",
	StateRetrying: "Retry",
}

func enum_demo() {
	fmt.Println("Data is the king")
}
