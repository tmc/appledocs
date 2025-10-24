// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMutableOrderedSet

// ExampleNewMutableOrderedSet demonstrates how to create a MutableOrderedSet instance.
// Initializes a newly allocated mutable ordered set.
func ExampleNewMutableOrderedSet() {
	_ = foundation.NewMutableOrderedSet()
	// Output:
}
// ExampleNewMutableOrderedSetWithCapacity demonstrates how to create a MutableOrderedSet instance using NewMutableOrderedSetWithCapacity.
// Returns an initialized mutable ordered set with a given initial capacity.
func ExampleNewMutableOrderedSetWithCapacity() {
	_ = foundation.NewMutableOrderedSetWithCapacity(
		0, // numItems uint
	)
	// Output:
}
// ExampleMutableOrderedSet_RemoveAllObjects demonstrates using RemoveAllObjects on a MutableOrderedSet instance.
// Removes all the objects from the mutable ordered set.
func ExampleMutableOrderedSet_RemoveAllObjects() {
	obj := foundation.NewMutableOrderedSet()
	obj.RemoveAllObjects()
	// Output:
	}

