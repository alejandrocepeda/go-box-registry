package bankart

import (
	"fmt"
	"payment-rewrite/box"
	"payment-rewrite/deps"
	"reflect"
)

func init() {
	box.RegisterBox("capture", func() box.Box {
		return box.Box{Executor: Box{}}
	})
}

type Box struct{}

func (b Box) Dependencies() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf(deps.Job{}),
	}
}

func (b Box) Apply(d box.Deps) (box.Result, error) {
	job := box.Get[deps.Job](d)

	fmt.Printf("Capturing using job %s\n", job)

	return box.Result{
		Output: map[string]any{"capture_id": "cap_123"},
	}, nil
}
