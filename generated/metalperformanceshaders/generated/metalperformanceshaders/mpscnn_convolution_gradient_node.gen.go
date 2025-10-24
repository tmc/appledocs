// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolutionGradientNode */


/* debug [class_header]: Header for MPSCNNConvolutionGradientNode */
// The class instance for the [CNNConvolutionGradientNode] class.
var (
	CNNConvolutionGradientNodeClass     _CNNConvolutionGradientNodeClass
	CNNConvolutionGradientNodeClassOnce sync.Once
)

func getCNNConvolutionGradientNodeClass() _CNNConvolutionGradientNodeClass {
	CNNConvolutionGradientNodeClassOnce.Do(func() {
		CNNConvolutionGradientNodeClass = _CNNConvolutionGradientNodeClass{objc.GetClass("MPSCNNConvolutionGradientNode")}
	})
	return CNNConvolutionGradientNodeClass
}

type _CNNConvolutionGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolutionGradientNode */
// An interface definition for the [CNNConvolutionGradientNode] class.
type ICNNConvolutionGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for CNNConvolutionGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolutionGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolutionGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionGradientNodeClass) Alloc() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionGradientNodeClass) New() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionGradientNode) Init() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionGradientNode) Autorelease() CNNConvolutionGradientNode {
	rv := objc.Send[CNNConvolutionGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionGradientNode creates a new CNNConvolutionGradientNode instance.
func NewCNNConvolutionGradientNode() CNNConvolutionGradientNode {
	return getCNNConvolutionGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolutionGradientNode */
// A representation of a gradient convolution kernel.


// A representation of a gradient convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientNode
type CNNConvolutionGradientNode struct {
	GradientFilterNode
}

// CNNConvolutionGradientNodeFrom constructs a [CNNConvolutionGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient convolution kernel.
func CNNConvolutionGradientNodeFrom(ptr unsafe.Pointer) CNNConvolutionGradientNode {
	return CNNConvolutionGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolutionGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientnode/2947999-initwithsourcegradient
func NewCNNConvolutionGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IImageNode, sourceImage IImageNode, gradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) CNNConvolutionGradientNode {
	instance := getCNNConvolutionGradientNodeClass().Alloc()
	rv := objc.Send[CNNConvolutionGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolutionGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientnode/2947984-nodewithsourcegradient
func (cc _CNNConvolutionGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IImageNode, sourceImage IImageNode, gradientState ICNNConvolutionGradientStateNode, weights unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageConvolutionGradientStateWeights) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolutionGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolutionGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolutionGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolutionGradientNode */


