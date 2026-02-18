package box

import "reflect"

type Result struct {
	Output   map[string]any `json:"output"`
	NewEvent string         `json:"new_event,omitempty"`
}

type Deps map[reflect.Type]any

func NewDeps(values ...any) Deps {
	d := Deps{}
	for _, v := range values {
		d[reflect.TypeOf(v)] = v
	}
	return d
}

func Get[T any](d Deps) T {
	t := reflect.TypeOf(new(T)).Elem()
	return d[t].(T)
}

type Executor interface {
	Dependencies() []reflect.Type
	Apply(d Deps) (Result, error)
}

type Box struct {
	ID string
	Executor
}
