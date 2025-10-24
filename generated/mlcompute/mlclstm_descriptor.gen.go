// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCLSTMDescriptor */


/* debug [class_header]: Header for MLCLSTMDescriptor */
// The class instance for the [CLSTMDescriptor] class.
var (
	CLSTMDescriptorClass     _CLSTMDescriptorClass
	CLSTMDescriptorClassOnce sync.Once
)

func getCLSTMDescriptorClass() _CLSTMDescriptorClass {
	CLSTMDescriptorClassOnce.Do(func() {
		CLSTMDescriptorClass = _CLSTMDescriptorClass{objc.GetClass("MLCLSTMDescriptor")}
	})
	return CLSTMDescriptorClass
}

type _CLSTMDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CLSTMDescriptor */
// An interface definition for the [CLSTMDescriptor] class.
type ICLSTMDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CLSTMDescriptor */
	// properties:
	BatchFirst() bool
	Dropout() float32
	HiddenSize() uint
	InputSize() uint
	IsBidirectional() bool
	LayerCount() uint
	ResultMode() CLSTMResultMode
	ReturnsSequences() bool
	UsesBiases() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CLSTMDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CLSTMDescriptor */
// Alloc allocates a new instance without initialization.
func (sc _CLSTMDescriptorClass) Alloc() CLSTMDescriptor {
	rv := objc.Send[CLSTMDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _CLSTMDescriptorClass) New() CLSTMDescriptor {
	rv := objc.Send[CLSTMDescriptor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ CLSTMDescriptor) Init() CLSTMDescriptor {
	rv := objc.Send[CLSTMDescriptor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ CLSTMDescriptor) Autorelease() CLSTMDescriptor {
	rv := objc.Send[CLSTMDescriptor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLSTMDescriptor creates a new CLSTMDescriptor instance.
func NewCLSTMDescriptor() CLSTMDescriptor {
	return getCLSTMDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CLSTMDescriptor */
// The configuration object you use to create the LSTM layer.


// The configuration object you use to create the LSTM layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor
type CLSTMDescriptor struct {
	objectivec.Object
}

// CLSTMDescriptorFrom constructs a [CLSTMDescriptor] from an unsafe.Pointer.
//
// The configuration object you use to create the LSTM layer.
func CLSTMDescriptorFrom(ptr unsafe.Pointer) CLSTMDescriptor {
	return CLSTMDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CLSTMDescriptor */

// Creates a batch first LSTM descriptor with the input size and number of layers you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/init(inputSize:hiddenSize:layerCount:)
func NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCount(inputSize uint, hiddenSize uint, layerCount uint) CLSTMDescriptor {
	rv := objc.Send[CLSTMDescriptor](objc.ID(getCLSTMDescriptorClass().class), objc.Sel("descriptorWithInputSize:hiddenSize:layerCount:"), inputSize, hiddenSize, layerCount)
	return rv
}/* debug [class_init_methods/constructor]: NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCount */


// Creates a batch first LSTM descriptor that allows you to indicate whether the input and output shape is batch first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/init(inputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:dropout:)
func NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalDropout(inputSize uint, hiddenSize uint, layerCount uint, usesBiases bool, batchFirst bool, isBidirectional bool, dropout float32) CLSTMDescriptor {
	rv := objc.Send[CLSTMDescriptor](objc.ID(getCLSTMDescriptorClass().class), objc.Sel("descriptorWithInputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:dropout:"), inputSize, hiddenSize, layerCount, usesBiases, batchFirst, isBidirectional, dropout)
	return rv
}/* debug [class_init_methods/constructor]: NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalDropout */


// Creates a batch first LSTM descriptor that allows you to indicate whether the layer returns output for all sequences, or output for only the last sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/init(inputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:returnsSequences:dropout:)
func NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropout(inputSize uint, hiddenSize uint, layerCount uint, usesBiases bool, batchFirst bool, isBidirectional bool, returnsSequences bool, dropout float32) CLSTMDescriptor {
	rv := objc.Send[CLSTMDescriptor](objc.ID(getCLSTMDescriptorClass().class), objc.Sel("descriptorWithInputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:returnsSequences:dropout:"), inputSize, hiddenSize, layerCount, usesBiases, batchFirst, isBidirectional, returnsSequences, dropout)
	return rv
}/* debug [class_init_methods/constructor]: NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropout */


// Creates a descriptor with the number of features and layers, dropout, and options for use of biases, batch order, return sequences, bidirectionality, and expected tensors you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/init(inputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:returnsSequences:dropout:resultMode:)
func NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropoutResultMode(inputSize uint, hiddenSize uint, layerCount uint, usesBiases bool, batchFirst bool, isBidirectional bool, returnsSequences bool, dropout float32, resultMode STMResultMode) CLSTMDescriptor {
	rv := objc.Send[CLSTMDescriptor](objc.ID(getCLSTMDescriptorClass().class), objc.Sel("descriptorWithInputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:returnsSequences:dropout:resultMode:"), inputSize, hiddenSize, layerCount, usesBiases, batchFirst, isBidirectional, returnsSequences, dropout, resultMode)
	return rv
}/* debug [class_init_methods/constructor]: NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropoutResultMode */


// Creates a batch first LSTM descriptor with bias and bidirectional options you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/init(inputSize:hiddenSize:layerCount:usesBiases:isBidirectional:dropout:)
func NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesIsBidirectionalDropout(inputSize uint, hiddenSize uint, layerCount uint, usesBiases bool, isBidirectional bool, dropout float32) CLSTMDescriptor {
	rv := objc.Send[CLSTMDescriptor](objc.ID(getCLSTMDescriptorClass().class), objc.Sel("descriptorWithInputSize:hiddenSize:layerCount:usesBiases:isBidirectional:dropout:"), inputSize, hiddenSize, layerCount, usesBiases, isBidirectional, dropout)
	return rv
}/* debug [class_init_methods/constructor]: NewCLSTMDescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesIsBidirectionalDropout */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CLSTMDescriptor */

// Creates a batch first LSTM descriptor with the input size and number of layers you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/init(inputSize:hiddenSize:layerCount:)
func (sc _CLSTMDescriptorClass) DescriptorWithInputSizeHiddenSizeLayerCount(inputSize uint, hiddenSize uint, layerCount uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("descriptorWithInputSize:hiddenSize:layerCount:"), inputSize, hiddenSize, layerCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithInputSizeHiddenSizeLayerCount) */


// Creates a batch first LSTM descriptor that allows you to indicate whether the input and output shape is batch first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/init(inputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:dropout:)
func (sc _CLSTMDescriptorClass) DescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalDropout(inputSize uint, hiddenSize uint, layerCount uint, usesBiases bool, batchFirst bool, isBidirectional bool, dropout float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("descriptorWithInputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:dropout:"), inputSize, hiddenSize, layerCount, usesBiases, batchFirst, isBidirectional, dropout)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalDropout) */


// Creates a batch first LSTM descriptor that allows you to indicate whether the layer returns output for all sequences, or output for only the last sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/init(inputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:returnsSequences:dropout:)
func (sc _CLSTMDescriptorClass) DescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropout(inputSize uint, hiddenSize uint, layerCount uint, usesBiases bool, batchFirst bool, isBidirectional bool, returnsSequences bool, dropout float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("descriptorWithInputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:returnsSequences:dropout:"), inputSize, hiddenSize, layerCount, usesBiases, batchFirst, isBidirectional, returnsSequences, dropout)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropout) */


// Creates a descriptor with the number of features and layers, dropout, and options for use of biases, batch order, return sequences, bidirectionality, and expected tensors you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/init(inputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:returnsSequences:dropout:resultMode:)
func (sc _CLSTMDescriptorClass) DescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropoutResultMode(inputSize uint, hiddenSize uint, layerCount uint, usesBiases bool, batchFirst bool, isBidirectional bool, returnsSequences bool, dropout float32, resultMode STMResultMode) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("descriptorWithInputSize:hiddenSize:layerCount:usesBiases:batchFirst:isBidirectional:returnsSequences:dropout:resultMode:"), inputSize, hiddenSize, layerCount, usesBiases, batchFirst, isBidirectional, returnsSequences, dropout, resultMode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesBatchFirstIsBidirectionalReturnsSequencesDropoutResultMode) */


// Creates a batch first LSTM descriptor with bias and bidirectional options you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/init(inputSize:hiddenSize:layerCount:usesBiases:isBidirectional:dropout:)
func (sc _CLSTMDescriptorClass) DescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesIsBidirectionalDropout(inputSize uint, hiddenSize uint, layerCount uint, usesBiases bool, isBidirectional bool, dropout float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("descriptorWithInputSize:hiddenSize:layerCount:usesBiases:isBidirectional:dropout:"), inputSize, hiddenSize, layerCount, usesBiases, isBidirectional, dropout)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithInputSizeHiddenSizeLayerCountUsesBiasesIsBidirectionalDropout) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CLSTMDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CLSTMDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CLSTMDescriptor */

// A Boolean that indicates whether the input and output shape is batch first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/batchFirst
func (s_ CLSTMDescriptor) BatchFirst() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("batchFirst"))
	return rv
}/* debug [instance_properties/getter]: batchFirst */


// The dropout probability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/dropout
func (s_ CLSTMDescriptor) Dropout() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("dropout"))
	return rv
}/* debug [instance_properties/getter]: dropout */


