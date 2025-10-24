// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCTensorDescriptor

// ExampleNewCTensorDescriptorConvolutionBiasesDescriptorWithFeatureChannelCountDataType demonstrates how to create a CTensorDescriptor instance using NewCTensorDescriptorConvolutionBiasesDescriptorWithFeatureChannelCountDataType.
// Creates a tensor descriptor with the number of feature channels and data type you specify.
func ExampleNewCTensorDescriptorConvolutionBiasesDescriptorWithFeatureChannelCountDataType() {
	_ = mlcompute.NewCTensorDescriptorConvolutionBiasesDescriptorWithFeatureChannelCountDataType(
		10, // featureChannelCount uint
		mlcompute.CDataType{}, // dataType CDataType
	)
	// Output:
}
// ExampleNewCTensorDescriptorConvolutionWeightsDescriptorWithInputFeatureChannelCountOutputFeatureChannelCountDataType demonstrates how to create a CTensorDescriptor instance using NewCTensorDescriptorConvolutionWeightsDescriptorWithInputFeatureChannelCountOutputFeatureChannelCountDataType.
// Creates a tensor descriptor with the number of feature channels and data type you specify.
func ExampleNewCTensorDescriptorConvolutionWeightsDescriptorWithInputFeatureChannelCountOutputFeatureChannelCountDataType() {
	_ = mlcompute.NewCTensorDescriptorConvolutionWeightsDescriptorWithInputFeatureChannelCountOutputFeatureChannelCountDataType(
		10, // inputFeatureChannelCount uint
		10, // outputFeatureChannelCount uint
		mlcompute.CDataType{}, // dataType CDataType
	)
	// Output:
}
// ExampleNewCTensorDescriptorConvolutionWeightsDescriptorWithWidthHeightInputFeatureChannelCountOutputFeatureChannelCountDataType demonstrates how to create a CTensorDescriptor instance using NewCTensorDescriptorConvolutionWeightsDescriptorWithWidthHeightInputFeatureChannelCountOutputFeatureChannelCountDataType.
// Creates a tensor descriptor with the sizing, number of feature channels, and data type you specify.
func ExampleNewCTensorDescriptorConvolutionWeightsDescriptorWithWidthHeightInputFeatureChannelCountOutputFeatureChannelCountDataType() {
	_ = mlcompute.NewCTensorDescriptorConvolutionWeightsDescriptorWithWidthHeightInputFeatureChannelCountOutputFeatureChannelCountDataType(
		100, // width uint
		100, // height uint
		10, // inputFeatureChannelCount uint
		10, // outputFeatureChannelCount uint
		mlcompute.CDataType{}, // dataType CDataType
	)
	// Output:
}
// ExampleNewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSize demonstrates how to create a CTensorDescriptor instance using NewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSize.
// Creates a tensor descriptor with the width and height, number of feature channels, and batch size you specify.
func ExampleNewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSize() {
	_ = mlcompute.NewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSize(
		100, // width uint
		100, // height uint
		0, // featureChannels uint
		10, // batchSize uint
	)
	// Output:
}
// ExampleNewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSizeDataType demonstrates how to create a CTensorDescriptor instance using NewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSizeDataType.
// Creates a tensor descriptor with the width and height, number of feature channels, batch size, and data type you specify.
func ExampleNewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSizeDataType() {
	_ = mlcompute.NewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSizeDataType(
		100, // width uint
		100, // height uint
		10, // featureChannelCount uint
		10, // batchSize uint
		mlcompute.CDataType{}, // dataType CDataType
	)
	// Output:
}
