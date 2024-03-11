package _map

import "sync"

type MDC struct {
	pairs map[string]string

	// In theory this shouldn't be necessary, but it's here to be safe. A
	// consumer of this package could be wrapping goroutines in a way that clones
	// the MDC, and we want to make sure that the MDC is thread-safe in that
	// case.
	lock *sync.Mutex
}

func NewMdc() *MDC {
	return &MDC{
		pairs: map[string]string{},
		lock:  &sync.Mutex{},
	}
}

// Get returns the value of a key in the mapped diagnostic context.
func (m *MDC) Get(key string) (exists bool, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	value, exists = m.pairs[key]
	return
}

// Put puts a key-value pair into the mapped diagnostic context.
func (m *MDC) Put(key, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.pairs[key] = value
}

// Remove removes a key-value pair from the mapped diagnostic context.
func (m *MDC) Remove(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.pairs, key)
}

// Clear removes all key-value pairs from the mapped diagnostic context.
func (m *MDC) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.pairs = map[string]string{}
}

// Copy returns a copy of the mapped diagnostic context.
func (m *MDC) Copy() map[string]string {
	m.lock.Lock()
	defer m.lock.Unlock()
	copy := map[string]string{}
	for k, v := range m.pairs {
		copy[k] = v
	}
	return copy
}
