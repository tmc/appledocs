// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewUnaryImageKernel

// ExampleUnaryImageKernel_Encode demonstrates using Encode on a UnaryImageKernel instance.
// Encodes a kernel into a command buffer, out of place.
func ExampleUnaryImageKernel_Encode() {
	obj := metalperformanceshaders.NewUnaryImageKernel()
	obj.Encode()
	// Output:
	}


