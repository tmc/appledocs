// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCMultiheadAttentionDescriptor

// ExampleNewCMultiheadAttentionDescriptorWithModelDimensionHeadCount demonstrates how to create a CMultiheadAttentionDescriptor instance using NewCMultiheadAttentionDescriptorWithModelDimensionHeadCount.
// Creates a multi-head attention descriptor with the model dimension and number of parallel attention heads you specify.
func ExampleNewCMultiheadAttentionDescriptorWithModelDimensionHeadCount() {
	_ = mlcompute.NewCMultiheadAttentionDescriptorWithModelDimensionHeadCount(
		0, // modelDimension uint
		10, // headCount uint
	)
	// Output:
}
// ExampleNewCMultiheadAttentionDescriptorWithModelDimensionKeyDimensionValueDimensionHeadCountDropoutHasBiasesHasAttentionBiasesAddsZeroAttention demonstrates how to create a CMultiheadAttentionDescriptor instance using NewCMultiheadAttentionDescriptorWithModelDimensionKeyDimensionValueDimensionHeadCountDropoutHasBiasesHasAttentionBiasesAddsZeroAttention.
// Creates a multi-head attention descriptor with the dimensions, number of attention heads, dropout rate, and bias and padding options you specify.
func ExampleNewCMultiheadAttentionDescriptorWithModelDimensionKeyDimensionValueDimensionHeadCountDropoutHasBiasesHasAttentionBiasesAddsZeroAttention() {
	_ = mlcompute.NewCMultiheadAttentionDescriptorWithModelDimensionKeyDimensionValueDimensionHeadCountDropoutHasBiasesHasAttentionBiasesAddsZeroAttention(
		0, // modelDimension uint
		0, // keyDimension uint
		0, // valueDimension uint
		10, // headCount uint
		0.0, // dropout float32
		false, // hasBiases bool
		false, // hasAttentionBiases bool
		false, // addsZeroAttention bool
	)
	// Output:
}
