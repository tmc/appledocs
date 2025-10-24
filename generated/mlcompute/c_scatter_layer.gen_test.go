// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCScatterLayer

// ExampleNewCScatterLayerWithDimensionReductionType demonstrates how to create a CScatterLayer instance using NewCScatterLayerWithDimensionReductionType.
// Creates a scatter layer with the dimension and reduction type you specify.
func ExampleNewCScatterLayerWithDimensionReductionType() {
	_ = mlcompute.NewCScatterLayerWithDimensionReductionType(
		0,                          // dimension uint
		mlcompute.CReductionType{}, // reductionType CReductionType
	)
	// Output:
}
