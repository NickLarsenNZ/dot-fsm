package gen_test

import (
	"bytes"
	"github.com/nicklarsennz/dot-fsm/gen"
	"strings"
	"testing"
)

func TestTemplateRender(t *testing.T) {
	template := `
{{range .Transitions -}}
{{.Description}}
{{end -}}
`
	expected := `
t1
t2
`
	templateData := struct {
		Transitions []struct {
			Name string
		}
	}{
		Transitions: []struct{ Name string }{
			{Name: "t1"},
			{Name: "t2"},
		},
	}

	byteWriter := bytes.NewBuffer([]byte{})
	stringReader := strings.NewReader(template)

	renderer := gen.NewRenderer(stringReader, templateData)
	err := renderer.Render(byteWriter)
	if err != nil {
		t.Fatal(err)
	}

	actual := byteWriter.String()

	if actual != expected {
		t.Fatalf("rendering the template failed. Expected '%s', got '%s'", expected, actual)
	}
}

func TestTemplateRenderWithCommands(t *testing.T) {
	template := `


 package gen

   func {{.Thing}}()     {
_ = 0
 }


`
	expected := `package gen

func blah() {
	_ = 0
}
`
	templateData := struct {
		Thing string
	}{
		Thing: "blah",
	}

	byteWriter := bytes.NewBuffer([]byte{})
	stringReader := strings.NewReader(template)

	renderer := gen.NewRenderer(stringReader, templateData)
	renderer.PipeTo("gofmt")
	err := renderer.Render(byteWriter)
	if err != nil {
		t.Fatal(err)
	}

	actual := byteWriter.String()

	if actual != expected {
		t.Fatalf("rendering the template failed. Expected '%s', got '%s'", expected, actual)
	}
}
