// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewInitialGradientNode

// ExampleNewInitialGradientNodeWithSource demonstrates how to create a InitialGradientNode instance using NewInitialGradientNodeWithSource.
func ExampleNewInitialGradientNodeWithSource() {
	_ = metalperformanceshaders.NewInitialGradientNodeWithSource(
		metalperformanceshaders.MPSNNImageNode{}, // source MPSNNImageNode
	)
	// Output:
}
