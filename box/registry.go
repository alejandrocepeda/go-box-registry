package box

import "fmt"

type BoxFactory func(deps Deps) Box

var factories = make(map[string]BoxFactory)

// RegisterBox reads the ID from the Box itself.
// The ID is defined once: in the struct's ID() method.
func RegisterBox(factory BoxFactory) {
	dummy := factory(Deps{})
	id := dummy.ID()
	if _, exists := factories[id]; exists {
		panic(fmt.Sprintf("box: duplicate registration for %q", id))
	}
	factories[id] = factory
}

type Registry struct {
	boxes map[string]Box
}

func BuildAll(deps Deps) Registry {
	reg := Registry{boxes: make(map[string]Box)}
	for id, factory := range factories {
		reg.boxes[id] = factory(deps)
	}
	return reg
}

func (r *Registry) Get(id string) (Box, error) {
	b, ok := r.boxes[id]
	if !ok {
		return nil, fmt.Errorf("box: %q not found in registry", id)
	}
	return b, nil
}

func (r *Registry) IDs() []string {
	ids := make([]string, 0, len(r.boxes))
	for id := range r.boxes {
		ids = append(ids, id)
	}
	return ids
}
