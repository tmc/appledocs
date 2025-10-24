// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCGramMatrixLayer

// ExampleNewCGramMatrixLayerWithScale demonstrates how to create a CGramMatrixLayer instance using NewCGramMatrixLayerWithScale.
// Creates a gram matrix layer with the scaling factor you specify.
func ExampleNewCGramMatrixLayerWithScale() {
	_ = mlcompute.NewCGramMatrixLayerWithScale(
		1.0, // scale float32
	)
	// Output:
}
