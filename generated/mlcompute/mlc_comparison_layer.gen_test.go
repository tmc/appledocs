// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCComparisonLayer

// ExampleNewCComparisonLayerWithOperation demonstrates how to create a CComparisonLayer instance using NewCComparisonLayerWithOperation.
// Creates a comparison layer with the operation you specify.
func ExampleNewCComparisonLayerWithOperation() {
	_ = mlcompute.NewCComparisonLayerWithOperation(
		mlcompute.CComparisonOperation{}, // operation CComparisonOperation
	)
	// Output:
}
