package transition

const (
    {{- range .Transitions}}
	{{.ID}} = "{{.Name}}"
    {{- end}}
)