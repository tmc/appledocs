// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNFullyConnectedNode */


/* debug [class_header]: Header for MPSCNNFullyConnectedNode */
// The class instance for the [CNNFullyConnectedNode] class.
var (
	CNNFullyConnectedNodeClass     _CNNFullyConnectedNodeClass
	CNNFullyConnectedNodeClassOnce sync.Once
)

func getCNNFullyConnectedNodeClass() _CNNFullyConnectedNodeClass {
	CNNFullyConnectedNodeClassOnce.Do(func() {
		CNNFullyConnectedNodeClass = _CNNFullyConnectedNodeClass{objc.GetClass("MPSCNNFullyConnectedNode")}
	})
	return CNNFullyConnectedNodeClass
}

type _CNNFullyConnectedNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNFullyConnectedNode */
// An interface definition for the [CNNFullyConnectedNode] class.
type ICNNFullyConnectedNode interface {
	ICNNConvolutionNode
	
/* debug [class_interface_properties]: Properties for CNNFullyConnectedNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNFullyConnectedNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNFullyConnectedNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNFullyConnectedNodeClass) Alloc() CNNFullyConnectedNode {
	rv := objc.Send[CNNFullyConnectedNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNFullyConnectedNodeClass) New() CNNFullyConnectedNode {
	rv := objc.Send[CNNFullyConnectedNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNFullyConnectedNode) Init() CNNFullyConnectedNode {
	rv := objc.Send[CNNFullyConnectedNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNFullyConnectedNode) Autorelease() CNNFullyConnectedNode {
	rv := objc.Send[CNNFullyConnectedNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNFullyConnectedNode creates a new CNNFullyConnectedNode instance.
func NewCNNFullyConnectedNode() CNNFullyConnectedNode {
	return getCNNFullyConnectedNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNFullyConnectedNode */
// A representation of a fully connected convolution layer, also known as an inner product layer.


// A representation of a fully connected convolution layer, also known as an inner product layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedNode
type CNNFullyConnectedNode struct {
	CNNConvolutionNode
}

// CNNFullyConnectedNodeFrom constructs a [CNNFullyConnectedNode] from an unsafe.Pointer.
//
// A representation of a fully connected convolution layer, also known as an inner product layer.
func CNNFullyConnectedNodeFrom(ptr unsafe.Pointer) CNNFullyConnectedNode {
	return CNNFullyConnectedNode{
		CNNConvolutionNode: CNNConvolutionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNFullyConnectedNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectednode/2866412-initwithsource
func NewCNNFullyConnectedNodeWithSourceWeights(sourceNode IImageNode, weights unsafe.Pointer) CNNFullyConnectedNode {
	instance := getCNNFullyConnectedNodeClass().Alloc()
	rv := objc.Send[CNNFullyConnectedNode](instance.ID, objc.Sel("initWithSource:weights:"), sourceNode, weights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNFullyConnectedNodeWithSourceWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNFullyConnectedNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectednode/2866458-nodewithsource
func (cc _CNNFullyConnectedNodeClass) NodeWithSourceWeights(sourceNode IImageNode, weights unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:weights:"), sourceNode, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceWeights) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNFullyConnectedNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNFullyConnectedNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNFullyConnectedNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNFullyConnectedNode */


