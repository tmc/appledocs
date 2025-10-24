// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronTanHNode */


/* debug [class_header]: Header for MPSCNNNeuronTanHNode */
// The class instance for the [CNNNeuronTanHNode] class.
var (
	CNNNeuronTanHNodeClass     _CNNNeuronTanHNodeClass
	CNNNeuronTanHNodeClassOnce sync.Once
)

func getCNNNeuronTanHNodeClass() _CNNNeuronTanHNodeClass {
	CNNNeuronTanHNodeClassOnce.Do(func() {
		CNNNeuronTanHNodeClass = _CNNNeuronTanHNodeClass{objc.GetClass("MPSCNNNeuronTanHNode")}
	})
	return CNNNeuronTanHNodeClass
}

type _CNNNeuronTanHNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronTanHNode */
// An interface definition for the [CNNNeuronTanHNode] class.
type ICNNNeuronTanHNode interface {
	ICNNNeuronNode
	
/* debug [class_interface_properties]: Properties for CNNNeuronTanHNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronTanHNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronTanHNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronTanHNodeClass) Alloc() CNNNeuronTanHNode {
	rv := objc.Send[CNNNeuronTanHNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronTanHNodeClass) New() CNNNeuronTanHNode {
	rv := objc.Send[CNNNeuronTanHNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronTanHNode) Init() CNNNeuronTanHNode {
	rv := objc.Send[CNNNeuronTanHNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronTanHNode) Autorelease() CNNNeuronTanHNode {
	rv := objc.Send[CNNNeuronTanHNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronTanHNode creates a new CNNNeuronTanHNode instance.
func NewCNNNeuronTanHNode() CNNNeuronTanHNode {
	return getCNNNeuronTanHNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronTanHNode */
// A representation of a hyperbolic tangent neuron filter.


// A representation of a hyperbolic tangent neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanHNode
type CNNNeuronTanHNode struct {
	CNNNeuronNode
}

// CNNNeuronTanHNodeFrom constructs a [CNNNeuronTanHNode] from an unsafe.Pointer.
//
// A representation of a hyperbolic tangent neuron filter.
func CNNNeuronTanHNodeFrom(ptr unsafe.Pointer) CNNNeuronTanHNode {
	return CNNNeuronTanHNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronTanHNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2921465-initwithsource
func NewCNNNeuronTanHNodeWithSource(sourceNode IImageNode) CNNNeuronTanHNode {
	instance := getCNNNeuronTanHNodeClass().Alloc()
	rv := objc.Send[CNNNeuronTanHNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronTanHNodeWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2866481-initwithsource
func NewCNNNeuronTanHNodeWithSourceAB(sourceNode IImageNode, a float32, b float32) CNNNeuronTanHNode {
	instance := getCNNNeuronTanHNodeClass().Alloc()
	rv := objc.Send[CNNNeuronTanHNode](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronTanHNodeWithSourceAB */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronTanHNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2866416-nodewithsource
func (cc _CNNNeuronTanHNodeClass) NodeWithSourceAB(sourceNode IImageNode, a float32, b float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceAB) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanhnode/2921451-nodewithsource
func (cc _CNNNeuronTanHNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronTanHNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronTanHNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronTanHNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronTanHNode */


