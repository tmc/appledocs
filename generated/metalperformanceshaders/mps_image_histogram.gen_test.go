// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewImageHistogram

// ExampleImageHistogram_HistogramSize demonstrates using HistogramSize on a ImageHistogram instance.
// The amount of space the histogram will take up in the output buffer.
func ExampleImageHistogram_HistogramSize() {
	obj := metalperformanceshaders.NewImageHistogram()
	obj.HistogramSize()
	// Output:
	}

// ExampleImageHistogram_Encode demonstrates using Encode on a ImageHistogram instance.
// Encodes the filter to a command buffer using a compute command encoder.
func ExampleImageHistogram_Encode() {
	obj := metalperformanceshaders.NewImageHistogram()
	obj.Encode()
	// Output:
	}

