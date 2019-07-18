package fsm

type State struct {
	Name          string
	PreviousState *State
}
