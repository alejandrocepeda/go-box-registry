package box

import "reflect"

type Result struct {
	Output   map[string]any `json:"output"`
	NewEvent string         `json:"new_event,omitempty"`
}

type Executor interface {
	Dependencies() []reflect.Type
	Apply(deps []any) (Result, error)
}

type Box struct {
	ID string
	Executor
}
