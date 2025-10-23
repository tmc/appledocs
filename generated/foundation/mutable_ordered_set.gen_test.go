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
// ExampleNewMutableOrderedSetWithCoder demonstrates how to create a MutableOrderedSet instance using NewMutableOrderedSetWithCoder.
func ExampleNewMutableOrderedSetWithCoder() {
	_ = foundation.NewMutableOrderedSetWithCoder(
		foundation.Coder{}, // coder Coder
	)
	// Output:
}
