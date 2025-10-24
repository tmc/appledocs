// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewImageHistogramEqualization

// ExampleImageHistogramEqualization_EncodeTransform demonstrates using EncodeTransform on a ImageHistogramEqualization instance.
// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
func ExampleImageHistogramEqualization_EncodeTransform() {
	obj := metalperformanceshaders.NewImageHistogramEqualization()
	obj.EncodeTransform()
	// Output:
	}

