package main

import "fmt"

type ServerState int

const (
	Pending ServerState = iota
	Running
	Stopped
	Retry
)

func (s ServerState) String() string {
	switch s {
	case Pending:
		return "pending"
	case Running:
		return "running"
	case Stopped:
		return "stopped"
	case Retry:
		return "retry"
	default:
		return "unknown"
	}
}

func demoEnums() {
	for _, state := range []ServerState{Pending, Running, Stopped, Retry} {
		fmt.Println(state)
	}
}
