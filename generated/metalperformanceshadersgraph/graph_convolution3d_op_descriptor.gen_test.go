// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshadersgraph"
)

// Suppress unused import errors
var _ = metalperformanceshadersgraph.NewGraphConvolution3DOpDescriptor

// ExampleNewGraphConvolution3DOpDescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBackPaddingStyleDataLayoutWeightsLayout demonstrates how to create a GraphConvolution3DOpDescriptor instance using NewGraphConvolution3DOpDescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBackPaddingStyleDataLayoutWeightsLayout.
// Creates a convolution descriptor with given values for parameters.
func ExampleNewGraphConvolution3DOpDescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBackPaddingStyleDataLayoutWeightsLayout() {
	_ = metalperformanceshadersgraph.NewGraphConvolution3DOpDescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBackPaddingStyleDataLayoutWeightsLayout(
		0, // strideInX uint
		0, // strideInY uint
		0, // strideInZ uint
		0, // dilationRateInX uint
		0, // dilationRateInY uint
		0, // dilationRateInZ uint
		0, // groups uint
		0, // paddingLeft uint
		0, // paddingRight uint
		0, // paddingTop uint
		0, // paddingBottom uint
		0, // paddingFront uint
		0, // paddingBack uint
		metalperformanceshadersgraph.GraphPaddingStyle{},          // paddingStyle GraphPaddingStyle
		metalperformanceshadersgraph.GraphTensorNamedDataLayout{}, // dataLayout GraphTensorNamedDataLayout
		metalperformanceshadersgraph.GraphTensorNamedDataLayout{}, // weightsLayout GraphTensorNamedDataLayout
	)
	// Output:
}
