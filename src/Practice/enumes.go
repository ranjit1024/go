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

func (ss ServerState) String() string {
	return stateName[ss]
}

func enum_demo() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)

	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConneted
	case StateConneted, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Unknow state %s", s))
	}

}
