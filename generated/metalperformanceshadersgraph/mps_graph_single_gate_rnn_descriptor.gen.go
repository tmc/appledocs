// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphSingleGateRNNDescriptor] class.
var (
	GraphSingleGateRNNDescriptorClass     _GraphSingleGateRNNDescriptorClass
	GraphSingleGateRNNDescriptorClassOnce sync.Once
)

func getGraphSingleGateRNNDescriptorClass() _GraphSingleGateRNNDescriptorClass {
	GraphSingleGateRNNDescriptorClassOnce.Do(func() {
		GraphSingleGateRNNDescriptorClass = _GraphSingleGateRNNDescriptorClass{objc.GetClass("MPSGraphSingleGateRNNDescriptor")}
	})
	return GraphSingleGateRNNDescriptorClass
}

type _GraphSingleGateRNNDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphSingleGateRNNDescriptor] class.
type IGraphSingleGateRNNDescriptor interface {
	IGraphObject
	Bidirectional() bool
	SetBidirectional(value bool)
	Activation() GraphRNNActivation
	SetActivation(value IGraphRNNActivation)
	Reverse() bool
	SetReverse(value bool)
	Training() bool
	SetTraining(value bool)
}

// The class that defines the parameters for a single gate RNN operation.
//
// Use this descriptor with the following methods:
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSingleGateRNNDescriptor
type GraphSingleGateRNNDescriptor struct {
	GraphObject
}

// GraphSingleGateRNNDescriptorFrom constructs a [GraphSingleGateRNNDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a single gate RNN operation.
func GraphSingleGateRNNDescriptorFrom(ptr unsafe.Pointer) GraphSingleGateRNNDescriptor {
	return GraphSingleGateRNNDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphSingleGateRNNDescriptorClass) Alloc() GraphSingleGateRNNDescriptor {
	rv := objc.Send[GraphSingleGateRNNDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphSingleGateRNNDescriptorClass) New() GraphSingleGateRNNDescriptor {
	rv := objc.Send[GraphSingleGateRNNDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphSingleGateRNNDescriptor) Init() GraphSingleGateRNNDescriptor {
	rv := objc.Send[GraphSingleGateRNNDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphSingleGateRNNDescriptor) Autorelease() GraphSingleGateRNNDescriptor {
	rv := objc.Send[GraphSingleGateRNNDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphSingleGateRNNDescriptor creates a new GraphSingleGateRNNDescriptor instance.
func NewGraphSingleGateRNNDescriptor() GraphSingleGateRNNDescriptor {
	return getGraphSingleGateRNNDescriptorClass().New()
}


// Creates a single gate RNN descriptor with default values.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSingleGateRNNDescriptor/descriptor
func (gc _GraphSingleGateRNNDescriptorClass) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptor"))
	return rv
}

// A parameter that defines a bidirectional RNN layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSingleGateRNNDescriptor/bidirectional
func (g_ GraphSingleGateRNNDescriptor) Bidirectional() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("bidirectional"))
	return rv
}


// SetBidirectional sets the value of the bidirectional property.
// A parameter that defines a bidirectional RNN layer.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSingleGateRNNDescriptor/bidirectional
func (g_ GraphSingleGateRNNDescriptor) SetBidirectional(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBidirectional:"), value)
}

// A parameter that defines the activation function to use with the RNN operation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/activation
func (g_ GraphSingleGateRNNDescriptor) Activation() GraphRNNActivation {
	rv := objc.Send[GraphRNNActivation](g_.ID, objc.Sel("activation"))
	return rv
}


// SetActivation sets the value of the activation property.
// A parameter that defines the activation function to use with the RNN operation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/activation
func (g_ GraphSingleGateRNNDescriptor) SetActivation(value IGraphRNNActivation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setActivation:"), value)
}

// A parameter that defines time direction of the input sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/reverse
func (g_ GraphSingleGateRNNDescriptor) Reverse() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("reverse"))
	return rv
}


// SetReverse sets the value of the reverse property.
// A parameter that defines time direction of the input sequence.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/reverse
func (g_ GraphSingleGateRNNDescriptor) SetReverse(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReverse:"), value)
}

// A parameter that makes the RNN layer support training.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/training
func (g_ GraphSingleGateRNNDescriptor) Training() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("training"))
	return rv
}


// SetTraining sets the value of the training property.
// A parameter that makes the RNN layer support training.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/training
func (g_ GraphSingleGateRNNDescriptor) SetTraining(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTraining:"), value)
}



