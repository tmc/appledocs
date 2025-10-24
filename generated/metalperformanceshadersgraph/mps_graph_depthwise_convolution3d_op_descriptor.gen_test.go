// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshadersgraph"
)

// Suppress unused import errors
var _ = metalperformanceshadersgraph.NewGraphDepthwiseConvolution3DOpDescriptor

// ExampleNewGraphDepthwiseConvolution3DOpDescriptorWithPaddingStyle demonstrates how to create a GraphDepthwiseConvolution3DOpDescriptor instance using NewGraphDepthwiseConvolution3DOpDescriptorWithPaddingStyle.
// Creates a 3D depthwise convolution descriptor with default values.
func ExampleNewGraphDepthwiseConvolution3DOpDescriptorWithPaddingStyle() {
	_ = metalperformanceshadersgraph.NewGraphDepthwiseConvolution3DOpDescriptorWithPaddingStyle(
		metalperformanceshadersgraph.GraphPaddingStyle{}, // paddingStyle GraphPaddingStyle
	)
	// Output:
}
