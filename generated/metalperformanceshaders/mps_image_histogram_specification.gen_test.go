// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewImageHistogramSpecification

// ExampleImageHistogramSpecification_EncodeTransform demonstrates using EncodeTransform on a ImageHistogramSpecification instance.
// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
func ExampleImageHistogramSpecification_EncodeTransform() {
	obj := metalperformanceshaders.NewImageHistogramSpecification()
	obj.EncodeTransform()
	// Output:
	}

