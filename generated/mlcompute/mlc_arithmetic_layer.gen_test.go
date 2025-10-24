// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCArithmeticLayer

// ExampleNewCArithmeticLayerWithOperation demonstrates how to create a CArithmeticLayer instance using NewCArithmeticLayerWithOperation.
// Creates an arithmetic layer with the operation you specify.
func ExampleNewCArithmeticLayerWithOperation() {
	_ = mlcompute.NewCArithmeticLayerWithOperation(
		mlcompute.CArithmeticOperation{}, // operation CArithmeticOperation
	)
	// Output:
}
