package main

type model_data struct {
	Transitions  []model_transition
	InitialState string
}

type model_transition struct {
	ID          string
	Description string
	From        []model_state
	To          model_state
}

type model_state struct {
	ID          string
	Description string
}
