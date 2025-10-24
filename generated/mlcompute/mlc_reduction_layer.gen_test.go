// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCReductionLayer

// ExampleNewCReductionLayerWithReductionTypeDimension demonstrates how to create a CReductionLayer instance using NewCReductionLayerWithReductionTypeDimension.
// Creates a reduction layer using the reduction type and dimension you specify.
func ExampleNewCReductionLayerWithReductionTypeDimension() {
	_ = mlcompute.NewCReductionLayerWithReductionTypeDimension(
		mlcompute.CReductionType{}, // reductionType CReductionType
		0, // dimension uint
	)
	// Output:
}
