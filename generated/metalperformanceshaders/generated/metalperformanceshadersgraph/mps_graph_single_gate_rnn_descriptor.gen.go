// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphSingleGateRNNDescriptor */


/* debug [class_header]: Header for MPSGraphSingleGateRNNDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphSingleGateRNNDescriptor */
// An interface definition for the [GraphSingleGateRNNDescriptor] class.
type IGraphSingleGateRNNDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphSingleGateRNNDescriptor */
	// properties:
	Activation() GraphRNNActivation
	SetActivation(value GraphRNNActivation)
	Bidirectional() bool
	SetBidirectional(value bool)
	Reverse() bool
	SetReverse(value bool)
	Training() bool
	SetTraining(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphSingleGateRNNDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphSingleGateRNNDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphSingleGateRNNDescriptorClass) Alloc() GraphSingleGateRNNDescriptor {
	rv := objc.Send[GraphSingleGateRNNDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphSingleGateRNNDescriptor */
// The class that defines the parameters for a single gate RNN operation.
//
// Use this descriptor with the following methods:


// The class that defines the parameters for a single gate RNN operation.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphSingleGateRNNDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphSingleGateRNNDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphSingleGateRNNDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphSingleGateRNNDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphSingleGateRNNDescriptor */

// A parameter that defines the activation function to use with the RNN operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/activation
func (g_ GraphSingleGateRNNDescriptor) Activation() GraphRNNActivation {
	rv := objc.Send[GraphRNNActivation](g_.ID, objc.Sel("activation"))
	return rv
}/* debug [instance_properties/getter]: activation */


// A parameter that defines the activation function to use with the RNN operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/activation
func (g_ GraphSingleGateRNNDescriptor) SetActivation(value GraphRNNActivation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setActivation:"), value)
}/* debug [instance_properties/setter]: activation */


// A parameter that defines a bidirectional RNN layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/bidirectional
func (g_ GraphSingleGateRNNDescriptor) Bidirectional() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("bidirectional"))
	return rv
}/* debug [instance_properties/getter]: bidirectional */


// A parameter that defines a bidirectional RNN layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/bidirectional
func (g_ GraphSingleGateRNNDescriptor) SetBidirectional(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBidirectional:"), value)
}/* debug [instance_properties/setter]: bidirectional */


// A parameter that defines time direction of the input sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/reverse
func (g_ GraphSingleGateRNNDescriptor) Reverse() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("reverse"))
	return rv
}/* debug [instance_properties/getter]: reverse */


// A parameter that defines time direction of the input sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/reverse
func (g_ GraphSingleGateRNNDescriptor) SetReverse(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReverse:"), value)
}/* debug [instance_properties/setter]: reverse */


// A parameter that makes the RNN layer support training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/training
func (g_ GraphSingleGateRNNDescriptor) Training() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("training"))
	return rv
}/* debug [instance_properties/getter]: training */


// A parameter that makes the RNN layer support training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsinglegaternndescriptor/training
func (g_ GraphSingleGateRNNDescriptor) SetTraining(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTraining:"), value)
}/* debug [instance_properties/setter]: training */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphSingleGateRNNDescriptor */



