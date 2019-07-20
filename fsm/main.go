package fsm

import "fmt"
import "github.com/google/go-cmp/cmp"

type Fsm interface {
	//InitialState() string
	//SetInitialState()
	Transitions() []Transition
	CreateTransition(name string, from string, to string) error
	TransitionNameExists(name string) bool
	TransitionExists(lookup Transition) bool
	TransitionExistsWithDestination(name string, destination string) bool
	//get
}

type fsm struct {
	transitions  []Transition
	initialState string
	//states       StateSet
}

func (f fsm) Transitions() []Transition {
	return f.transitions
}

func (f *fsm) CreateTransition(name string, from string, to string) error {

	transition := Transition{
		Name: name,
		From: from,
		To:   to,
	}

	// Is it an exact duplicate
	if f.TransitionExists(transition) {
		fmt.Printf("warning: a duplicate transition '%s' has been declared\n", name)
		return nil
	}

	// Is the transition name already used, and if so, are we trying to create a transition with a different destination?
	if f.TransitionNameExists(name) && f.TransitionExistsWithDestination(name, to) {
		return fmt.Errorf("the transition '%s' already exists with a different destination state", name)
	}

	f.transitions = append(f.transitions, transition)

	return nil
}

func (f fsm) TransitionNameExists(name string) bool {
	for _, t := range f.transitions {
		if t.Name == name {
			return true
		}
	}

	return false
}

func (f fsm) TransitionExists(lookup Transition) bool {
	for _, existing := range f.transitions {
		if cmp.Equal(&existing, &lookup) {
			return true
		}
	}

	return false
}

func (f fsm) TransitionExistsWithDestination(name string, destination string) bool {
	for _, t := range f.transitions {
		if t.Name == name && t.To != destination {
			return true
		}
	}

	return false
}

//func (f *fsm) InitialState() *State {
//	return f.initialState
//}
//
//func (f *fsm) SetInitialState(state *State) {
//	f.states.Add(state)
//	f.initialState = state
//}

func NewFsm() Fsm {
	return &fsm{}
}
