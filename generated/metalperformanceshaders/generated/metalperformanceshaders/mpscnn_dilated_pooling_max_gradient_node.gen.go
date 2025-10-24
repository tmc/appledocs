// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNDilatedPoolingMaxGradientNode */


/* debug [class_header]: Header for MPSCNNDilatedPoolingMaxGradientNode */
// The class instance for the [CNNDilatedPoolingMaxGradientNode] class.
var (
	CNNDilatedPoolingMaxGradientNodeClass     _CNNDilatedPoolingMaxGradientNodeClass
	CNNDilatedPoolingMaxGradientNodeClassOnce sync.Once
)

func getCNNDilatedPoolingMaxGradientNodeClass() _CNNDilatedPoolingMaxGradientNodeClass {
	CNNDilatedPoolingMaxGradientNodeClassOnce.Do(func() {
		CNNDilatedPoolingMaxGradientNodeClass = _CNNDilatedPoolingMaxGradientNodeClass{objc.GetClass("MPSCNNDilatedPoolingMaxGradientNode")}
	})
	return CNNDilatedPoolingMaxGradientNodeClass
}

type _CNNDilatedPoolingMaxGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNDilatedPoolingMaxGradientNode */
// An interface definition for the [CNNDilatedPoolingMaxGradientNode] class.
type ICNNDilatedPoolingMaxGradientNode interface {
	ICNNPoolingGradientNode
	
/* debug [class_interface_properties]: Properties for CNNDilatedPoolingMaxGradientNode */
	// properties:
	DilationRateX() objectivec.IObject
	SetDilationRateX(value objectivec.IObject)
	DilationRateY() objectivec.IObject
	SetDilationRateY(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNDilatedPoolingMaxGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNDilatedPoolingMaxGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNDilatedPoolingMaxGradientNodeClass) Alloc() CNNDilatedPoolingMaxGradientNode {
	rv := objc.Send[CNNDilatedPoolingMaxGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDilatedPoolingMaxGradientNodeClass) New() CNNDilatedPoolingMaxGradientNode {
	rv := objc.Send[CNNDilatedPoolingMaxGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDilatedPoolingMaxGradientNode) Init() CNNDilatedPoolingMaxGradientNode {
	rv := objc.Send[CNNDilatedPoolingMaxGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDilatedPoolingMaxGradientNode) Autorelease() CNNDilatedPoolingMaxGradientNode {
	rv := objc.Send[CNNDilatedPoolingMaxGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDilatedPoolingMaxGradientNode creates a new CNNDilatedPoolingMaxGradientNode instance.
func NewCNNDilatedPoolingMaxGradientNode() CNNDilatedPoolingMaxGradientNode {
	return getCNNDilatedPoolingMaxGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNDilatedPoolingMaxGradientNode */
// A representation of a gradient dilated max pooling filter.


// A representation of a gradient dilated max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradientNode
type CNNDilatedPoolingMaxGradientNode struct {
	CNNPoolingGradientNode
}

// CNNDilatedPoolingMaxGradientNodeFrom constructs a [CNNDilatedPoolingMaxGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient dilated max pooling filter.
func CNNDilatedPoolingMaxGradientNodeFrom(ptr unsafe.Pointer) CNNDilatedPoolingMaxGradientNode {
	return CNNDilatedPoolingMaxGradientNode{
		CNNPoolingGradientNode: CNNPoolingGradientNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNDilatedPoolingMaxGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradientnode/2948026-initwithsourcegradient
func NewCNNDilatedPoolingMaxGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) CNNDilatedPoolingMaxGradientNode {
	instance := getCNNDilatedPoolingMaxGradientNodeClass().Alloc()
	rv := objc.Send[CNNDilatedPoolingMaxGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDilatedPoolingMaxGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNDilatedPoolingMaxGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradientnode/2948012-nodewithsourcegradient
func (cc _CNNDilatedPoolingMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNDilatedPoolingMaxGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNDilatedPoolingMaxGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNDilatedPoolingMaxGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradientnode/2947996-dilationratex
func (c_ CNNDilatedPoolingMaxGradientNode) DilationRateX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateX"))
	return rv
}/* debug [instance_properties/getter]: dilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradientnode/2947996-dilationratex
func (c_ CNNDilatedPoolingMaxGradientNode) SetDilationRateX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateX:"), value)
}/* debug [instance_properties/setter]: dilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradientnode/2948037-dilationratey
func (c_ CNNDilatedPoolingMaxGradientNode) DilationRateY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateY"))
	return rv
}/* debug [instance_properties/getter]: dilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradientnode/2948037-dilationratey
func (c_ CNNDilatedPoolingMaxGradientNode) SetDilationRateY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateY:"), value)
}/* debug [instance_properties/setter]: dilationRateY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNDilatedPoolingMaxGradientNode */


