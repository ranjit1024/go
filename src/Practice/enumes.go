package main

import "fmt"

type jobState int
type eventState int

const (
	Pending jobState = iota
	Running
	Completed
	Failed
	Retrying
)

const (
	Start eventState = iota
	Success
	Failure
	Retry
)

var JobMap = map[jobState]string{
	Pending:   "Pending",
	Running:   "running",
	Completed: "Competed",
	Failed:    "Failed",
	Retrying:  "Retrying",
}

var EvnetMap = map[eventState]string{
	Start:   "Start",
	Success: "Sucess",
	Failure: "failed",
	Retry:   "retry",
}

func (s jobState) String() string {
	return JobMap[s]
}
func transition(state jobState, event eventState) jobState {
	if state == Pending {
		if event == Start {
			state = Running
		}
	}
	return state
}

func enum_demo() {
	test := transition(Pending, Start)
	fmt.Println(test)
	fmt.Println("Data is the king")
}
