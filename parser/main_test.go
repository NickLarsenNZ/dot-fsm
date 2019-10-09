package parser_test

import (
	"github.com/nicklarsennz/dot-fsm/parser"
	"strings"
	"testing"
)

const digraph string = `
digraph dummy {
   a
   b

   a -> b
}
`

// https://endler.dev/2018/go-io-testing/
func TestParseDotFile(t *testing.T) {
	expected_transitions := 1

	dummy := strings.NewReader(digraph)

	fsm, err := parser.ParseDotFile(dummy)
	if err != nil {
		t.Fatalf("parsing error: %s", err)
	}

	transitions := fsm.Transitions()
	actual_transitions := len(transitions)
	if actual_transitions != expected_transitions {
		t.Fatalf("wrong transition count: expected %d, got %d", expected_transitions, actual_transitions)
	}
}
