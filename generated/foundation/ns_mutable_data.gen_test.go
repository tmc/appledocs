// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMutableData

// ExampleNewMutableDataWithCapacity demonstrates how to create a MutableData instance using NewMutableDataWithCapacity.
// Returns an initialized mutable data object capable of holding the specified number of bytes.
func ExampleNewMutableDataWithCapacity() {
	_ = foundation.NewMutableDataWithCapacity(
		0, // capacity uint
	)
	// Output:
}
// ExampleNewMutableDataWithLength demonstrates how to create a MutableData instance using NewMutableDataWithLength.
// Initializes and returns a mutable data object containing a given number of zeroed bytes.
func ExampleNewMutableDataWithLength() {
	_ = foundation.NewMutableDataWithLength(
		10, // length uint
	)
	// Output:
}
