// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingGradientNode */


/* debug [class_header]: Header for MPSCNNPoolingGradientNode */
// The class instance for the [CNNPoolingGradientNode] class.
var (
	CNNPoolingGradientNodeClass     _CNNPoolingGradientNodeClass
	CNNPoolingGradientNodeClassOnce sync.Once
)

func getCNNPoolingGradientNodeClass() _CNNPoolingGradientNodeClass {
	CNNPoolingGradientNodeClassOnce.Do(func() {
		CNNPoolingGradientNodeClass = _CNNPoolingGradientNodeClass{objc.GetClass("MPSCNNPoolingGradientNode")}
	})
	return CNNPoolingGradientNodeClass
}

type _CNNPoolingGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingGradientNode */
// An interface definition for the [CNNPoolingGradientNode] class.
type ICNNPoolingGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for CNNPoolingGradientNode */
	// properties:
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	StrideInPixelsX() objectivec.IObject
	SetStrideInPixelsX(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	StrideInPixelsY() objectivec.IObject
	SetStrideInPixelsY(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingGradientNodeClass) Alloc() CNNPoolingGradientNode {
	rv := objc.Send[CNNPoolingGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingGradientNodeClass) New() CNNPoolingGradientNode {
	rv := objc.Send[CNNPoolingGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingGradientNode) Init() CNNPoolingGradientNode {
	rv := objc.Send[CNNPoolingGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingGradientNode) Autorelease() CNNPoolingGradientNode {
	rv := objc.Send[CNNPoolingGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingGradientNode creates a new CNNPoolingGradientNode instance.
func NewCNNPoolingGradientNode() CNNPoolingGradientNode {
	return getCNNPoolingGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingGradientNode */
// A representation of a gradient pooling kernel.


// A representation of a gradient pooling kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode
type CNNPoolingGradientNode struct {
	GradientFilterNode
}

// CNNPoolingGradientNodeFrom constructs a [CNNPoolingGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient pooling kernel.
func CNNPoolingGradientNodeFrom(ptr unsafe.Pointer) CNNPoolingGradientNode {
	return CNNPoolingGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948011-initwithsourcegradient
func NewCNNPoolingGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy unsafe.Pointer) CNNPoolingGradientNode {
	instance := getCNNPoolingGradientNodeClass().Alloc()
	rv := objc.Send[CNNPoolingGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948045-nodewithsourcegradient
func (cc _CNNPoolingGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2947992-kernelheight
func (c_ CNNPoolingGradientNode) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2947992-kernelheight
func (c_ CNNPoolingGradientNode) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948018-strideinpixelsx
func (c_ CNNPoolingGradientNode) StrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsX"))
	return rv
}/* debug [instance_properties/getter]: strideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948018-strideinpixelsx
func (c_ CNNPoolingGradientNode) SetStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsX:"), value)
}/* debug [instance_properties/setter]: strideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948034-kernelwidth
func (c_ CNNPoolingGradientNode) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948034-kernelwidth
func (c_ CNNPoolingGradientNode) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948048-strideinpixelsy
func (c_ CNNPoolingGradientNode) StrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsY"))
	return rv
}/* debug [instance_properties/getter]: strideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948048-strideinpixelsy
func (c_ CNNPoolingGradientNode) SetStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsY:"), value)
}/* debug [instance_properties/setter]: strideInPixelsY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingGradientNode */


