// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewReshapeNode

// ExampleNewReshapeNodeWithSourceResultWidthResultHeightResultFeatureChannels demonstrates how to create a ReshapeNode instance using NewReshapeNodeWithSourceResultWidthResultHeightResultFeatureChannels.
func ExampleNewReshapeNodeWithSourceResultWidthResultHeightResultFeatureChannels() {
	_ = metalperformanceshaders.NewReshapeNodeWithSourceResultWidthResultHeightResultFeatureChannels(
		metalperformanceshaders.MPSNNImageNode{}, // source MPSNNImageNode
		0,                                        // resultWidth uint
		0,                                        // resultHeight uint
		0,                                        // resultFeatureChannels uint
	)
	// Output:
}
