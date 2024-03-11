package mdc

import "github.com/telenornorway/mdc/internal/local"

// Get returns a boolean indicating if the key exists and the value of the key if it exists.
//
//goland:noinspection GoUnusedExportedFunction
func Get(key string) (bool, string) { return local.Get(key) }

// Put adds or updates the value of the MDC.
func Put(key, value string) { local.Put(key, value) }

// Remove removes the key from the MDC.
//
//goland:noinspection GoUnusedExportedFunction
func Remove(key string) { local.Remove(key) }

// Clear removes all keys from the MDC.
//
//goland:noinspection GoUnusedExportedFunction
func Clear() { local.Clear() }

// Copy returns a copy of the MDC.
func Copy() map[string]string { return local.Copy() }
