// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewCountedSet

// ExampleNewCountedSetWithCapacity demonstrates how to create a CountedSet instance using NewCountedSetWithCapacity.
// Returns a counted set object initialized with enough memory to hold a given number of objects.
func ExampleNewCountedSetWithCapacity() {
	_ = foundation.NewCountedSetWithCapacity(
		0, // numItems uint
	)
	// Output:
}
// ExampleCountedSet_ObjectEnumerator demonstrates using ObjectEnumerator on a CountedSet instance.
// Returns an enumerator object that lets you access each object in the set once, independent of its count.
func ExampleCountedSet_ObjectEnumerator() {
	obj := foundation.NewCountedSet()
	_ = obj.ObjectEnumerator()
	// Output:
	}

