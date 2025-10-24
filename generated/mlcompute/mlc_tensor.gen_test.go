// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCTensor

// ExampleNewCTensorWithSequenceLengthFeatureChannelCountBatchSize demonstrates how to create a CTensor instance using NewCTensorWithSequenceLengthFeatureChannelCountBatchSize.
// Creates a tensor without data, with the sequence length, number of feature channels, and batch size you specify.
func ExampleNewCTensorWithSequenceLengthFeatureChannelCountBatchSize() {
	_ = mlcompute.NewCTensorWithSequenceLengthFeatureChannelCountBatchSize(
		10, // sequenceLength uint
		10, // featureChannelCount uint
		10, // batchSize uint
	)
	// Output:
}
// ExampleNewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType demonstrates how to create a CTensor instance using NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType.
// Creates a tensor with the sequence length, number of feature channels, batch size, and random initializer type you specify.
func ExampleNewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType() {
	_ = mlcompute.NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType(
		10, // sequenceLength uint
		10, // featureChannelCount uint
		10, // batchSize uint
		mlcompute.CRandomInitializerType{}, // randomInitializerType CRandomInitializerType
	)
	// Output:
}
// ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSize demonstrates how to create a CTensor instance using NewCTensorWithWidthHeightFeatureChannelCountBatchSize.
// Creates a tensor without data, with the sizes and number of feature channels you specify.
func ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSize() {
	_ = mlcompute.NewCTensorWithWidthHeightFeatureChannelCountBatchSize(
		100, // width uint
		100, // height uint
		10, // featureChannelCount uint
		10, // batchSize uint
	)
	// Output:
}
// ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType demonstrates how to create a CTensor instance using NewCTensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType.
// Creates a tensor with the sizes and number of feature channels, and filled with the data and type you specify.
func ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType() {
	_ = mlcompute.NewCTensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType(
		100, // width uint
		100, // height uint
		10, // featureChannelCount uint
		10, // batchSize uint
		0.0, // fillData float32
		mlcompute.CDataType{}, // dataType CDataType
	)
	// Output:
}
// ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType demonstrates how to create a CTensor instance using NewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType.
// Creates a tensor with the sizes, number of feature channels, and random data using the random initializer type you specify.
func ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType() {
	_ = mlcompute.NewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType(
		100, // width uint
		100, // height uint
		10, // featureChannelCount uint
		10, // batchSize uint
		mlcompute.CRandomInitializerType{}, // randomInitializerType CRandomInitializerType
	)
	// Output:
}
