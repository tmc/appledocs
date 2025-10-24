// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronLogarithmNode */


/* debug [class_header]: Header for MPSCNNNeuronLogarithmNode */
// The class instance for the [CNNNeuronLogarithmNode] class.
var (
	CNNNeuronLogarithmNodeClass     _CNNNeuronLogarithmNodeClass
	CNNNeuronLogarithmNodeClassOnce sync.Once
)

func getCNNNeuronLogarithmNodeClass() _CNNNeuronLogarithmNodeClass {
	CNNNeuronLogarithmNodeClassOnce.Do(func() {
		CNNNeuronLogarithmNodeClass = _CNNNeuronLogarithmNodeClass{objc.GetClass("MPSCNNNeuronLogarithmNode")}
	})
	return CNNNeuronLogarithmNodeClass
}

type _CNNNeuronLogarithmNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronLogarithmNode */
// An interface definition for the [CNNNeuronLogarithmNode] class.
type ICNNNeuronLogarithmNode interface {
	ICNNNeuronNode
	
/* debug [class_interface_properties]: Properties for CNNNeuronLogarithmNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronLogarithmNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronLogarithmNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronLogarithmNodeClass) Alloc() CNNNeuronLogarithmNode {
	rv := objc.Send[CNNNeuronLogarithmNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronLogarithmNodeClass) New() CNNNeuronLogarithmNode {
	rv := objc.Send[CNNNeuronLogarithmNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronLogarithmNode) Init() CNNNeuronLogarithmNode {
	rv := objc.Send[CNNNeuronLogarithmNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronLogarithmNode) Autorelease() CNNNeuronLogarithmNode {
	rv := objc.Send[CNNNeuronLogarithmNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronLogarithmNode creates a new CNNNeuronLogarithmNode instance.
func NewCNNNeuronLogarithmNode() CNNNeuronLogarithmNode {
	return getCNNNeuronLogarithmNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronLogarithmNode */
// A representation of a logarithm neuron filter.


// A representation of a logarithm neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithmNode
type CNNNeuronLogarithmNode struct {
	CNNNeuronNode
}

// CNNNeuronLogarithmNodeFrom constructs a [CNNNeuronLogarithmNode] from an unsafe.Pointer.
//
// A representation of a logarithm neuron filter.
func CNNNeuronLogarithmNodeFrom(ptr unsafe.Pointer) CNNNeuronLogarithmNode {
	return CNNNeuronLogarithmNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronLogarithmNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlogarithmnode/2951949-initwithsource
func NewCNNNeuronLogarithmNodeWithSource(sourceNode IImageNode) CNNNeuronLogarithmNode {
	instance := getCNNNeuronLogarithmNodeClass().Alloc()
	rv := objc.Send[CNNNeuronLogarithmNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronLogarithmNodeWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlogarithmnode/2951939-initwithsource
func NewCNNNeuronLogarithmNodeWithSourceABC(sourceNode IImageNode, a float32, b float32, c float32) CNNNeuronLogarithmNode {
	instance := getCNNNeuronLogarithmNodeClass().Alloc()
	rv := objc.Send[CNNNeuronLogarithmNode](instance.ID, objc.Sel("initWithSource:a:b:c:"), sourceNode, a, b, c)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronLogarithmNodeWithSourceABC */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronLogarithmNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlogarithmnode/2951931-nodewithsource
func (cc _CNNNeuronLogarithmNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlogarithmnode/2951943-nodewithsource
func (cc _CNNNeuronLogarithmNodeClass) NodeWithSourceABC(sourceNode IImageNode, a float32, b float32, c float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:b:c:"), sourceNode, a, b, c)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceABC) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronLogarithmNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronLogarithmNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronLogarithmNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronLogarithmNode */


