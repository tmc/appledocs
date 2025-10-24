// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronReLUNode */


/* debug [class_header]: Header for MPSCNNNeuronReLUNode */
// The class instance for the [CNNNeuronReLUNode] class.
var (
	CNNNeuronReLUNodeClass     _CNNNeuronReLUNodeClass
	CNNNeuronReLUNodeClassOnce sync.Once
)

func getCNNNeuronReLUNodeClass() _CNNNeuronReLUNodeClass {
	CNNNeuronReLUNodeClassOnce.Do(func() {
		CNNNeuronReLUNodeClass = _CNNNeuronReLUNodeClass{objc.GetClass("MPSCNNNeuronReLUNode")}
	})
	return CNNNeuronReLUNodeClass
}

type _CNNNeuronReLUNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronReLUNode */
// An interface definition for the [CNNNeuronReLUNode] class.
type ICNNNeuronReLUNode interface {
	ICNNNeuronNode
	
/* debug [class_interface_properties]: Properties for CNNNeuronReLUNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronReLUNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronReLUNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronReLUNodeClass) Alloc() CNNNeuronReLUNode {
	rv := objc.Send[CNNNeuronReLUNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronReLUNodeClass) New() CNNNeuronReLUNode {
	rv := objc.Send[CNNNeuronReLUNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronReLUNode) Init() CNNNeuronReLUNode {
	rv := objc.Send[CNNNeuronReLUNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronReLUNode) Autorelease() CNNNeuronReLUNode {
	rv := objc.Send[CNNNeuronReLUNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronReLUNode creates a new CNNNeuronReLUNode instance.
func NewCNNNeuronReLUNode() CNNNeuronReLUNode {
	return getCNNNeuronReLUNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronReLUNode */
// A representation a ReLU neuron filter.


// A representation a ReLU neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNode
type CNNNeuronReLUNode struct {
	CNNNeuronNode
}

// CNNNeuronReLUNodeFrom constructs a [CNNNeuronReLUNode] from an unsafe.Pointer.
//
// A representation a ReLU neuron filter.
func CNNNeuronReLUNodeFrom(ptr unsafe.Pointer) CNNNeuronReLUNode {
	return CNNNeuronReLUNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronReLUNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunode/2921464-initwithsource
func NewCNNNeuronReLUNodeWithSource(sourceNode IImageNode) CNNNeuronReLUNode {
	instance := getCNNNeuronReLUNodeClass().Alloc()
	rv := objc.Send[CNNNeuronReLUNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronReLUNodeWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunode/2921462-initwithsource
func NewCNNNeuronReLUNodeWithSourceA(sourceNode IImageNode, a float32) CNNNeuronReLUNode {
	instance := getCNNNeuronReLUNodeClass().Alloc()
	rv := objc.Send[CNNNeuronReLUNode](instance.ID, objc.Sel("initWithSource:a:"), sourceNode, a)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronReLUNodeWithSourceA */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronReLUNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunode/2866494-nodewithsource
func (cc _CNNNeuronReLUNodeClass) NodeWithSourceA(sourceNode IImageNode, a float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:a:"), sourceNode, a)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceA) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunode/2921460-nodewithsource
func (cc _CNNNeuronReLUNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronReLUNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronReLUNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronReLUNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronReLUNode */


