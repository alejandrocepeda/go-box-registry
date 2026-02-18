package box

import "fmt"

type BoxFactory func() Box

var factories = make(map[string]BoxFactory)

func RegisterBox(id string, factory BoxFactory) {
	if _, exists := factories[id]; exists {
		panic(fmt.Sprintf("box: duplicate registration for %q", id))
	}
	factories[id] = factory
}

type Registry struct {
	boxes map[string]Box
}

func BuildAll() Registry {
	reg := Registry{boxes: make(map[string]Box)}
	for id, factory := range factories {
		b := factory()
		b.ID = id
		reg.boxes[id] = b
	}
	return reg
}

func (r *Registry) Get(id string) (Box, error) {
	b, ok := r.boxes[id]
	if !ok {
		return Box{}, fmt.Errorf("box: %q not found in registry", id)
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
