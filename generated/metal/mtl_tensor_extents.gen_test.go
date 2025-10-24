// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal_test

import (
	"github.com/tmc/appledocs/generated/metal"
)

// Suppress unused import errors
var _ = metal.NewTensorExtents

// ExampleNewTensorExtentsWithRankValues demonstrates how to create a TensorExtents instance using NewTensorExtentsWithRankValues.
// Creates a new tensor extents with the rank and extent values you provide.
func ExampleNewTensorExtentsWithRankValues() {
	_ = metal.NewTensorExtentsWithRankValues(
		0, // rank uint
		0, // values int
	)
	// Output:
}
