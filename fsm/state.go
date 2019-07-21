package fsm

type State struct {
	ID          string
	Description string
}

type StateSet map[*State]bool

func (s StateSet) List() []*State {
	var list []*State
	for k, _ := range s {
		list = append(list, k)
	}
	return list
}

func (s StateSet) Add(state *State) {
	s[state] = true
}
