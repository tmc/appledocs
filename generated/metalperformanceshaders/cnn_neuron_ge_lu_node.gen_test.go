// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewCNNNeuronGeLUNode

// ExampleNewCNNNeuronGeLUNodeWithSource demonstrates how to create a CNNNeuronGeLUNode instance using NewCNNNeuronGeLUNodeWithSource.
func ExampleNewCNNNeuronGeLUNodeWithSource() {
	_ = metalperformanceshaders.NewCNNNeuronGeLUNodeWithSource(
		metalperformanceshaders.MPSNNImageNode{}, // sourceNode MPSNNImageNode
	)
	// Output:
}
