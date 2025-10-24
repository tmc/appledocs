// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNFullyConnectedGradientNode */


/* debug [class_header]: Header for MPSCNNFullyConnectedGradientNode */
// The class instance for the [CNNFullyConnectedGradientNode] class.
var (
	CNNFullyConnectedGradientNodeClass     _CNNFullyConnectedGradientNodeClass
	CNNFullyConnectedGradientNodeClassOnce sync.Once
)

func getCNNFullyConnectedGradientNodeClass() _CNNFullyConnectedGradientNodeClass {
	CNNFullyConnectedGradientNodeClassOnce.Do(func() {
		CNNFullyConnectedGradientNodeClass = _CNNFullyConnectedGradientNodeClass{objc.GetClass("MPSCNNFullyConnectedGradientNode")}
	})
	return CNNFullyConnectedGradientNodeClass
}

type _CNNFullyConnectedGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNFullyConnectedGradientNode */
// An interface definition for the [CNNFullyConnectedGradientNode] class.
type ICNNFullyConnectedGradientNode interface {
	ICNNConvolutionGradientNode
	
/* debug [class_interface_properties]: Properties for CNNFullyConnectedGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNFullyConnectedGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNFullyConnectedGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNFullyConnectedGradientNodeClass) Alloc() CNNFullyConnectedGradientNode {
	rv := objc.Send[CNNFullyConnectedGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNFullyConnectedGradientNodeClass) New() CNNFullyConnectedGradientNode {
	rv := objc.Send[CNNFullyConnectedGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNFullyConnectedGradientNode) Init() CNNFullyConnectedGradientNode {
	rv := objc.Send[CNNFullyConnectedGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNFullyConnectedGradientNode) Autorelease() CNNFullyConnectedGradientNode {
	rv := objc.Send[CNNFullyConnectedGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNFullyConnectedGradientNode creates a new CNNFullyConnectedGradientNode instance.
func NewCNNFullyConnectedGradientNode() CNNFullyConnectedGradientNode {
	return getCNNFullyConnectedGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNFullyConnectedGradientNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedGradientNode
type CNNFullyConnectedGradientNode struct {
	CNNConvolutionGradientNode
}

// CNNFullyConnectedGradientNodeFrom constructs a [CNNFullyConnectedGradientNode] from an unsafe.Pointer.
func CNNFullyConnectedGradientNodeFrom(ptr unsafe.Pointer) CNNFullyConnectedGradientNode {
	return CNNFullyConnectedGradientNode{
		CNNConvolutionGradientNode: CNNConvolutionGradientNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNFullyConnectedGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradientnode/3152566-initwithsourcegradient
func NewCNNFullyConnectedGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IImageNode, sourceImage IImageNode, gradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) CNNFullyConnectedGradientNode {
	instance := getCNNFullyConnectedGradientNodeClass().Alloc()
	rv := objc.Send[CNNFullyConnectedGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNFullyConnectedGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNFullyConnectedGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradientnode/3152567-nodewithsourcegradient
func (cc _CNNFullyConnectedGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IImageNode, sourceImage IImageNode, gradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageConvolutionGradientStateWeights) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNFullyConnectedGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNFullyConnectedGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNFullyConnectedGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNFullyConnectedGradientNode */


