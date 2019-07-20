package fsm

import "fmt"
import "github.com/google/go-cmp/cmp"

type Fsm interface {
	//InitialState() string
	//SetInitialState()
	Transitions() []Transition
	CreateTransition(name string, from string, to string) error
	TransitionNameExists(name string) bool
	TransitionWithStatesExists(lookup Transition) bool
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
	if f.TransitionWithStatesExists(transition) {
		fmt.Printf("warning: a duplicate transition '%s' has been declared\n", name)
		return nil
	}

	// Is the transition name already used
	if f.TransitionNameExists(name) {
		return fmt.Errorf("the transition '%s' already exists with different from and to states", name)
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

func (f fsm) TransitionWithStatesExists(lookup Transition) bool {
	for _, existing := range f.transitions {
		if cmp.Equal(&existing, &lookup) {
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
