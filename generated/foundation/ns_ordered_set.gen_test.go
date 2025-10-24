// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewOrderedSet

// ExampleNewOrderedSet demonstrates how to create a OrderedSet instance.
// Initializes a newly allocated ordered set.
func ExampleNewOrderedSet() {
	_ = foundation.NewOrderedSet()
	// Output:
}
// ExampleOrderedSet_ObjectEnumerator demonstrates using ObjectEnumerator on a OrderedSet instance.
// Returns an enumerator object that lets you access each object in the ordered set.
func ExampleOrderedSet_ObjectEnumerator() {
	obj := foundation.NewOrderedSet()
	_ = obj.ObjectEnumerator()
	// Output:
	}

// ExampleOrderedSet_ReverseObjectEnumerator demonstrates using ReverseObjectEnumerator on a OrderedSet instance.
// Returns an enumerator object that lets you access each object in the ordered set.
func ExampleOrderedSet_ReverseObjectEnumerator() {
	obj := foundation.NewOrderedSet()
	_ = obj.ReverseObjectEnumerator()
	// Output:
	}

