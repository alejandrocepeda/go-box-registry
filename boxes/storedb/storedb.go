package boxes

import (
	"fmt"
	"payment-rewrite/box"
	"payment-rewrite/deps"
	"reflect"
)

func init() {
	box.RegisterBox("store_db", func() box.Box {
		return box.Box{Executor: Box{}}
	})
}

type Box struct{}

func (b Box) Dependencies() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf(deps.DB{}),
	}
}

func (b Box) Apply(d box.Deps) (box.Result, error) {
	db := box.Get[deps.DB](d)

	fmt.Printf("Store using DB at %s\n", db)

	return box.Result{
		Output: map[string]any{"stored": true, "capture_id": ""},
	}, nil
}
