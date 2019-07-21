package utils_test

import (
	"github.com/nicklarsennz/dot-fsm/utils"
	"testing"
)

func TestIdentifierFromText(t *testing.T) {
	text := "this is an id for #5"

	expected := "THIS_IS_AN_ID_FOR_5"
	actual := utils.TextToIdentifier(text)

	if actual != expected {
		t.Fatalf("expected identifier '%s', got '%s' instead", expected, actual)
	}
}