// The number of features in the hidden state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/hiddenSize
func (s_ CLSTMDescriptor) HiddenSize() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("hiddenSize"))
	return rv
}/* debug [instance_properties/getter]: hiddenSize */


// The number of expected features in the input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/inputSize
func (s_ CLSTMDescriptor) InputSize() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("inputSize"))
	return rv
}/* debug [instance_properties/getter]: inputSize */


// A Boolean that indicates whether the layer is bidirectional.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/isBidirectional
func (s_ CLSTMDescriptor) IsBidirectional() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isBidirectional"))
	return rv
}/* debug [instance_properties/getter]: isBidirectional */


// The number of recurrent layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/layerCount
func (s_ CLSTMDescriptor) LayerCount() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("layerCount"))
	return rv
}/* debug [instance_properties/getter]: layerCount */


// The mode that indicates whether the layer produces a single result tensor or three result tensors — final output, last hidden state, and the cell state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/resultMode
func (s_ CLSTMDescriptor) ResultMode() CLSTMResultMode {
	rv := objc.Send[CLSTMResultMode](s_.ID, objc.Sel("resultMode"))
	return rv
}/* debug [instance_properties/getter]: resultMode */


// A Boolean that indicates whether the layer returns output for all sequences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/returnsSequences
func (s_ CLSTMDescriptor) ReturnsSequences() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("returnsSequences"))
	return rv
}/* debug [instance_properties/getter]: returnsSequences */


// A Boolean that indicates whether you use bias weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMDescriptor/usesBiases
func (s_ CLSTMDescriptor) UsesBiases() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("usesBiases"))
	return rv
}/* debug [instance_properties/getter]: usesBiases */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCLSTMDescriptor */


