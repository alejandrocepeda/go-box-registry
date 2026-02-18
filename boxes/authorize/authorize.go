package authorize

import (
	"fmt"
	"payment-rewrite/box"
	"payment-rewrite/deps"
	"reflect"
)

func init() {
	box.RegisterBox("authorize", func() box.Box {
		return box.Box{Executor: Box{}}
	})
}

type Box struct{}

func (b Box) Dependencies() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf(deps.Job{}),
	}
}

func (b Box) Apply(d []any) (box.Result, error) {
	job := d[0].(deps.Job)

	fmt.Printf("Authorizing using job %s\n", job)

	return box.Result{
		Output: map[string]any{"authorize_id": "auth_001", "amount": 0},
	}, nil
}
