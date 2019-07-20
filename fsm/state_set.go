package fsm

type StateSet map[*State]bool

func (s StateSet) Add(state *State) {
	s[state] = true
}

func (s StateSet) List() []*State {
	var list []*State
	for k, _ := range s {
		list = append(list, k)
	}
	return list
}

func (s StateSet) Exists(state string) bool {
	for k, _ := range s {
		if k.Name == state {
			return true
		}
	}
	return false
}
