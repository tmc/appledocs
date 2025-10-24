// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCMatMulDescriptor

// ExampleNewCMatMulDescriptor demonstrates how to create a CMatMulDescriptor instance using NewCMatMulDescriptor.
// Creates a batched matrix multiplication descriptor.
func ExampleNewCMatMulDescriptor() {
	_ = mlcompute.NewCMatMulDescriptor()
	// Output:
}
// ExampleNewCMatMulDescriptorWithAlphaTransposesXTransposesY demonstrates how to create a CMatMulDescriptor instance using NewCMatMulDescriptorWithAlphaTransposesXTransposesY.
// Creates a batched matrix multiplication descriptor with the alpha value and transpose options you specify.
func ExampleNewCMatMulDescriptorWithAlphaTransposesXTransposesY() {
	_ = mlcompute.NewCMatMulDescriptorWithAlphaTransposesXTransposesY(
		1.0, // alpha float32
		false, // transposesX bool
		false, // transposesY bool
	)
	// Output:
}
