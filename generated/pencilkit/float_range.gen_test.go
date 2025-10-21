// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit_test

import (
	"github.com/tmc/appledocs/generated/pencilkit"
)

// Suppress unused import errors
var _ = pencilkit.NewFloatRange

// ExampleNewFloatRangeWithLowerBoundUpperBound demonstrates how to create a FloatRange instance using NewFloatRangeWithLowerBoundUpperBound.
// A utility class used to contain ranges returned by the PKStroke API.
func ExampleNewFloatRangeWithLowerBoundUpperBound() {
	_ = pencilkit.NewFloatRangeWithLowerBoundUpperBound(
		0.0, // lowerBound float64
		0.0, // upperBound float64
	)
	// Output:
}
