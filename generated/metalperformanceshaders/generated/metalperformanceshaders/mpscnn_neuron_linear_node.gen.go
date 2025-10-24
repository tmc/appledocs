// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronLinearNode */


/* debug [class_header]: Header for MPSCNNNeuronLinearNode */
// The class instance for the [CNNNeuronLinearNode] class.
var (
	CNNNeuronLinearNodeClass     _CNNNeuronLinearNodeClass
	CNNNeuronLinearNodeClassOnce sync.Once
)

func getCNNNeuronLinearNodeClass() _CNNNeuronLinearNodeClass {
	CNNNeuronLinearNodeClassOnce.Do(func() {
		CNNNeuronLinearNodeClass = _CNNNeuronLinearNodeClass{objc.GetClass("MPSCNNNeuronLinearNode")}
	})
	return CNNNeuronLinearNodeClass
}

type _CNNNeuronLinearNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronLinearNode */
// An interface definition for the [CNNNeuronLinearNode] class.
type ICNNNeuronLinearNode interface {
	ICNNNeuronNode
	
/* debug [class_interface_properties]: Properties for CNNNeuronLinearNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronLinearNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronLinearNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronLinearNodeClass) Alloc() CNNNeuronLinearNode {
	rv := objc.Send[CNNNeuronLinearNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronLinearNodeClass) New() CNNNeuronLinearNode {
	rv := objc.Send[CNNNeuronLinearNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronLinearNode) Init() CNNNeuronLinearNode {
	rv := objc.Send[CNNNeuronLinearNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronLinearNode) Autorelease() CNNNeuronLinearNode {
	rv := objc.Send[CNNNeuronLinearNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronLinearNode creates a new CNNNeuronLinearNode instance.
func NewCNNNeuronLinearNode() CNNNeuronLinearNode {
	return getCNNNeuronLinearNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronLinearNode */
// A representation of a linear neuron filter.


// A representation of a linear neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinearNode
type CNNNeuronLinearNode struct {
	CNNNeuronNode
}

// CNNNeuronLinearNodeFrom constructs a [CNNNeuronLinearNode] from an unsafe.Pointer.
//
// A representation of a linear neuron filter.
func CNNNeuronLinearNodeFrom(ptr unsafe.Pointer) CNNNeuronLinearNode {
	return CNNNeuronLinearNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronLinearNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlinearnode/2921456-initwithsource
func NewCNNNeuronLinearNodeWithSource(sourceNode IImageNode) CNNNeuronLinearNode {
	instance := getCNNNeuronLinearNodeClass().Alloc()
	rv := objc.Send[CNNNeuronLinearNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronLinearNodeWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlinearnode/2866495-initwithsource
func NewCNNNeuronLinearNodeWithSourceAB(sourceNode IImageNode, a float32, b float32) CNNNeuronLinearNode {
	instance := getCNNNeuronLinearNodeClass().Alloc()
	rv := objc.Send[CNNNeuronLinearNode](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronLinearNodeWithSourceAB */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronLinearNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlinearnode/2866445-nodewithsource
func (cc _CNNNeuronLinearNodeClass) NodeWithSourceAB(sourceNode IImageNode, a float32, b float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceAB) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlinearnode/2921450-nodewithsource
func (cc _CNNNeuronLinearNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronLinearNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronLinearNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronLinearNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronLinearNode */


