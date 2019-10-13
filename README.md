# dot-fsm
Finite State Machine code generator for converting Dot digraphs to loop/fsm states and events

## Usage

```sh
dot2fsm <dot-file> <template-path>
dot2fsm <dot-file> <template-path> -p 'gofmt' -p 'cat -n'
```

## Core

- cmd/dot2fsm: binary artifact
- fsm: internal representation of an FSM
- gen: take an interface and render a template
- parser: parse dot and return a simple structure
- utils: internal utilities

