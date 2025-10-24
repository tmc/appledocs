// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewKernel

// ExampleNewKernelWithCoder demonstrates how to create a Kernel instance using NewKernelWithCoder.
func ExampleNewKernelWithCoder() {
	_ = metalperformanceshaders.NewKernelWithCoder(
		metalperformanceshaders.Coder /* not a class type */{}, // aDecoder Coder /* not a class type */
	)
	// Output:
}
