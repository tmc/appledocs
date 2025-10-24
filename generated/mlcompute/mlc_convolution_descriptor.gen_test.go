// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCConvolutionDescriptor

// ExampleNewCConvolutionDescriptorConvolutionTransposeDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountOutputFeatureChannelCount demonstrates how to create a CConvolutionDescriptor instance using NewCConvolutionDescriptorConvolutionTransposeDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountOutputFeatureChannelCount.
// Creates a descriptor for convolution transpose with the kernel sizes and number of feature channels you specify.
func ExampleNewCConvolutionDescriptorConvolutionTransposeDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountOutputFeatureChannelCount() {
	_ = mlcompute.NewCConvolutionDescriptorConvolutionTransposeDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountOutputFeatureChannelCount(
		100, // kernelWidth uint
		100, // kernelHeight uint
		10, // inputFeatureChannelCount uint
		10, // outputFeatureChannelCount uint
	)
	// Output:
}
// ExampleNewCConvolutionDescriptorDepthwiseConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountChannelMultiplier demonstrates how to create a CConvolutionDescriptor instance using NewCConvolutionDescriptorDepthwiseConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountChannelMultiplier.
// Creates a descriptor for depthwise convolution with the kernel sizes, number of input feature channels, and channel multiplier you specify.
func ExampleNewCConvolutionDescriptorDepthwiseConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountChannelMultiplier() {
	_ = mlcompute.NewCConvolutionDescriptorDepthwiseConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountChannelMultiplier(
		100, // kernelWidth uint
		100, // kernelHeight uint
		10, // inputFeatureChannelCount uint
		0, // channelMultiplier uint
	)
	// Output:
}
