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
		0, // sequenceLength uint
		0, // featureChannelCount uint
		0, // batchSize uint
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
