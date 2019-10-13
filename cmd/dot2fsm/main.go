package main

import (
	"fmt"
	"github.com/nicklarsennz/dot-fsm/fsm"
	"github.com/nicklarsennz/dot-fsm/gen"
	"github.com/nicklarsennz/dot-fsm/parser"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage:\n%s <dot-file> <template-path>\n", os.Args[0])
		os.Exit(1)
	}

	dot_file := os.Args[1]
	template_dir := os.Args[2]

	file, err := os.Stat(dot_file)
	exitOnError(err)
	if !file.Mode().IsRegular() {
		fmt.Printf("%s is not a file\n", dot_file)
		os.Exit(1)
	}

	file, err = os.Stat(template_dir)
	exitOnError(err)
	if !file.Mode().IsDir() {
		fmt.Printf("%s is not a directory\n", template_dir)
		os.Exit(1)
	}

	fmt.Printf("file: %s\ntemplate: %s\n", dot_file, template_dir)

	dot, err := os.Open(dot_file)
	exitOnError(err)
	fsm, err := parser.ParseDotFile(dot)
	exitOnError(err)

	model := convertToSimple(fsm)

	// TODO: remove the two lines below
	fixed_file := "./templates/looplab-fsm/fsm.go.tpl"
	fixed_template, _ := os.Open(fixed_file)

	// TODO: for each file in the path, drop the template_path from it and write to an output directory:
	r := gen.NewRenderer(fixed_template, model)
	r.PipeTo("gofmt") // TODO: user will specify pipe commands
	fmt.Fprintf(os.Stderr, "rendering file %s\n", fixed_file)
	r.Render(os.Stdout)

}

func convertToSimple(fsm fsm.Fsm) *model_data {
	data := &model_data{}
	// data.InitialState = fsm.InitialState()

	for _, fsm_transition := range fsm.Transitions() {
		transition := model_transition{}

		transition.ID = fsm_transition.ID
		transition.Description = fsm_transition.Description

		from := []model_state{}
		for _, from_state := range fsm_transition.From {
			f := model_state{}
			f.ID = from_state.ID
			f.Description = from_state.Description
			from = append(from, f)
		}
		transition.From = from

		to := model_state{}
		to.ID = fsm_transition.To.ID
		to.Description = fsm_transition.To.Description
		transition.To = to

		data.Transitions = append(data.Transitions, transition)
	}

	return data
}

func exitOnError(err error) {
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
