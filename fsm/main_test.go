package fsm_test

import (
	"github.com/nicklarsennz/dot-fsm/fsm"
	"testing"
)

func TestCreateTransition(t *testing.T) {
	machine := fsm.NewFsm()
	expectedTransitions := 1

	err := machine.CreateTransition("t1", "a", "b")
	if err != nil {
		t.Fatalf("unexpected error creating first transition")
	}

	actualTransitions := len(machine.Transitions())
	if expectedTransitions != actualTransitions {
		t.Fatalf("expected %d transition(s), got %d transition(s)", expectedTransitions, actualTransitions)
	}
}

func TestTransitionWithDifferentDestination(t *testing.T) {
	machine := fsm.NewFsm()

	// a -> b
	// c -> d
	err1 := machine.CreateTransition("t1", "a", "b")
	err2 := machine.CreateTransition("t1", "c", "d")

	if err1 != nil {
		t.Fatalf("unexpected error creating first transition")
	}

	if err2 == nil {
		t.Fatalf("expected an error on same name transition, got nil instead")
	}
}

func TestTransitionWithSameDestination(t *testing.T) {
	machine := fsm.NewFsm()
	expectedTransitions := 2

	// {a b} -> c
	err1 := machine.CreateTransition("t1", "a", "c")
	err2 := machine.CreateTransition("t1", "b", "c")

	if err1 != nil || err2 != nil {
		t.Fatalf("unexpected error creating transitions")
	}

	actualTransitions := len(machine.Transitions())
	if expectedTransitions != actualTransitions {
		t.Fatalf("expected %d transition(s), got %d transition(s)", expectedTransitions, actualTransitions)
	}
}

func TestDuplicateTransition(t *testing.T) {
	machine := fsm.NewFsm()
	expectedTransitions := 1

	// a -> b
	// a -> b
	err1 := machine.CreateTransition("t1", "a", "b")
	err2 := machine.CreateTransition("t1", "a", "b")

	if err1 != nil || err2 != nil {
		t.Fatalf("unexpected error creating duplicate transitions")
	}

	actualTransitions := len(machine.Transitions())
	if expectedTransitions != actualTransitions {
		t.Fatalf("expected %d transition(s), got %d transition(s)", expectedTransitions, actualTransitions)
	}
}
