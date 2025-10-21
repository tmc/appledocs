// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewGramMatrixCalculationNode

// ExampleNewGramMatrixCalculationNodeWithSource demonstrates how to create a GramMatrixCalculationNode instance using NewGramMatrixCalculationNodeWithSource.
func ExampleNewGramMatrixCalculationNodeWithSource() {
	_ = metalperformanceshaders.NewGramMatrixCalculationNodeWithSource(
		metalperformanceshaders.MPSNNImageNode{}, // sourceNode MPSNNImageNode
	)
	// Output:
}
