package fsm

import (
	"github.com/looplab/fsm"
)

var transitions = []fsm.EventDesc{
	{{- range .Transitions }}
	{
		Name: transition.{{.ID}},
		Src: []string{
			{{- range .From}}
			state.{{.ID}},
			{{- end}}
		},
		Dst: state.{{.To.ID}},
	},
	{{- end}}
}

func Fsm(callbacks fsm.Callbacks) *fsm.FSM {
	return fsm.NewFSM(
	    {{- if .InitialState }}
		state.{{.InitialState}}, // Initial State
		{{- end }}
		transitions,
        callbacks,
	)
}