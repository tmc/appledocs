// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphGRUDescriptor */


/* debug [class_header]: Header for MPSGraphGRUDescriptor */
// The class instance for the [GraphGRUDescriptor] class.
var (
	GraphGRUDescriptorClass     _GraphGRUDescriptorClass
	GraphGRUDescriptorClassOnce sync.Once
)

func getGraphGRUDescriptorClass() _GraphGRUDescriptorClass {
	GraphGRUDescriptorClassOnce.Do(func() {
		GraphGRUDescriptorClass = _GraphGRUDescriptorClass{objc.GetClass("MPSGraphGRUDescriptor")}
	})
	return GraphGRUDescriptorClass
}

type _GraphGRUDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphGRUDescriptor */
// An interface definition for the [GraphGRUDescriptor] class.
type IGraphGRUDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphGRUDescriptor */
	// properties:
	Bidirectional() bool
	SetBidirectional(value bool)
	FlipZ() bool
	SetFlipZ(value bool)
	OutputGateActivation() GraphRNNActivation
	SetOutputGateActivation(value GraphRNNActivation)
	ResetAfter() bool
	SetResetAfter(value bool)
	ResetGateActivation() GraphRNNActivation
	SetResetGateActivation(value GraphRNNActivation)
	ResetGateFirst() bool
	SetResetGateFirst(value bool)
	Reverse() bool
	SetReverse(value bool)
	Training() bool
	SetTraining(value bool)
	UpdateGateActivation() GraphRNNActivation
	SetUpdateGateActivation(value GraphRNNActivation)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphGRUDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphGRUDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphGRUDescriptorClass) Alloc() GraphGRUDescriptor {
	rv := objc.Send[GraphGRUDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphGRUDescriptorClass) New() GraphGRUDescriptor {
	rv := objc.Send[GraphGRUDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphGRUDescriptor) Init() GraphGRUDescriptor {
	rv := objc.Send[GraphGRUDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphGRUDescriptor) Autorelease() GraphGRUDescriptor {
	rv := objc.Send[GraphGRUDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphGRUDescriptor creates a new GraphGRUDescriptor instance.
func NewGraphGRUDescriptor() GraphGRUDescriptor {
	return getGraphGRUDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphGRUDescriptor */
// The class that defines the parameters for a gated recurrent unit (GRU) operation.
//
// Use this descriptor with the following methods:


// The class that defines the parameters for a gated recurrent unit (GRU) operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor
type GraphGRUDescriptor struct {
	GraphObject
}

// GraphGRUDescriptorFrom constructs a [GraphGRUDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a gated recurrent unit (GRU) operation.
func GraphGRUDescriptorFrom(ptr unsafe.Pointer) GraphGRUDescriptor {
	return GraphGRUDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphGRUDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphGRUDescriptor */

// Creates an GRU descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/descriptor
func (gc _GraphGRUDescriptorClass) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Descriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphGRUDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphGRUDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphGRUDescriptor */

// A parameter that defines a bidirectional GRU layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/bidirectional
func (g_ GraphGRUDescriptor) Bidirectional() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("bidirectional"))
	return rv
}/* debug [instance_properties/getter]: bidirectional */


// A parameter that defines a bidirectional GRU layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/bidirectional
func (g_ GraphGRUDescriptor) SetBidirectional(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBidirectional:"), value)
}/* debug [instance_properties/setter]: bidirectional */


// A parameter that chooses between two variants for the final output computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/flipZ
func (g_ GraphGRUDescriptor) FlipZ() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("flipZ"))
	return rv
}/* debug [instance_properties/getter]: flipZ */


// A parameter that chooses between two variants for the final output computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/flipZ
func (g_ GraphGRUDescriptor) SetFlipZ(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setFlipZ:"), value)
}/* debug [instance_properties/setter]: flipZ */


// A parameter that defines the activation function to use with the output-gate of the GRU operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/outputGateActivation
func (g_ GraphGRUDescriptor) OutputGateActivation() GraphRNNActivation {
	rv := objc.Send[GraphRNNActivation](g_.ID, objc.Sel("outputGateActivation"))
	return rv
}/* debug [instance_properties/getter]: outputGateActivation */


// A parameter that defines the activation function to use with the output-gate of the GRU operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/outputGateActivation
func (g_ GraphGRUDescriptor) SetOutputGateActivation(value GraphRNNActivation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputGateActivation:"), value)
}/* debug [instance_properties/setter]: outputGateActivation */


// A parameter that chooses between two variants for the reset gate computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/resetAfter
func (g_ GraphGRUDescriptor) ResetAfter() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("resetAfter"))
	return rv
}/* debug [instance_properties/getter]: resetAfter */


// A parameter that chooses between two variants for the reset gate computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/resetAfter
func (g_ GraphGRUDescriptor) SetResetAfter(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResetAfter:"), value)
}/* debug [instance_properties/setter]: resetAfter */


// A parameter that defines the activation function to use with the reset-gate of the GRU operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/resetGateActivation
func (g_ GraphGRUDescriptor) ResetGateActivation() GraphRNNActivation {
	rv := objc.Send[GraphRNNActivation](g_.ID, objc.Sel("resetGateActivation"))
	return rv
}/* debug [instance_properties/getter]: resetGateActivation */


// A parameter that defines the activation function to use with the reset-gate of the GRU operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/resetGateActivation
func (g_ GraphGRUDescriptor) SetResetGateActivation(value GraphRNNActivation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResetGateActivation:"), value)
}/* debug [instance_properties/setter]: resetGateActivation */


// A parameter that controls the internal order of the GRU gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/resetGateFirst
func (g_ GraphGRUDescriptor) ResetGateFirst() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("resetGateFirst"))
	return rv
}/* debug [instance_properties/getter]: resetGateFirst */


// A parameter that controls the internal order of the GRU gates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/resetGateFirst
func (g_ GraphGRUDescriptor) SetResetGateFirst(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResetGateFirst:"), value)
}/* debug [instance_properties/setter]: resetGateFirst */


// A parameter that defines the time direction of the input sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/reverse
func (g_ GraphGRUDescriptor) Reverse() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("reverse"))
	return rv
}/* debug [instance_properties/getter]: reverse */


// A parameter that defines the time direction of the input sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/reverse
func (g_ GraphGRUDescriptor) SetReverse(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReverse:"), value)
}/* debug [instance_properties/setter]: reverse */


// A parameter that enables the GRU layer to support training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/training
func (g_ GraphGRUDescriptor) Training() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("training"))
	return rv
}/* debug [instance_properties/getter]: training */


// A parameter that enables the GRU layer to support training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/training
func (g_ GraphGRUDescriptor) SetTraining(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTraining:"), value)
}/* debug [instance_properties/setter]: training */


// A parameter that defines the activation function to use with the update-gate of the GRU operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/updateGateActivation
func (g_ GraphGRUDescriptor) UpdateGateActivation() GraphRNNActivation {
	rv := objc.Send[GraphRNNActivation](g_.ID, objc.Sel("updateGateActivation"))
	return rv
}/* debug [instance_properties/getter]: updateGateActivation */


// A parameter that defines the activation function to use with the update-gate of the GRU operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/updateGateActivation
func (g_ GraphGRUDescriptor) SetUpdateGateActivation(value GraphRNNActivation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUpdateGateActivation:"), value)
}/* debug [instance_properties/setter]: updateGateActivation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphGRUDescriptor */



