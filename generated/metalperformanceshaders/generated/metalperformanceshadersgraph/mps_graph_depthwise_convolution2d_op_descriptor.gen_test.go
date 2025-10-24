// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshadersgraph"
)

// Suppress unused import errors
var _ = metalperformanceshadersgraph.NewGraphDepthwiseConvolution2DOpDescriptor

// ExampleNewGraphDepthwiseConvolution2DOpDescriptorWithDataLayoutWeightsLayout demonstrates how to create a GraphDepthwiseConvolution2DOpDescriptor instance using NewGraphDepthwiseConvolution2DOpDescriptorWithDataLayoutWeightsLayout.
// Creates a 2D-depthwise convolution descriptor with given properties and default values.
func ExampleNewGraphDepthwiseConvolution2DOpDescriptorWithDataLayoutWeightsLayout() {
	_ = metalperformanceshadersgraph.NewGraphDepthwiseConvolution2DOpDescriptorWithDataLayoutWeightsLayout(
		metalperformanceshadersgraph.GraphTensorNamedDataLayout{}, // dataLayout GraphTensorNamedDataLayout
		metalperformanceshadersgraph.GraphTensorNamedDataLayout{}, // weightsLayout GraphTensorNamedDataLayout
	)
	// Output:
}
// ExampleNewGraphDepthwiseConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout demonstrates how to create a GraphDepthwiseConvolution2DOpDescriptor instance using NewGraphDepthwiseConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout.
// Creates a 2D-depthwise convolution descriptor with given values.
func ExampleNewGraphDepthwiseConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout() {
	_ = metalperformanceshadersgraph.NewGraphDepthwiseConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(
		0, // strideInX uint
		0, // strideInY uint
		0, // dilationRateInX uint
		0, // dilationRateInY uint
		0, // paddingLeft uint
		0, // paddingRight uint
		0, // paddingTop uint
		0, // paddingBottom uint
		metalperformanceshadersgraph.GraphPaddingStyle{}, // paddingStyle GraphPaddingStyle
		metalperformanceshadersgraph.GraphTensorNamedDataLayout{}, // dataLayout GraphTensorNamedDataLayout
		metalperformanceshadersgraph.GraphTensorNamedDataLayout{}, // weightsLayout GraphTensorNamedDataLayout
	)
	// Output:
}
