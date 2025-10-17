// Package-level convenience helpers for ObjectiveC.
//
// This file contains hand-written helpers that complement the generated bindings.
// It is not overwritten during code generation.
package objectivec

import (
	"github.com/tmc/appledocs/generated/objc"
)

// GetID returns the underlying objc.ID for the Object.
// This method makes Object satisfy the objc.IObject interface.
func (o Object) GetID() objc.ID {
	return o.ID
}
