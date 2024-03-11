package revertable

import (
	"github.com/telenornorway/mdc/internal/local"
	val "github.com/telenornorway/mdc/internal/value"
)

type Inner struct {
	differences map[string]*val.Diff
	reverted    bool
}

func newInner(differences map[string]*val.Diff) *Inner {
	return &Inner{
		differences: differences,
		reverted:    false,
	}
}

func (i *Inner) Revert() {
	if i.reverted {
		panic("revertable changes can only be reverted once")
	}
	i.reverted = true

	copy := local.Copy()

	for key, diff := range i.differences {
		nowValue, nowExists := copy[key]

		if diff.New.Exists() != nowExists || diff.New.Value() != nowValue {
			// Revertable changes are not consistent with current state
			//
			// We don't want to restore the original MDC value
			// when the current MDC value is not equal to the
			// MDC value that was set in this snapshot. That
			// means that another part of the application took
			// over the MDC key, and thus that part of application
			// should clean up after themselves. This is a feature
			// and not a bug. We don't want to accidentally break
			// the application if the application expects the
			// current value to remain the same after this snapshot
			// is restored.
			continue
		}

		if diff.Old.Exists() {
			local.Put(key, diff.Old.Value())
		} else {
			local.Remove(key)
		}
	}
}
