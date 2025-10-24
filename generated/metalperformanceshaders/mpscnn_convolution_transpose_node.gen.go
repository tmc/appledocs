// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolutionTransposeNode */


/* debug [class_header]: Header for MPSCNNConvolutionTransposeNode */
// The class instance for the [CNNConvolutionTransposeNode] class.
var (
	CNNConvolutionTransposeNodeClass     _CNNConvolutionTransposeNodeClass
	CNNConvolutionTransposeNodeClassOnce sync.Once
)

func getCNNConvolutionTransposeNodeClass() _CNNConvolutionTransposeNodeClass {
	CNNConvolutionTransposeNodeClassOnce.Do(func() {
		CNNConvolutionTransposeNodeClass = _CNNConvolutionTransposeNodeClass{objc.GetClass("MPSCNNConvolutionTransposeNode")}
	})
	return CNNConvolutionTransposeNodeClass
}

type _CNNConvolutionTransposeNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolutionTransposeNode */
// An interface definition for the [CNNConvolutionTransposeNode] class.
type ICNNConvolutionTransposeNode interface {
	ICNNConvolutionNode
	
/* debug [class_interface_properties]: Properties for CNNConvolutionTransposeNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolutionTransposeNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolutionTransposeNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionTransposeNodeClass) Alloc() CNNConvolutionTransposeNode {
	rv := objc.Send[CNNConvolutionTransposeNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionTransposeNodeClass) New() CNNConvolutionTransposeNode {
	rv := objc.Send[CNNConvolutionTransposeNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionTransposeNode) Init() CNNConvolutionTransposeNode {
	rv := objc.Send[CNNConvolutionTransposeNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionTransposeNode) Autorelease() CNNConvolutionTransposeNode {
	rv := objc.Send[CNNConvolutionTransposeNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionTransposeNode creates a new CNNConvolutionTransposeNode instance.
func NewCNNConvolutionTransposeNode() CNNConvolutionTransposeNode {
	return getCNNConvolutionTransposeNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolutionTransposeNode */
// A representation of a transposed convolution.


// A representation of a transposed convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeNode
type CNNConvolutionTransposeNode struct {
	CNNConvolutionNode
}

// CNNConvolutionTransposeNodeFrom constructs a [CNNConvolutionTransposeNode] from an unsafe.Pointer.
//
// A representation of a transposed convolution.
func CNNConvolutionTransposeNodeFrom(ptr unsafe.Pointer) CNNConvolutionTransposeNode {
	return CNNConvolutionTransposeNode{
		CNNConvolutionNode: CNNConvolutionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolutionTransposeNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposenode/2942641-initwithsource
func NewCNNConvolutionTransposeNodeWithSourceConvolutionGradientStateWeights(sourceNode IImageNode, convolutionGradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) CNNConvolutionTransposeNode {
	instance := getCNNConvolutionTransposeNodeClass().Alloc()
	rv := objc.Send[CNNConvolutionTransposeNode](instance.ID, objc.Sel("initWithSource:convolutionGradientState:weights:"), sourceNode, convolutionGradientState, weights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionTransposeNodeWithSourceConvolutionGradientStateWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolutionTransposeNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposenode/2942636-nodewithsource
func (cc _CNNConvolutionTransposeNodeClass) NodeWithSourceConvolutionGradientStateWeights(sourceNode IImageNode, convolutionGradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:convolutionGradientState:weights:"), sourceNode, convolutionGradientState, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceConvolutionGradientStateWeights) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolutionTransposeNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolutionTransposeNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolutionTransposeNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolutionTransposeNode */


