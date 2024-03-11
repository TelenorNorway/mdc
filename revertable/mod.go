package revertable

import ir "github.com/telenornorway/mdc/internal/revertable"

type Holder struct {
	inner *ir.Inner
}

func (h Holder) Revert() {
	h.inner.Revert()
}

type Builder struct {
	inner *ir.Builder
}

func (b Builder) Put(key, value string) Builder {
	b.inner.Put(key, value)
	return b
}

func (b Builder) PutIfNotEmpty(key, value string) Builder {
	b.inner.PutIfNotEmpty(key, value)
	return b
}

func (b Builder) Remove(key string) Builder {
	b.inner.Remove(key)
	return b
}

func (b Builder) Clear() Builder {
	b.inner.Clear()
	return b
}

func (b Builder) Apply() Holder {
	return Holder{inner: b.inner.Apply()}
}

func New() Builder {
	return Builder{inner: ir.NewRevertableBuilder()}
}

func Put(key, value string) Builder {
	return New().Put(key, value)
}

func PutIfNotEmpty(key, value string) Builder {
	return New().PutIfNotEmpty(key, value)
}

func Remove(key string) Builder {
	return New().Remove(key)
}

func Clear() Builder {
	return New().Clear()
}
