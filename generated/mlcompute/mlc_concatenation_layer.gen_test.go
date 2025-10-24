// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCConcatenationLayer

// ExampleNewCConcatenationLayer demonstrates how to create a CConcatenationLayer instance using NewCConcatenationLayer.
// Creates a concatenation layer with a dimension value of 1, which typically represents feature channels.
func ExampleNewCConcatenationLayer() {
	_ = mlcompute.NewCConcatenationLayer()
	// Output:
}
// ExampleNewCConcatenationLayerWithDimension demonstrates how to create a CConcatenationLayer instance using NewCConcatenationLayerWithDimension.
// Creates a concatenation layer with the dimension you specify.
func ExampleNewCConcatenationLayerWithDimension() {
	_ = mlcompute.NewCConcatenationLayerWithDimension(
		0, // dimension uint
	)
	// Output:
}
