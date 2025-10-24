// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCSoftmaxLayer

// ExampleNewCSoftmaxLayerWithOperation demonstrates how to create a CSoftmaxLayer instance using NewCSoftmaxLayerWithOperation.
// Creates a softmax layer with the operation you specify.
func ExampleNewCSoftmaxLayerWithOperation() {
	_ = mlcompute.NewCSoftmaxLayerWithOperation(
		mlcompute.CSoftmaxOperation{}, // operation CSoftmaxOperation
	)
	// Output:
}
// ExampleNewCSoftmaxLayerWithOperationDimension demonstrates how to create a CSoftmaxLayer instance using NewCSoftmaxLayerWithOperationDimension.
// Creates a softmax layer with the operation and dimension you specify.
func ExampleNewCSoftmaxLayerWithOperationDimension() {
	_ = mlcompute.NewCSoftmaxLayerWithOperationDimension(
		mlcompute.CSoftmaxOperation{}, // operation CSoftmaxOperation
		0, // dimension uint
	)
	// Output:
}
