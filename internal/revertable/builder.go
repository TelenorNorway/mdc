package revertable

import (
	"github.com/telenornorway/mdc/internal/local"
	val "github.com/telenornorway/mdc/internal/value"
)

type Builder struct {
	changes        map[string]*val.Value
	hasBeenApplied bool
}

func NewRevertableBuilder() *Builder {
	return &Builder{
		changes: make(map[string]*val.Value),
	}
}

func (r *Builder) Put(key, value string) {
	r.changes[key] = val.Some(value)
}

func (r *Builder) PutIfNotEmpty(key, value string) {
	if value != "" {
		r.changes[key] = val.Some(value)
	}
}

func (r *Builder) Remove(key string) {
	r.changes[key] = val.None
}

func (r *Builder) Clear() {
	copy := local.Copy()
	for k := range copy {
		r.changes[k] = val.None
	}
}

func (r *Builder) Apply() *Inner {
	if r.hasBeenApplied {
		panic("revertable changes can only be applied once")
	}
	r.hasBeenApplied = true

	differences := make(map[string]*val.Diff)

	for k, v := range r.changes {
		differences[k] = val.NewDiff(local.GetAsValue(k), v)
		if v.Exists() {
			local.Put(k, v.Value())
		} else {
			local.Remove(k)
		}
	}

	return newInner(differences)
}
