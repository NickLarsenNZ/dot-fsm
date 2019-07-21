package state

const (
    {{- range .States}}
	{{.ID}} = "{{.Name}}"
    {{- end}}
)