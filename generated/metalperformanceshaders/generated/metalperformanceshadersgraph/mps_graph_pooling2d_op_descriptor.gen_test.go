// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshadersgraph"
)

// Suppress unused import errors
var _ = metalperformanceshadersgraph.NewGraphPooling2DOpDescriptor

// ExampleNewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayout demonstrates how to create a GraphPooling2DOpDescriptor instance using NewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayout.
// Creates a 2D pooling descriptor with given values.
func ExampleNewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayout() {
	_ = metalperformanceshadersgraph.NewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayout(
		100, // kernelWidth uint
		100, // kernelHeight uint
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
	)
	// Output:
}
// ExampleNewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout demonstrates how to create a GraphPooling2DOpDescriptor instance using NewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout.
// Creates a 2D pooling descriptor with given values.
func ExampleNewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout() {
	_ = metalperformanceshadersgraph.NewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout(
		100, // kernelWidth uint
		100, // kernelHeight uint
		0, // strideInX uint
		0, // strideInY uint
		metalperformanceshadersgraph.GraphPaddingStyle{}, // paddingStyle GraphPaddingStyle
		metalperformanceshadersgraph.GraphTensorNamedDataLayout{}, // dataLayout GraphTensorNamedDataLayout
	)
	// Output:
}
