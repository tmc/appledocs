// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronPReLUNode */


/* debug [class_header]: Header for MPSCNNNeuronPReLUNode */
// The class instance for the [CNNNeuronPReLUNode] class.
var (
	CNNNeuronPReLUNodeClass     _CNNNeuronPReLUNodeClass
	CNNNeuronPReLUNodeClassOnce sync.Once
)

func getCNNNeuronPReLUNodeClass() _CNNNeuronPReLUNodeClass {
	CNNNeuronPReLUNodeClassOnce.Do(func() {
		CNNNeuronPReLUNodeClass = _CNNNeuronPReLUNodeClass{objc.GetClass("MPSCNNNeuronPReLUNode")}
	})
	return CNNNeuronPReLUNodeClass
}

type _CNNNeuronPReLUNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronPReLUNode */
// An interface definition for the [CNNNeuronPReLUNode] class.
type ICNNNeuronPReLUNode interface {
	ICNNNeuronNode
	
/* debug [class_interface_properties]: Properties for CNNNeuronPReLUNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronPReLUNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronPReLUNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronPReLUNodeClass) Alloc() CNNNeuronPReLUNode {
	rv := objc.Send[CNNNeuronPReLUNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronPReLUNodeClass) New() CNNNeuronPReLUNode {
	rv := objc.Send[CNNNeuronPReLUNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronPReLUNode) Init() CNNNeuronPReLUNode {
	rv := objc.Send[CNNNeuronPReLUNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronPReLUNode) Autorelease() CNNNeuronPReLUNode {
	rv := objc.Send[CNNNeuronPReLUNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronPReLUNode creates a new CNNNeuronPReLUNode instance.
func NewCNNNeuronPReLUNode() CNNNeuronPReLUNode {
	return getCNNNeuronPReLUNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronPReLUNode */
// A representation a PReLU neuron filter.


// A representation a PReLU neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPReLUNode
type CNNNeuronPReLUNode struct {
	CNNNeuronNode
}

// CNNNeuronPReLUNodeFrom constructs a [CNNNeuronPReLUNode] from an unsafe.Pointer.
//
// A representation a PReLU neuron filter.
func CNNNeuronPReLUNodeFrom(ptr unsafe.Pointer) CNNNeuronPReLUNode {
	return CNNNeuronPReLUNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronPReLUNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronprelunode/2921595-initwithsource
func NewCNNNeuronPReLUNodeWithSourceAData(sourceNode IImageNode, aData foundation.Data) CNNNeuronPReLUNode {
	instance := getCNNNeuronPReLUNodeClass().Alloc()
	rv := objc.Send[CNNNeuronPReLUNode](instance.ID, objc.Sel("initWithSource:aData:"), sourceNode, aData)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronPReLUNodeWithSourceAData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronPReLUNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronprelunode/2921597-nodewithsource
func (cc _CNNNeuronPReLUNodeClass) NodeWithSourceAData(sourceNode IImageNode, aData foundation.Data) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:aData:"), sourceNode, aData)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceAData) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronPReLUNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronPReLUNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronPReLUNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronPReLUNode */


