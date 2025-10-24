// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewBinaryImageKernel

// ExampleBinaryImageKernel_Encode demonstrates using Encode on a BinaryImageKernel instance.
// This method attempts to apply a kernel in place on a texture.
func ExampleBinaryImageKernel_Encode() {
	obj := metalperformanceshaders.NewBinaryImageKernel()
	obj.Encode()
	// Output:
	}

// ExampleBinaryImageKernel_SecondarySourceRegion demonstrates using SecondarySourceRegion on a BinaryImageKernel instance.
// Determines the region of the secondary source texture that will be read for an encode operation.
func ExampleBinaryImageKernel_SecondarySourceRegion() {
	obj := metalperformanceshaders.NewBinaryImageKernel()
	obj.SecondarySourceRegion()
	// Output:
	}

// ExampleBinaryImageKernel_PrimarySourceRegion demonstrates using PrimarySourceRegion on a BinaryImageKernel instance.
// Determines the region of the primary source texture that will be read for an encode operation.
func ExampleBinaryImageKernel_PrimarySourceRegion() {
	obj := metalperformanceshaders.NewBinaryImageKernel()
	obj.PrimarySourceRegion()
	// Output:
	}

