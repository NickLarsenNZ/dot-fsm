package fsm

import "github.com/nicklarsennz/dot-fsm/utils"

// This is our internal structure for keeping track of transitions from many states to a single state
type Transition struct {
	ID          string
	Description string
	From        []*State
	To          *State
	states      StateSet
}

func (t *Transition) AddFromState(from *State) {
	t.states.Add(from)
	t.From = append(t.From, from)
}

// Return pointers of all unique states connected to this transition
func (t Transition) States() []*State {
	return t.states.List()
}

func (t Transition) IndividualTransitions() []IndividualTransition {
	panic("Not implemented")
}

// This represents any transition between a pair of states
type IndividualTransition struct {
	ID          string
	Description string
	From        string
	To          string
}

func NewTransition(name string, from []*State, to *State) *Transition {
	stateset := StateSet{}
	for _, state := range from {
		stateset.Add(state)
	}
	stateset.Add(to)

	return &Transition{
		ID:          utils.TextToIdentifier(name),
		Description: name,
		From:        from,
		To:          to,
		states:      stateset,
	}
}
