package gen

import (
	"github.com/joncalhoun/pipe"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"text/template"
)

type Renderer interface {
	Render(buffer io.Writer) error
	PipeTo(command string, arg ...string)
}

type renderer struct {
	template        string
	data            interface{}
	commandPipeline []*exec.Cmd
}

func (r *renderer) PipeTo(command string, arg ...string) {
	r.commandPipeline = append(r.commandPipeline, exec.Command(command, arg...))
}

func (r renderer) Render(buffer io.Writer) error {
	t := template.Must(template.New("fsm").Parse(r.template))

	if len(r.commandPipeline) == 0 {
		// If there are no commands in the pipeline, simply render the template
		return t.Execute(buffer, &r.data)
	} else {
		// Otherwise we need to render, then pipe the output through the commands in the list
		for _, cmd := range r.commandPipeline {
			cmd.Stderr = os.Stderr
		}
		rc, wc, errCh := pipe.Commands(
			r.commandPipeline...,
		)

		go func() {
			select {
			case err, ok := <-errCh:
				if ok && err != nil {
					os.Exit(2)
				}
			}
		}()

		err := t.Execute(wc, &r.data)
		if err != nil {
			return err
		}

		err = wc.Close()
		if err != nil {
			return err
		}

		_, err = io.Copy(buffer, rc)
		if err != nil {
			return err
		}

		return nil
	}
}

func NewRenderer(template io.Reader, data interface{}) Renderer {
	buf, err := ioutil.ReadAll(template)
	if err != nil {
		log.Fatal(err)
	}

	return &renderer{
		template: string(buf),
		data:     data,
	}
}
