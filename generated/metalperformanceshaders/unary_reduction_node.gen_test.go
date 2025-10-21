// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewUnaryReductionNode

// ExampleNewUnaryReductionNodeWithSource demonstrates how to create a UnaryReductionNode instance using NewUnaryReductionNodeWithSource.
func ExampleNewUnaryReductionNodeWithSource() {
	_ = metalperformanceshaders.NewUnaryReductionNodeWithSource(
		metalperformanceshaders.MPSNNImageNode{}, // sourceNode MPSNNImageNode
	)
	// Output:
}
