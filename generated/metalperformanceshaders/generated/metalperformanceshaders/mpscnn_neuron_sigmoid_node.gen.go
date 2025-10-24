// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronSigmoidNode */


/* debug [class_header]: Header for MPSCNNNeuronSigmoidNode */
// The class instance for the [CNNNeuronSigmoidNode] class.
var (
	CNNNeuronSigmoidNodeClass     _CNNNeuronSigmoidNodeClass
	CNNNeuronSigmoidNodeClassOnce sync.Once
)

func getCNNNeuronSigmoidNodeClass() _CNNNeuronSigmoidNodeClass {
	CNNNeuronSigmoidNodeClassOnce.Do(func() {
		CNNNeuronSigmoidNodeClass = _CNNNeuronSigmoidNodeClass{objc.GetClass("MPSCNNNeuronSigmoidNode")}
	})
	return CNNNeuronSigmoidNodeClass
}

type _CNNNeuronSigmoidNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronSigmoidNode */
// An interface definition for the [CNNNeuronSigmoidNode] class.
type ICNNNeuronSigmoidNode interface {
	ICNNNeuronNode
	
/* debug [class_interface_properties]: Properties for CNNNeuronSigmoidNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronSigmoidNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronSigmoidNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronSigmoidNodeClass) Alloc() CNNNeuronSigmoidNode {
	rv := objc.Send[CNNNeuronSigmoidNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronSigmoidNodeClass) New() CNNNeuronSigmoidNode {
	rv := objc.Send[CNNNeuronSigmoidNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronSigmoidNode) Init() CNNNeuronSigmoidNode {
	rv := objc.Send[CNNNeuronSigmoidNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronSigmoidNode) Autorelease() CNNNeuronSigmoidNode {
	rv := objc.Send[CNNNeuronSigmoidNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronSigmoidNode creates a new CNNNeuronSigmoidNode instance.
func NewCNNNeuronSigmoidNode() CNNNeuronSigmoidNode {
	return getCNNNeuronSigmoidNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronSigmoidNode */
// A representation of a sigmoid neuron filter.


// A representation of a sigmoid neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSigmoidNode
type CNNNeuronSigmoidNode struct {
	CNNNeuronNode
}

// CNNNeuronSigmoidNodeFrom constructs a [CNNNeuronSigmoidNode] from an unsafe.Pointer.
//
// A representation of a sigmoid neuron filter.
func CNNNeuronSigmoidNodeFrom(ptr unsafe.Pointer) CNNNeuronSigmoidNode {
	return CNNNeuronSigmoidNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronSigmoidNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsigmoidnode/2921458-initwithsource
func NewCNNNeuronSigmoidNodeWithSource(sourceNode IImageNode) CNNNeuronSigmoidNode {
	instance := getCNNNeuronSigmoidNodeClass().Alloc()
	rv := objc.Send[CNNNeuronSigmoidNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronSigmoidNodeWithSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronSigmoidNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsigmoidnode/2866467-nodewithsource
func (cc _CNNNeuronSigmoidNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronSigmoidNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronSigmoidNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronSigmoidNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronSigmoidNode */


