// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCDropoutLayer

// ExampleNewCDropoutLayerWithRateSeed demonstrates how to create a CDropoutLayer instance using NewCDropoutLayerWithRateSeed.
// Creates a dropout layer with the probability rate and random number generator seed you specify.
func ExampleNewCDropoutLayerWithRateSeed() {
	_ = mlcompute.NewCDropoutLayerWithRateSeed(
		0.0, // rate float32
		0, // seed uint
	)
	// Output:
}
