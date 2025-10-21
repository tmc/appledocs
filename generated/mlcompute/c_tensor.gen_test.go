// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCTensor

// ExampleNewCTensorWithDescriptor demonstrates how to create a CTensor instance using NewCTensorWithDescriptor.
// Creates a tensor without data, using the descriptor you specify.
func ExampleNewCTensorWithDescriptor() {
	_ = mlcompute.NewCTensorWithDescriptor(
		mlcompute.MLCTensorDescriptor{}, // tensorDescriptor MLCTensorDescriptor
	)
	// Output:
}
// ExampleNewCTensorWithDescriptorData demonstrates how to create a CTensor instance using NewCTensorWithDescriptorData.
// Creates a tensor with the descriptor and data you specify.
func ExampleNewCTensorWithDescriptorData() {
	_ = mlcompute.NewCTensorWithDescriptorData(
		mlcompute.MLCTensorDescriptor{}, // tensorDescriptor MLCTensorDescriptor
		mlcompute.MLCTensorData{}, // data MLCTensorData
	)
	// Output:
}
// ExampleNewCTensorWithDescriptorRandomInitializerType demonstrates how to create a CTensor instance using NewCTensorWithDescriptorRandomInitializerType.
// Creates a tensor with the descriptor and random initializer type you specify.
func ExampleNewCTensorWithDescriptorRandomInitializerType() {
	_ = mlcompute.NewCTensorWithDescriptorRandomInitializerType(
		mlcompute.MLCTensorDescriptor{}, // tensorDescriptor MLCTensorDescriptor
		mlcompute.CRandomInitializerType{}, // randomInitializerType CRandomInitializerType
	)
	// Output:
}
// ExampleNewCTensorWithSequenceLengthFeatureChannelCountBatchSize demonstrates how to create a CTensor instance using NewCTensorWithSequenceLengthFeatureChannelCountBatchSize.
// Creates a tensor without data, with the sequence length, number of feature channels, and batch size you specify.
func ExampleNewCTensorWithSequenceLengthFeatureChannelCountBatchSize() {
	_ = mlcompute.NewCTensorWithSequenceLengthFeatureChannelCountBatchSize(
		0, // sequenceLength uint
		0, // featureChannelCount uint
		0, // batchSize uint
	)
	// Output:
}
// ExampleNewCTensorWithSequenceLengthFeatureChannelCountBatchSizeData demonstrates how to create a CTensor instance using NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeData.
// Creates a tensor with the sequence length, number of feature channels, batch size, and data you specify.
func ExampleNewCTensorWithSequenceLengthFeatureChannelCountBatchSizeData() {
	_ = mlcompute.NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeData(
		0, // sequenceLength uint
		0, // featureChannelCount uint
		0, // batchSize uint
		mlcompute.MLCTensorData{}, // data MLCTensorData
	)
	// Output:
}
// ExampleNewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType demonstrates how to create a CTensor instance using NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType.
// Creates a tensor with the sequence length, number of feature channels, batch size, and random initializer type you specify.
func ExampleNewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType() {
	_ = mlcompute.NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType(
		0, // sequenceLength uint
		0, // featureChannelCount uint
		0, // batchSize uint
		mlcompute.CRandomInitializerType{}, // randomInitializerType CRandomInitializerType
	)
	// Output:
}
// ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSize demonstrates how to create a CTensor instance using NewCTensorWithWidthHeightFeatureChannelCountBatchSize.
// Creates a tensor without data, with the sizes and number of feature channels you specify.
func ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSize() {
	_ = mlcompute.NewCTensorWithWidthHeightFeatureChannelCountBatchSize(
		0, // width uint
		0, // height uint
		0, // featureChannelCount uint
		0, // batchSize uint
	)
	// Output:
}
// ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSizeData demonstrates how to create a CTensor instance using NewCTensorWithWidthHeightFeatureChannelCountBatchSizeData.
// Creates a tensor with the sizes, number of feature channels, and data you specify.
func ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSizeData() {
	_ = mlcompute.NewCTensorWithWidthHeightFeatureChannelCountBatchSizeData(
		0, // width uint
		0, // height uint
		0, // featureChannelCount uint
		0, // batchSize uint
		mlcompute.MLCTensorData{}, // data MLCTensorData
	)
	// Output:
}
// ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType demonstrates how to create a CTensor instance using NewCTensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType.
// Creates a tensor with the sizes, number of feature channels, data, and data type you specify.
func ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType() {
	_ = mlcompute.NewCTensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType(
		0, // width uint
		0, // height uint
		0, // featureChannelCount uint
		0, // batchSize uint
		mlcompute.MLCTensorData{}, // data MLCTensorData
		mlcompute.CDataType{}, // dataType CDataType
	)
	// Output:
}
// ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType demonstrates how to create a CTensor instance using NewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType.
// Creates a tensor with the sizes, number of feature channels, and random data using the random initializer type you specify.
func ExampleNewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType() {
	_ = mlcompute.NewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType(
		0, // width uint
		0, // height uint
		0, // featureChannelCount uint
		0, // batchSize uint
		mlcompute.CRandomInitializerType{}, // randomInitializerType CRandomInitializerType
	)
	// Output:
}
