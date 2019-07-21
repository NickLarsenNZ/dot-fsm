package fsm

import (
	"fmt"
	"github.com/nicklarsennz/dot-fsm/utils"
)

type Fsm interface {
	//InitialState() string
	//SetInitialState()
	Transitions() []*Transition
	CreateTransition(name string, from string, to string) error
	TransitionExists(name string, from string, to string) bool
	FindTransitionById(id string) *Transition
	FindStateById(id string) *State
}

type fsm struct {
	transitions  []*Transition
	initialState string
	//states       StateSet
}

func (f fsm) Transitions() []*Transition {
	return f.transitions
}

func (f *fsm) CreateTransition(name string, from string, to string) error {
	var fromState, toState *State

	// Lookup or create new From state
	fromState = f.FindStateById(utils.TextToIdentifier(from))
	if fromState == nil {
		fromState = &State{
			ID:          utils.TextToIdentifier(from),
			Description: from,
		}
	}

	// Lookup or create new To state
	toState = f.FindStateById(utils.TextToIdentifier(to))
	if toState == nil {
		toState = &State{
			ID:          utils.TextToIdentifier(to),
			Description: to,
		}
	}

	// Lookup or create new Transition
	transition := f.FindTransitionById(utils.TextToIdentifier(name))
	if transition == nil {
		transition := NewTransition(name, []*State{fromState}, toState)
		f.transitions = append(f.transitions, transition)

		return nil
	}

	// Ensure the transition has the same To as "to", otherwise it is invalid
	if transition.To.ID == utils.TextToIdentifier(to) {
		// skip if the same transition has been defined
		if f.TransitionExists(name, from, to) {
			fmt.Printf("warning: the transition '%s' from '%s' to '%s' has already been declared\n", name, from, to)
		} else {
			transition.AddFromState(fromState)
		}
	} else {
		return fmt.Errorf("unable to add the transition '%s' to the state '%s', because the transition already exists to state '%s'", name, to, transition.To.Description)
	}

	return nil
}

func (f fsm) FindTransitionById(id string) *Transition {
	for _, transition := range f.transitions {
		if transition.ID == id {
			return transition
		}
	}
	return nil
}

func (f fsm) FindStateById(id string) *State {
	for _, transition := range f.transitions {
		for _, state := range transition.States() {
			if state.ID == id {
				return state
			}
		}
	}
	return nil
}

func (f fsm) TransitionExists(name string, from string, to string) bool {
	for _, transition := range f.transitions {
		if transition.ID == utils.TextToIdentifier(name) {
			if transition.To.ID == utils.TextToIdentifier(to) {
				for _, state := range transition.From {
					if state.ID == utils.TextToIdentifier(from) {
						return true
					}
				}
			}
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
