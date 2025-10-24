// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCLSTMDescriptor

// ExampleNewCLSTMDescriptorWithInputSizeHiddenSizeLayerCount demonstrates how to create a CLSTMDescriptor instance using NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCount.
// Creates a batch first LSTM descriptor with the input size and number of layers you specify.
func ExampleNewCLSTMDescriptorWithInputSizeHiddenSizeLayerCount() {
	_ = mlcompute.NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCount(
		10, // inputSize uint
		10, // hiddenSize uint
		10, // layerCount uint
	)
	// Output:
}
// ExampleNewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalDropout demonstrates how to create a CLSTMDescriptor instance using NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalDropout.
// Creates a batch first LSTM descriptor that allows you to indicate whether the input and output shape is batch first.
func ExampleNewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalDropout() {
	_ = mlcompute.NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalDropout(
		10, // inputSize uint
		10, // hiddenSize uint
		10, // layerCount uint
		false, // usesBiases bool
		false, // batchFirst bool
		false, // isBidirectional bool
		0.0, // dropout float32
	)
	// Output:
}
// ExampleNewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropout demonstrates how to create a CLSTMDescriptor instance using NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropout.
// Creates a batch first LSTM descriptor that allows you to indicate whether the layer returns output for all sequences, or output for only the last sequence.
func ExampleNewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropout() {
	_ = mlcompute.NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropout(
		10, // inputSize uint
		10, // hiddenSize uint
		10, // layerCount uint
		false, // usesBiases bool
		false, // batchFirst bool
		false, // isBidirectional bool
		false, // returnsSequences bool
		0.0, // dropout float32
	)
	// Output:
}
// ExampleNewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropoutResultMode demonstrates how to create a CLSTMDescriptor instance using NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropoutResultMode.
// Creates a descriptor with the number of features and layers, dropout, and options for use of biases, batch order, return sequences, bidirectionality, and expected tensors you specify.
func ExampleNewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropoutResultMode() {
	_ = mlcompute.NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropoutResultMode(
		10, // inputSize uint
		10, // hiddenSize uint
		10, // layerCount uint
		false, // usesBiases bool
		false, // batchFirst bool
		false, // isBidirectional bool
		false, // returnsSequences bool
		0.0, // dropout float32
		mlcompute.CLSTMResultMode{}, // resultMode CLSTMResultMode
	)
	// Output:
}
// ExampleNewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesIsBidirectionalDropout demonstrates how to create a CLSTMDescriptor instance using NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesIsBidirectionalDropout.
// Creates a batch first LSTM descriptor with bias and bidirectional options you specify.
func ExampleNewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesIsBidirectionalDropout() {
	_ = mlcompute.NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesIsBidirectionalDropout(
		10, // inputSize uint
		10, // hiddenSize uint
		10, // layerCount uint
		false, // usesBiases bool
		false, // isBidirectional bool
		0.0, // dropout float32
	)
	// Output:
}
