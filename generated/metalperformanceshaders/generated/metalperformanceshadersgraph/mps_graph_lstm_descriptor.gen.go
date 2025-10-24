// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphLSTMDescriptor */


/* debug [class_header]: Header for MPSGraphLSTMDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphLSTMDescriptor */
// An interface definition for the [GraphLSTMDescriptor] class.
type IGraphLSTMDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphLSTMDescriptor */
	// properties:
	Activation() GraphRNNActivation
	SetActivation(value GraphRNNActivation)
	Bidirectional() bool
	SetBidirectional(value bool)
	CellGateActivation() GraphRNNActivation
	SetCellGateActivation(value GraphRNNActivation)
	ForgetGateActivation() GraphRNNActivation
	SetForgetGateActivation(value GraphRNNActivation)
	ForgetGateLast() bool
	SetForgetGateLast(value bool)
	InputGateActivation() GraphRNNActivation
	SetInputGateActivation(value GraphRNNActivation)
	OutputGateActivation() GraphRNNActivation
	SetOutputGateActivation(value GraphRNNActivation)
	ProduceCell() bool
	SetProduceCell(value bool)
	Reverse() bool
	SetReverse(value bool)
	Training() bool
	SetTraining(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphLSTMDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphLSTMDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphLSTMDescriptorClass) Alloc() GraphLSTMDescriptor {
	rv := objc.Send[GraphLSTMDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphLSTMDescriptor */
// The class that defines the parameters for a long short-term memory (LSTM) operation.
//
// Use this descriptor with the following methods:


// The class that defines the parameters for a long short-term memory (LSTM) operation.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphLSTMDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphLSTMDescriptor */

// Creates an LSTM descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/descriptor
func (gc _GraphLSTMDescriptorClass) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Descriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphLSTMDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphLSTMDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphLSTMDescriptor */

// A parameter that defines the activation function used with the current cell value of the LSTM operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/activation
func (g_ GraphLSTMDescriptor) Activation() GraphRNNActivation {
	rv := objc.Send[GraphRNNActivation](g_.ID, objc.Sel("activation"))
	return rv
}/* debug [instance_properties/getter]: activation */


// A parameter that defines the activation function used with the current cell value of the LSTM operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/activation
func (g_ GraphLSTMDescriptor) SetActivation(value GraphRNNActivation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setActivation:"), value)
}/* debug [instance_properties/setter]: activation */


// A parameter that defines a bidirectional LSTM layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/bidirectional
func (g_ GraphLSTMDescriptor) Bidirectional() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("bidirectional"))
	return rv
}/* debug [instance_properties/getter]: bidirectional */


// A parameter that defines a bidirectional LSTM layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/bidirectional
func (g_ GraphLSTMDescriptor) SetBidirectional(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBidirectional:"), value)
}/* debug [instance_properties/setter]: bidirectional */


// A parameter that defines the activation function used with the cell gate of the LSTM operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/cellGateActivation
func (g_ GraphLSTMDescriptor) CellGateActivation() GraphRNNActivation {
	rv := objc.Send[GraphRNNActivation](g_.ID, objc.Sel("cellGateActivation"))
	return rv
}/* debug [instance_properties/getter]: cellGateActivation */


// A parameter that defines the activation function used with the cell gate of the LSTM operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/cellGateActivation
func (g_ GraphLSTMDescriptor) SetCellGateActivation(value GraphRNNActivation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCellGateActivation:"), value)
}/* debug [instance_properties/setter]: cellGateActivation */


// A parameter that defines the activation function used with the forget gate of the LSTM operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/forgetGateActivation
func (g_ GraphLSTMDescriptor) ForgetGateActivation() GraphRNNActivation {
	rv := objc.Send[GraphRNNActivation](g_.ID, objc.Sel("forgetGateActivation"))
	return rv
}/* debug [instance_properties/getter]: forgetGateActivation */


// A parameter that defines the activation function used with the forget gate of the LSTM operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/forgetGateActivation
func (g_ GraphLSTMDescriptor) SetForgetGateActivation(value GraphRNNActivation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setForgetGateActivation:"), value)
}/* debug [instance_properties/setter]: forgetGateActivation */


// A parameter that controls the internal order of the LSTM gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/forgetGateLast
func (g_ GraphLSTMDescriptor) ForgetGateLast() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("forgetGateLast"))
	return rv
}/* debug [instance_properties/getter]: forgetGateLast */


// A parameter that controls the internal order of the LSTM gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/forgetGateLast
func (g_ GraphLSTMDescriptor) SetForgetGateLast(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setForgetGateLast:"), value)
}/* debug [instance_properties/setter]: forgetGateLast */


// A parameter that defines the activation function used with the input gate of the LSTM operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/inputGateActivation
func (g_ GraphLSTMDescriptor) InputGateActivation() GraphRNNActivation {
	rv := objc.Send[GraphRNNActivation](g_.ID, objc.Sel("inputGateActivation"))
	return rv
}/* debug [instance_properties/getter]: inputGateActivation */


// A parameter that defines the activation function used with the input gate of the LSTM operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/inputGateActivation
func (g_ GraphLSTMDescriptor) SetInputGateActivation(value GraphRNNActivation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInputGateActivation:"), value)
}/* debug [instance_properties/setter]: inputGateActivation */


// A parameter that defines the activation function used with the output gate of the LSTM operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/outputGateActivation
func (g_ GraphLSTMDescriptor) OutputGateActivation() GraphRNNActivation {
	rv := objc.Send[GraphRNNActivation](g_.ID, objc.Sel("outputGateActivation"))
	return rv
}/* debug [instance_properties/getter]: outputGateActivation */


// A parameter that defines the activation function used with the output gate of the LSTM operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/outputGateActivation
func (g_ GraphLSTMDescriptor) SetOutputGateActivation(value GraphRNNActivation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputGateActivation:"), value)
}/* debug [instance_properties/setter]: outputGateActivation */


// A parameter that controls whether or not to return the output cell from the LSTM layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/produceCell
func (g_ GraphLSTMDescriptor) ProduceCell() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("produceCell"))
	return rv
}/* debug [instance_properties/getter]: produceCell */


// A parameter that controls whether or not to return the output cell from the LSTM layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/produceCell
func (g_ GraphLSTMDescriptor) SetProduceCell(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProduceCell:"), value)
}/* debug [instance_properties/setter]: produceCell */


// A parameter that defines time direction of the input sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/reverse
func (g_ GraphLSTMDescriptor) Reverse() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("reverse"))
	return rv
}/* debug [instance_properties/getter]: reverse */


// A parameter that defines time direction of the input sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/reverse
func (g_ GraphLSTMDescriptor) SetReverse(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReverse:"), value)
}/* debug [instance_properties/setter]: reverse */


// A parameter that enables the LSTM layer to support training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/training
func (g_ GraphLSTMDescriptor) Training() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("training"))
	return rv
}/* debug [instance_properties/getter]: training */


// A parameter that enables the LSTM layer to support training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/training
func (g_ GraphLSTMDescriptor) SetTraining(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTraining:"), value)
}/* debug [instance_properties/setter]: training */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphLSTMDescriptor */



