package main

import (
	"fmt"
	"gonum.org/v1/gonum/graph/formats/dot"
	"gonum.org/v1/gonum/graph/formats/dot/ast"
	"strings"
)

type FsmDefinition struct {
	Transitions []Transition
	States      StateSet
}

type Transition struct {
	Name string
	From *State
	To   *State
}

type State struct {
	Name          string
	PreviousState *State
}

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

type States map[string]bool
type Transitions map[string]bool

func (s States) Add(state string) {
	s[state] = true
}

func (s States) List() []string {
	var list []string
	for k, _ := range s {
		list = append(list, k)
	}
	return list
}

func (t Transitions) Add(transition string) {
	t[transition] = true
}

func (t Transitions) List() []string {
	var list []string
	for k, _ := range t {
		list = append(list, k)
	}
	return list
}

func main() {
	dot, err := dot.ParseFile("./testdata/complex.dot")
	if err != nil {
		fmt.Println(err)
	}

	states := &States{}
	transitions := &Transitions{}
	initial_transition := "initial"

	graphs := dot.Graphs
	fmt.Printf("Found %d graph(s)\n", len(graphs))

	for _, graph := range graphs {
		if !graph.Directed {
			fmt.Println("WARNING: graph should be directed to be able to form a correct Finite State Machine")
		}

		if !graph.Strict {
			fmt.Println("WARNING: graph is not strict. We'll try our best to form a correct Finite State Machine")
		}

		if graph.ID != "" {
			fmt.Printf("It has been named %s\n", graph.ID)
		}

		for _, stmt := range graph.Stmts {
			_, ok := stmt.(*ast.Attr)
			if ok {
				continue
			}
			_, ok = stmt.(*ast.AttrStmt)
			if ok {
				continue
			}

			// Ignore node definitions. We actually only care about nodes that have state transitions
			//node, ok := stmt.(*ast.NodeStmt)
			//if ok {
			//	fmt.Printf("Node Statement: STATE=%s\n", strings.Trim(node.Node.ID, " \"'"))
			//	continue
			//}

			edge, ok := stmt.(*ast.EdgeStmt)
			if ok {
				left := edge.From
				right := edge.To.Vertex
				transition := "unlabeled"

				for _, attr := range edge.Attrs {
					if attr.Key == "label" {
						// strip space and quotes
						transition = strings.Trim(attr.Val, " \"'")
					}
				}

				fmt.Printf("Edge Statement: TRANSITION=%s\n", transition)
				if transition == initial_transition {
					fmt.Println("Skipping initial transition, although we will need to find out the right side node so we know what our initial state is")
					continue
				}

				transitions.Add(transition)

				// Is the left side a node or subgraph?
				l_node, ok := left.(*ast.Node)
				if ok {
					fmt.Println("\tLeft Node: ", l_node.ID)
					states.Add(l_node.ID)
				} else {
					l_sub := left.(*ast.Subgraph)
					for i, stmt2 := range l_sub.Stmts {
						l_node_stmt, ok := stmt2.(*ast.NodeStmt)
						if ok {
							fmt.Printf("\tLeft Node %d: %s\n", i, l_node_stmt.Node.ID)
							states.Add(l_node_stmt.Node.ID)
						} else {
							fmt.Printf("\tLeft something else")
						}
					}
				}

				// If the right a node or subgraph?
				r_node, ok := right.(*ast.Node)
				if ok {
					fmt.Println("\tRight Node: ", r_node.ID)
					states.Add(r_node.ID)
				} else {
					r_sub := right.(*ast.Subgraph)
					for i, stmt2 := range r_sub.Stmts {
						r_node_stmt, ok := stmt2.(*ast.NodeStmt)
						if ok {
							fmt.Printf("\tRight Node %d: %s\n", i, r_node_stmt.Node.ID)
							states.Add(r_node_stmt.Node.ID)
						} else {
							fmt.Printf("\tRight something else")
						}
					}
				}
				continue
			}

			// sub, ok := edge.To.Vertex.(*ast.Subgraph)
			// if ok {
			// 	//fmt.Println("To sub: ", sub.String())
			// 	for _, s := range sub.Stmts {
			// 		fmt.Println("\tTo: ", s.String())
			// 	}
			// } else {
			// 	fmt.Println("To: ", edge.To.Vertex.String())
			// }
		}
	}

	fmt.Println("States:", states.List())
	fmt.Println("transitions:", transitions.List())
}
