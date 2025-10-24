// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingNode */


/* debug [class_header]: Header for MPSCNNPoolingNode */
// The class instance for the [CNNPoolingNode] class.
var (
	CNNPoolingNodeClass     _CNNPoolingNodeClass
	CNNPoolingNodeClassOnce sync.Once
)

func getCNNPoolingNodeClass() _CNNPoolingNodeClass {
	CNNPoolingNodeClassOnce.Do(func() {
		CNNPoolingNodeClass = _CNNPoolingNodeClass{objc.GetClass("MPSCNNPoolingNode")}
	})
	return CNNPoolingNodeClass
}

type _CNNPoolingNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingNode */
// An interface definition for the [CNNPoolingNode] class.
type ICNNPoolingNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNPoolingNode */
	// properties:
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	StrideInPixelsX() objectivec.IObject
	SetStrideInPixelsX(value objectivec.IObject)
	StrideInPixelsY() objectivec.IObject
	SetStrideInPixelsY(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingNodeClass) Alloc() CNNPoolingNode {
	rv := objc.Send[CNNPoolingNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingNodeClass) New() CNNPoolingNode {
	rv := objc.Send[CNNPoolingNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingNode) Init() CNNPoolingNode {
	rv := objc.Send[CNNPoolingNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingNode) Autorelease() CNNPoolingNode {
	rv := objc.Send[CNNPoolingNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingNode creates a new CNNPoolingNode instance.
func NewCNNPoolingNode() CNNPoolingNode {
	return getCNNPoolingNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingNode */
// A representation of a MPS CNN pooling kernel.


// A representation of a MPS CNN pooling kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode
type CNNPoolingNode struct {
	FilterNode
}

// CNNPoolingNodeFrom constructs a [CNNPoolingNode] from an unsafe.Pointer.
//
// A representation of a MPS CNN pooling kernel.
func CNNPoolingNodeFrom(ptr unsafe.Pointer) CNNPoolingNode {
	return CNNPoolingNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866488-initwithsource
func NewCNNPoolingNodeWithSourceFilterSize(sourceNode IImageNode, size uint) CNNPoolingNode {
	instance := getCNNPoolingNodeClass().Alloc()
	rv := objc.Send[CNNPoolingNode](instance.ID, objc.Sel("initWithSource:filterSize:"), sourceNode, size)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingNodeWithSourceFilterSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866444-initwithsource
func NewCNNPoolingNodeWithSourceFilterSizeStride(sourceNode IImageNode, size uint, stride uint) CNNPoolingNode {
	instance := getCNNPoolingNodeClass().Alloc()
	rv := objc.Send[CNNPoolingNode](instance.ID, objc.Sel("initWithSource:filterSize:stride:"), sourceNode, size, stride)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingNodeWithSourceFilterSizeStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866471-initwithsource
func NewCNNPoolingNodeWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(sourceNode IImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingNode {
	instance := getCNNPoolingNodeClass().Alloc()
	rv := objc.Send[CNNPoolingNode](instance.ID, objc.Sel("initWithSource:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), sourceNode, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingNodeWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866508-nodewithsource
func (cc _CNNPoolingNodeClass) NodeWithSourceFilterSize(sourceNode IImageNode, size uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:filterSize:"), sourceNode, size)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceFilterSize) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2890831-nodewithsource
func (cc _CNNPoolingNodeClass) NodeWithSourceFilterSizeStride(sourceNode IImageNode, size uint, stride uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:filterSize:stride:"), sourceNode, size, stride)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceFilterSizeStride) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993001-kernelheight
func (c_ CNNPoolingNode) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993001-kernelheight
func (c_ CNNPoolingNode) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993002-kernelwidth
func (c_ CNNPoolingNode) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993002-kernelwidth
func (c_ CNNPoolingNode) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993003-strideinpixelsx
func (c_ CNNPoolingNode) StrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsX"))
	return rv
}/* debug [instance_properties/getter]: strideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993003-strideinpixelsx
func (c_ CNNPoolingNode) SetStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsX:"), value)
}/* debug [instance_properties/setter]: strideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993004-strideinpixelsy
func (c_ CNNPoolingNode) StrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsY"))
	return rv
}/* debug [instance_properties/getter]: strideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993004-strideinpixelsy
func (c_ CNNPoolingNode) SetStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsY:"), value)
}/* debug [instance_properties/setter]: strideInPixelsY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingNode */


