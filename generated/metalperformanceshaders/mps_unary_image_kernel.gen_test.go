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

// ExampleUnaryImageKernel_SourceRegion demonstrates using SourceRegion on a UnaryImageKernel instance.
// Determines the region of the source texture that will be read for an encode operation.
func ExampleUnaryImageKernel_SourceRegion() {
	obj := metalperformanceshaders.NewUnaryImageKernel()
	obj.SourceRegion()
	// Output:
	}

