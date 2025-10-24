// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCSplitLayer

// ExampleNewCSplitLayerWithSplitCountDimension demonstrates how to create a CSplitLayer instance using NewCSplitLayerWithSplitCountDimension.
// Creates a split layer with the number of splits and dimension you specify.
func ExampleNewCSplitLayerWithSplitCountDimension() {
	_ = mlcompute.NewCSplitLayerWithSplitCountDimension(
		10, // splitCount uint
		0, // dimension uint
	)
	// Output:
}
