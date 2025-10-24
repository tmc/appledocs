// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewMatrixMultiplication

// ExampleMatrixMultiplication_Encode demonstrates using Encode on a MatrixMultiplication instance.
// Encodes a matrix multiplication kernel to a command buffer.
func ExampleMatrixMultiplication_Encode() {
	obj := metalperformanceshaders.NewMatrixMultiplication()
	obj.Encode()
	// Output:
	}

