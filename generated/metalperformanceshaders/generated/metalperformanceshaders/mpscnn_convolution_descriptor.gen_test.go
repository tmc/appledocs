// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewCNNConvolutionDescriptor

// ExampleNewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels demonstrates how to create a CNNConvolutionDescriptor instance using NewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels.
func ExampleNewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels() {
	_ = metalperformanceshaders.NewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(
		100, // kernelWidth uint
		100, // kernelHeight uint
		0, // inputFeatureChannels uint
		0, // outputFeatureChannels uint
	)
	// Output:
}
// ExampleNewCNNConvolutionDescriptorWithCoder demonstrates how to create a CNNConvolutionDescriptor instance using NewCNNConvolutionDescriptorWithCoder.
func ExampleNewCNNConvolutionDescriptorWithCoder() {
	_ = metalperformanceshaders.NewCNNConvolutionDescriptorWithCoder(
		metalperformanceshaders.Coder /* not a class type */{}, // aDecoder Coder /* not a class type */
	)
	// Output:
}
// ExampleCNNConvolutionDescriptor_Encode demonstrates using Encode on a CNNConvolutionDescriptor instance.
func ExampleCNNConvolutionDescriptor_Encode() {
	obj := metalperformanceshaders.NewCNNConvolutionDescriptor()
	obj.Encode()
	// Output:
	}

// ExampleCNNConvolutionDescriptor_SetBatchNormalizationParametersForInferenceWithMean demonstrates using SetBatchNormalizationParametersForInferenceWithMean on a CNNConvolutionDescriptor instance.
func ExampleCNNConvolutionDescriptor_SetBatchNormalizationParametersForInferenceWithMean() {
	obj := metalperformanceshaders.NewCNNConvolutionDescriptor()
	obj.SetBatchNormalizationParametersForInferenceWithMean()
	// Output:
	}

