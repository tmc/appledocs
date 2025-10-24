// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshadersgraph"
)

// Suppress unused import errors
var _ = metalperformanceshadersgraph.NewGraphImToColOpDescriptor

// ExampleNewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYDataLayout demonstrates how to create a GraphImToColOpDescriptor instance using NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYDataLayout.
// Creates column to image descriptor with given values for parameters.
func ExampleNewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYDataLayout() {
	_ = metalperformanceshadersgraph.NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYDataLayout(
		100, // kernelWidth uint
		100, // kernelHeight uint
		0, // strideInX uint
		0, // strideInY uint
		0, // dilationRateInX uint
		0, // dilationRateInY uint
		metalperformanceshadersgraph.GraphTensorNamedDataLayout{}, // dataLayout GraphTensorNamedDataLayout
	)
	// Output:
}
// ExampleNewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomDataLayout demonstrates how to create a GraphImToColOpDescriptor instance using NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomDataLayout.
// Creates an image to column descriptor with given values for parameters.
func ExampleNewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomDataLayout() {
	_ = metalperformanceshadersgraph.NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomDataLayout(
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
		metalperformanceshadersgraph.GraphTensorNamedDataLayout{}, // dataLayout GraphTensorNamedDataLayout
	)
	// Output:
}
