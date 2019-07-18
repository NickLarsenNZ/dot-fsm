package fsm

type FsmDefinition struct {
	Transitions  []Transition
	initialState *State
	states       StateSet
}

func (f *FsmDefinition) InitialState() *State {
	return f.initialState
}

func (f *FsmDefinition) SetInitialState(state *State) {
	f.states.Add(state)
	f.initialState = state
}
