// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphLSTMDescriptor] class.
var (
	GraphLSTMDescriptorClass     _GraphLSTMDescriptorClass
	GraphLSTMDescriptorClassOnce sync.Once
)

func getGraphLSTMDescriptorClass() _GraphLSTMDescriptorClass {
	GraphLSTMDescriptorClassOnce.Do(func() {
		GraphLSTMDescriptorClass = _GraphLSTMDescriptorClass{objc.GetClass("MPSGraphLSTMDescriptor")}
	})
	return GraphLSTMDescriptorClass
}

type _GraphLSTMDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphLSTMDescriptor] class.
type IGraphLSTMDescriptor interface {
	IGraphObject
}

// The class that defines the parameters for a long short-term memory (LSTM) operation.
//
// Use this descriptor with the following methods:
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor
type GraphLSTMDescriptor struct {
	GraphObject
}

// GraphLSTMDescriptorFrom constructs a [GraphLSTMDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a long short-term memory (LSTM) operation.
func GraphLSTMDescriptorFrom(ptr unsafe.Pointer) GraphLSTMDescriptor {
	return GraphLSTMDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphLSTMDescriptorClass) Alloc() GraphLSTMDescriptor {
	rv := objc.Send[GraphLSTMDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphLSTMDescriptorClass) New() GraphLSTMDescriptor {
	rv := objc.Send[GraphLSTMDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphLSTMDescriptor) Init() GraphLSTMDescriptor {
	rv := objc.Send[GraphLSTMDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphLSTMDescriptor) Autorelease() GraphLSTMDescriptor {
	rv := objc.Send[GraphLSTMDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphLSTMDescriptor creates a new GraphLSTMDescriptor instance.
func NewGraphLSTMDescriptor() GraphLSTMDescriptor {
	return getGraphLSTMDescriptorClass().New()
}


// A parameter that defines the activation function used with the input gate of the LSTM operation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/inputgateactivation
func (g_ GraphLSTMDescriptor) InputGateActivation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("inputGateActivation"))
	return rv
}


// SetInputGateActivation sets the value of the inputGateActivation property.
// A parameter that defines the activation function used with the input gate of the LSTM operation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/inputgateactivation
func (g_ GraphLSTMDescriptor) SetInputGateActivation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInputGateActivation:"), value)
}

// A parameter that defines the activation function used with the current cell value of the LSTM operation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/activation
func (g_ GraphLSTMDescriptor) Activation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("activation"))
	return rv
}


// SetActivation sets the value of the activation property.
// A parameter that defines the activation function used with the current cell value of the LSTM operation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/activation
func (g_ GraphLSTMDescriptor) SetActivation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setActivation:"), value)
}

// A parameter that controls whether or not to return the output cell from the LSTM layer.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/producecell
func (g_ GraphLSTMDescriptor) ProduceCell() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("produceCell"))
	return rv
}


// SetProduceCell sets the value of the produceCell property.
// A parameter that controls whether or not to return the output cell from the LSTM layer.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/producecell
func (g_ GraphLSTMDescriptor) SetProduceCell(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProduceCell:"), value)
}

// A parameter that defines the activation function used with the output gate of the LSTM operation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/outputgateactivation
func (g_ GraphLSTMDescriptor) OutputGateActivation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("outputGateActivation"))
	return rv
}


// SetOutputGateActivation sets the value of the outputGateActivation property.
// A parameter that defines the activation function used with the output gate of the LSTM operation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/outputgateactivation
func (g_ GraphLSTMDescriptor) SetOutputGateActivation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputGateActivation:"), value)
}

// A parameter that controls the internal order of the LSTM gates.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/forgetgatelast
func (g_ GraphLSTMDescriptor) ForgetGateLast() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("forgetGateLast"))
	return rv
}


// SetForgetGateLast sets the value of the forgetGateLast property.
// A parameter that controls the internal order of the LSTM gates.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/forgetgatelast
func (g_ GraphLSTMDescriptor) SetForgetGateLast(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setForgetGateLast:"), value)
}

// A parameter that defines a bidirectional LSTM layer.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/bidirectional
func (g_ GraphLSTMDescriptor) Bidirectional() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("bidirectional"))
	return rv
}


// SetBidirectional sets the value of the bidirectional property.
// A parameter that defines a bidirectional LSTM layer.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/bidirectional
func (g_ GraphLSTMDescriptor) SetBidirectional(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBidirectional:"), value)
}

// A parameter that defines time direction of the input sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/reverse
func (g_ GraphLSTMDescriptor) Reverse() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("reverse"))
	return rv
}


// SetReverse sets the value of the reverse property.
// A parameter that defines time direction of the input sequence.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/reverse
func (g_ GraphLSTMDescriptor) SetReverse(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReverse:"), value)
}

// A parameter that defines the activation function used with the forget gate of the LSTM operation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/forgetgateactivation
func (g_ GraphLSTMDescriptor) ForgetGateActivation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("forgetGateActivation"))
	return rv
}


// SetForgetGateActivation sets the value of the forgetGateActivation property.
// A parameter that defines the activation function used with the forget gate of the LSTM operation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/forgetgateactivation
func (g_ GraphLSTMDescriptor) SetForgetGateActivation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setForgetGateActivation:"), value)
}

// A parameter that enables the LSTM layer to support training.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/training
func (g_ GraphLSTMDescriptor) Training() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("training"))
	return rv
}


// SetTraining sets the value of the training property.
// A parameter that enables the LSTM layer to support training.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/training
func (g_ GraphLSTMDescriptor) SetTraining(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTraining:"), value)
}

// A parameter that defines the activation function used with the cell gate of the LSTM operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/cellGateActivation
func (g_ GraphLSTMDescriptor) CellGateActivation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("cellGateActivation"))
	return rv
}


// SetCellGateActivation sets the value of the cellGateActivation property.
// A parameter that defines the activation function used with the cell gate of the LSTM operation.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/cellGateActivation
func (g_ GraphLSTMDescriptor) SetCellGateActivation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCellGateActivation:"), value)
}



