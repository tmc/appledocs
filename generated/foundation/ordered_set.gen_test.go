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
// ExampleNewOrderedSetWithCoder demonstrates how to create a OrderedSet instance using NewOrderedSetWithCoder.
func ExampleNewOrderedSetWithCoder() {
	_ = foundation.NewOrderedSetWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
