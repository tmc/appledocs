// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMutableSet

// ExampleNewMutableSet demonstrates how to create a MutableSet instance.
// Initializes a newly allocated set.
func ExampleNewMutableSet() {
	_ = foundation.NewMutableSet()
	// Output:
}

// ExampleNewMutableSetWithCapacity demonstrates how to create a MutableSet instance using NewMutableSetWithCapacity.
// Returns an initialized mutable set with a given initial capacity.
func ExampleNewMutableSetWithCapacity() {
	_ = foundation.NewMutableSetWithCapacity(
		0, // numItems uint
	)
	// Output:
}
