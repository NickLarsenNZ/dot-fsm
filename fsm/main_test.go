package fsm_test

import (
	"github.com/nicklarsennz/dot-fsm/fsm"
	"testing"
)

func TestTransitionCreatesStates(t *testing.T) {
	machine := fsm.NewFsm

	// machine.CreateTransition("a", "b")
	// machine.States == ["a", "b"]
}

func TestTransitionSetsPreviousState(t *testing.T) {
	machine := fsm.NewFsm

	// machine.CreateTransition("a", "b")
	//
}
