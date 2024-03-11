package local

import (
	"github.com/telenornorway/mdc/internal/map"
	"github.com/telenornorway/mdc/internal/value"
	"github.com/timandy/routine"
)

var localMDC = routine.NewThreadLocalWithInitial[*_map.MDC](_map.NewMdc)

// Get returns the value of a key in the mapped diagnostic context for the current goroutine.
func Get(key string) (bool, string) {
	return localMDC.Get().Get(key)
}

func GetAsValue(key string) *value.Value {
	return value.From(Get(key))
}

// Put puts a key-value pair into the mapped diagnostic context for the current goroutine.
func Put(key, value string) {
	localMDC.Get().Put(key, value)
}

// Remove removes a key-value pair from the mapped diagnostic context for the current goroutine.
func Remove(key string) {
	localMDC.Get().Remove(key)
}

// Clear removes all key-value pairs from the mapped diagnostic context for the current goroutine.
func Clear() {
	localMDC.Get().Clear()
}

// Copy returns a copy of the mapped diagnostic context for the current goroutine.
func Copy() map[string]string {
	return localMDC.Get().Copy()
}
