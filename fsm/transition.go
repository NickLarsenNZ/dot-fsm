package fsm

type Transition struct {
	Name string
	From *State
	To   *State
}
