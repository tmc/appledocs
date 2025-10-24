// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronSoftPlusNode */


/* debug [class_header]: Header for MPSCNNNeuronSoftPlusNode */
// The class instance for the [CNNNeuronSoftPlusNode] class.
var (
	CNNNeuronSoftPlusNodeClass     _CNNNeuronSoftPlusNodeClass
	CNNNeuronSoftPlusNodeClassOnce sync.Once
)

func getCNNNeuronSoftPlusNodeClass() _CNNNeuronSoftPlusNodeClass {
	CNNNeuronSoftPlusNodeClassOnce.Do(func() {
		CNNNeuronSoftPlusNodeClass = _CNNNeuronSoftPlusNodeClass{objc.GetClass("MPSCNNNeuronSoftPlusNode")}
	})
	return CNNNeuronSoftPlusNodeClass
}

type _CNNNeuronSoftPlusNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronSoftPlusNode */
// An interface definition for the [CNNNeuronSoftPlusNode] class.
type ICNNNeuronSoftPlusNode interface {
	ICNNNeuronNode
	
/* debug [class_interface_properties]: Properties for CNNNeuronSoftPlusNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronSoftPlusNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronSoftPlusNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronSoftPlusNodeClass) Alloc() CNNNeuronSoftPlusNode {
	rv := objc.Send[CNNNeuronSoftPlusNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronSoftPlusNodeClass) New() CNNNeuronSoftPlusNode {
	rv := objc.Send[CNNNeuronSoftPlusNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronSoftPlusNode) Init() CNNNeuronSoftPlusNode {
	rv := objc.Send[CNNNeuronSoftPlusNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronSoftPlusNode) Autorelease() CNNNeuronSoftPlusNode {
	rv := objc.Send[CNNNeuronSoftPlusNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronSoftPlusNode creates a new CNNNeuronSoftPlusNode instance.
func NewCNNNeuronSoftPlusNode() CNNNeuronSoftPlusNode {
	return getCNNNeuronSoftPlusNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronSoftPlusNode */
// A representation of a parametric softplus neuron filter.


// A representation of a parametric softplus neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlusNode
type CNNNeuronSoftPlusNode struct {
	CNNNeuronNode
}

// CNNNeuronSoftPlusNodeFrom constructs a [CNNNeuronSoftPlusNode] from an unsafe.Pointer.
//
// A representation of a parametric softplus neuron filter.
func CNNNeuronSoftPlusNodeFrom(ptr unsafe.Pointer) CNNNeuronSoftPlusNode {
	return CNNNeuronSoftPlusNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronSoftPlusNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2921457-initwithsource
func NewCNNNeuronSoftPlusNodeWithSource(sourceNode IImageNode) CNNNeuronSoftPlusNode {
	instance := getCNNNeuronSoftPlusNodeClass().Alloc()
	rv := objc.Send[CNNNeuronSoftPlusNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronSoftPlusNodeWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2866413-initwithsource
func NewCNNNeuronSoftPlusNodeWithSourceAB(sourceNode IImageNode, a float32, b float32) CNNNeuronSoftPlusNode {
	instance := getCNNNeuronSoftPlusNodeClass().Alloc()
	rv := objc.Send[CNNNeuronSoftPlusNode](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronSoftPlusNodeWithSourceAB */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronSoftPlusNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2866450-nodewithsource
func (cc _CNNNeuronSoftPlusNodeClass) NodeWithSourceAB(sourceNode IImageNode, a float32, b float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceAB) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2921449-nodewithsource
func (cc _CNNNeuronSoftPlusNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronSoftPlusNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronSoftPlusNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronSoftPlusNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronSoftPlusNode */


