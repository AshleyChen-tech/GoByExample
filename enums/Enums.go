// Enumerated types (enums) are a special case of sum types. An enum is a type that has a fixed number of possible values, each with a distinct name.
// Go doesn’t have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.

package main

import (
	"fmt"
)

type ServerState int 

const (
	StateId = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string {
	StateId: "idle",
	StateConnected: "connected",
	StateError: "error",
	StateRetrying: "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateId)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateId:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateId
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

