package parser

import (
	"fmt"
	"github.com/gonum/graph/formats/dot"
	"github.com/gonum/graph/formats/dot/ast"
	"github.com/nicklarsennz/dot-fsm/fsm"
	"io"
	"strings"
)

func ParseDotFile(handle io.Reader) (fsm.Fsm, error) {
	parsed, err := dot.Parse(handle)
	if err != nil {
		return nil, err
	}

	//f := fsm.NewFsm()

	// For now we'll only concern ourselves with the first graph in the file and warn if there are more
	graphs := parsed.Graphs
	if len(graphs) != 1 {
		fmt.Printf("currently only one graph is supported, selecting 1/%d\n", len(graphs))
	}
	graph := graphs[0]

	if !graph.Directed {
		return nil, fmt.Errorf("graph must be a digraph to be able to form a correct Finite State Machine")
	}

	// produce an FSM
	return ProduceFSM(graph)
}

func ProduceFSM(digraph *ast.Graph) (fsm.Fsm, error) {

	machine := fsm.NewFsm()
	transition_count := 0

	for _, stmt := range digraph.Stmts {

		switch t := stmt.(type) {
		case *ast.EdgeStmt:
			transition_count++
			transition_name := fmt.Sprintf("transition_%d", transition_count)

			for _, attr := range t.Attrs {
				if attr.Key == "label" {
					// strip space and quotes
					transition_name = strings.Trim(attr.Val, " \"'")
				}
			}

			left := t.From
			right := t.To.Vertex

			// Fail if the Right side is not a Node (eg: subgraph) which makes
			to, ok := right.(*ast.Node)
			if !ok {
				return nil, fmt.Errorf("a transition cannot have multiple destination states: %s", t.String())
			}

			switch from := left.(type) {
			case *ast.Node:
				machine.CreateTransition(transition_name, from.ID, to.ID)
			case *ast.Subgraph:
				for _, sub_stmt := range from.Stmts {
					switch from := sub_stmt.(type) {
					case *ast.NodeStmt:
						machine.CreateTransition(transition_name, from.Node.ID, to.ID)
					default:
						fmt.Printf("unexpected statement in subgraph of from states for transition %s: %s\n", transition_name, t.String())
					}
				}
			}
		}
	}

	return machine, nil
}
