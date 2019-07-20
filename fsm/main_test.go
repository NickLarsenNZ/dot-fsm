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
		t.Fatalf("expected only %d transition(s), got %d transition(s)", expectedTransitions, actualTransitions)
	}
}

// If transition name exists, and from/to states differ from the original, error
func TestSameNameTransition(t *testing.T) {
	machine := fsm.NewFsm()

	err1 := machine.CreateTransition("t1", "a", "b")
	err2 := machine.CreateTransition("t1", "c", "d")

	if err1 != nil {
		t.Fatalf("unexpected error creating first transition")
	}

	if err2 == nil {
		t.Fatalf("expected an error on same name transition, got nil instead")
	}
}

// If transition name exists, and from/to states match the original transition, no error (log warning: duplicate transition)
func TestDuplicateTransition(t *testing.T) {
	machine := fsm.NewFsm()
	expectedTransitions := 1

	err1 := machine.CreateTransition("t1", "a", "b")
	err2 := machine.CreateTransition("t1", "a", "b")

	if err1 != nil || err2 != nil {
		t.Fatalf("unexpected error creating duplicate transitions")
	}

	actualTransitions := len(machine.Transitions())
	if expectedTransitions != actualTransitions {
		t.Fatalf("expected only %d transition(s), got %d transition(s)", expectedTransitions, actualTransitions)
	}
}
