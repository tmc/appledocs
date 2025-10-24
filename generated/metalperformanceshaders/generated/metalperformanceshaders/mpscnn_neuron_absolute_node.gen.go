// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronAbsoluteNode */


/* debug [class_header]: Header for MPSCNNNeuronAbsoluteNode */
// The class instance for the [CNNNeuronAbsoluteNode] class.
var (
	CNNNeuronAbsoluteNodeClass     _CNNNeuronAbsoluteNodeClass
	CNNNeuronAbsoluteNodeClassOnce sync.Once
)

func getCNNNeuronAbsoluteNodeClass() _CNNNeuronAbsoluteNodeClass {
	CNNNeuronAbsoluteNodeClassOnce.Do(func() {
		CNNNeuronAbsoluteNodeClass = _CNNNeuronAbsoluteNodeClass{objc.GetClass("MPSCNNNeuronAbsoluteNode")}
	})
	return CNNNeuronAbsoluteNodeClass
}

type _CNNNeuronAbsoluteNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronAbsoluteNode */
// An interface definition for the [CNNNeuronAbsoluteNode] class.
type ICNNNeuronAbsoluteNode interface {
	ICNNNeuronNode
	
/* debug [class_interface_properties]: Properties for CNNNeuronAbsoluteNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronAbsoluteNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronAbsoluteNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronAbsoluteNodeClass) Alloc() CNNNeuronAbsoluteNode {
	rv := objc.Send[CNNNeuronAbsoluteNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronAbsoluteNodeClass) New() CNNNeuronAbsoluteNode {
	rv := objc.Send[CNNNeuronAbsoluteNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronAbsoluteNode) Init() CNNNeuronAbsoluteNode {
	rv := objc.Send[CNNNeuronAbsoluteNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronAbsoluteNode) Autorelease() CNNNeuronAbsoluteNode {
	rv := objc.Send[CNNNeuronAbsoluteNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronAbsoluteNode creates a new CNNNeuronAbsoluteNode instance.
func NewCNNNeuronAbsoluteNode() CNNNeuronAbsoluteNode {
	return getCNNNeuronAbsoluteNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronAbsoluteNode */
// A representation of an absolute neuron filter.


// A representation of an absolute neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronAbsoluteNode
type CNNNeuronAbsoluteNode struct {
	CNNNeuronNode
}

// CNNNeuronAbsoluteNodeFrom constructs a [CNNNeuronAbsoluteNode] from an unsafe.Pointer.
//
// A representation of an absolute neuron filter.
func CNNNeuronAbsoluteNodeFrom(ptr unsafe.Pointer) CNNNeuronAbsoluteNode {
	return CNNNeuronAbsoluteNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronAbsoluteNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronabsolutenode/2921448-initwithsource
func NewCNNNeuronAbsoluteNodeWithSource(sourceNode IImageNode) CNNNeuronAbsoluteNode {
	instance := getCNNNeuronAbsoluteNodeClass().Alloc()
	rv := objc.Send[CNNNeuronAbsoluteNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronAbsoluteNodeWithSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronAbsoluteNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronabsolutenode/2866431-nodewithsource
func (cc _CNNNeuronAbsoluteNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronAbsoluteNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronAbsoluteNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronAbsoluteNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronAbsoluteNode */


