// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewPointerArray

// ExampleNewPointerArrayWithOptions demonstrates how to create a PointerArray instance using NewPointerArrayWithOptions.
// Initializes the receiver to use the given options.
func ExampleNewPointerArrayWithOptions() {
	_ = foundation.NewPointerArrayWithOptions(
		foundation.PointerFunctionsOptions{}, // options PointerFunctionsOptions
	)
	// Output:
}
// ExamplePointerArray_Compact demonstrates using Compact on a PointerArray instance.
// Removes   values from the receiver.
func ExamplePointerArray_Compact() {
	obj := foundation.NewPointerArray()
	obj.Compact()
	// Output:
	}

